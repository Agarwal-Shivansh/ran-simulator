// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mobility

import (
	"context"
	"github.com/onosproject/onos-api/go/onos/ransim/types"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/ran-simulator/pkg/model"
	"github.com/onosproject/ran-simulator/pkg/store/cells"
	"github.com/onosproject/ran-simulator/pkg/store/routes"
	"github.com/onosproject/ran-simulator/pkg/store/ues"
	"github.com/onosproject/ran-simulator/pkg/utils"
	"math"
	"math/rand"
	"time"
)

var log = logging.GetLogger("mobility", "driver")

// Driver is an abstraction of an entity driving the UE mobility
type Driver interface {
	// Start starts the driving engine
	Start(ctx context.Context)

	// Stop stops the driving engine
	Stop()

	// GenerateRoutes generates routes for all UEs that currently do not have a route; remove routes with no UEs
	GenerateRoutes(ctx context.Context, minSpeed uint32, maxSpeed uint32, speedStdDev uint32)
}

type driver struct {
	cellStore  cells.Store
	routeStore routes.Store
	ueStore    ues.Store
	apiKey     string
	ticker     *time.Ticker
	done       chan bool
	min        *model.Coordinate
	max        *model.Coordinate
}

// NewMobilityDriver returns a driving engine capable of "driving" UEs along pre-specified routes
func NewMobilityDriver(cellStore cells.Store, routeStore routes.Store, ueStore ues.Store, apiKey string) Driver {
	return &driver{
		cellStore:  cellStore,
		routeStore: routeStore,
		ueStore:    ueStore,
	}
}

var tickUnit = time.Second

const tickFrequency = 1

func (d *driver) Start(ctx context.Context) {
	log.Info("Driver starting")

	// Iterate over all routes and position the UEs at the start of their routes
	for _, route := range d.routeStore.List(ctx) {
		d.initializeUEPosition(ctx, route)
	}

	d.ticker = time.NewTicker(tickFrequency * tickUnit)
	d.done = make(chan bool)

	go d.drive(ctx)
}

func (d *driver) Stop() {
	log.Info("Driver stopping")
	d.ticker.Stop()
	d.done <- true
}

func (d *driver) drive(ctx context.Context) {
	for {
		select {
		case <-d.done:
			ctx.Done()
			close(d.done)
			return
		case <-d.ticker.C:
			for _, route := range d.routeStore.List(ctx) {
				if route.NextPoint == 0 && !route.Reverse {
					d.initializeUEPosition(ctx, route)
				}
				d.updateUEPosition(ctx, route)
				d.updateUESignalStrength(ctx, route.IMSI)
			}
		}
	}
}

// Initializes UE positions to the start of its routes.
func (d *driver) initializeUEPosition(ctx context.Context, route *model.Route) {
	bearing := utils.InitialBearing(*route.Points[0], *route.Points[1])
	_ = d.ueStore.MoveToCoordinate(ctx, route.IMSI, *route.Points[0], uint32(math.Round(bearing)))
	_ = d.routeStore.Start(ctx, route.IMSI, route.SpeedAvg, route.SpeedStdDev)
}

func (d *driver) updateUEPosition(ctx context.Context, route *model.Route) {
	// Get the UE
	ue, err := d.ueStore.Get(ctx, route.IMSI)
	if err != nil {
		log.Warn("Unable to find UE %d", route.IMSI)
		return
	}

	// Determine speed and heading
	speed := float64(route.SpeedAvg) + rand.NormFloat64()*float64(route.SpeedStdDev)
	distanceDriven := (tickFrequency * speed) / 3600.0

	// Determine bearing and distance to the next point
	bearing := utils.InitialBearing(ue.Location, *route.Points[route.NextPoint])
	remainingDistance := utils.Distance(ue.Location, *route.Points[route.NextPoint])

	// If distance is less than to the next waypoint, determine the coordinate along that vector
	// Otherwise just use the next waypoint
	newPoint := *route.Points[route.NextPoint]
	reachedWaypoint := remainingDistance <= distanceDriven
	if !reachedWaypoint {
		newPoint = utils.TargetPoint(ue.Location, bearing, distanceDriven)
	}

	// Move the UE to the determined coordinate; update heading if necessary
	err = d.ueStore.MoveToCoordinate(ctx, route.IMSI, newPoint, uint32(math.Round(bearing)))
	if err != nil {
		log.Warn("Unable to update UE %d coordinates", route.IMSI)
	}

	// Update the route if necessary
	if reachedWaypoint {
		_ = d.routeStore.Advance(ctx, route.IMSI)
	}
}

func (d *driver) updateUESignalStrength(ctx context.Context, imsi types.IMSI) {
	ue, err := d.ueStore.Get(ctx, imsi)
	if err != nil {
		log.Warn("Unable to find UE %d", imsi)
		return
	}

	for _, uecell := range ue.Cells {
		cell, err := d.cellStore.Get(ctx, uecell.ECGI)
		if err != nil {
			log.Warn("Unable to find cell %d", uecell.ECGI)
			continue
		}
		uecell.Strength = StrengthAtLocation(ue.Location, *cell)
	}
	err = d.ueStore.UpdateCells(ctx, imsi, ue.Cells)
	if err != nil {
		log.Warn("Unable to update UE %d cell info", imsi)
	}
}