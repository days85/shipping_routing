package routing

import (
    cargo `github.com/days85/shipping_cargo`
)

// Service provides access to an external routing service.
type Service interface {
    // FetchRoutesForSpecification finds all possible routes that satisfy a
    // given specification.
    FetchRoutesForSpecification(rs cargo.RouteSpecification) []cargo.Itinerary
}
