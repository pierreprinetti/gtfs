package gtfs

// "log"

// Trip.Direction possible values: (There is no "in" or "out" directions, they symbolize two opposite directions)
const (
	DirectionOut = iota // 0 - travel in one direction (e.g. outbound travel)
	DirectionIn         //  1 - travel in the opposite direction (e.g. inbound travel)
)

// Trip.WheelchairAccessible possible values:
const (
	_                       = iota // 0 - no accessibility information for the trip
	WheelchairAccessible           // 1 - the vehicle being used on this particular trip can accommodate at least one rider in a wheelchair
	WheelchairNotAccessible        // 2 - no riders in wheelchairs can be accommodated on this trip
)

const (
	_               = iota // 0 - no bike information for the trip
	BikesAllowed           // 1 - the vehicle being used on this particular trip can accommodate at least one bicycle
	BikesNotAllowed        // 2 - no bicycles are allowed on this trip
)

// trips.txt
type Trip struct {

	// route_id - Required. The route_id field contains an ID that uniquely identifies a route.
	// This value is referenced from the routes.txt file.
	Route *Route

	// service_id - Required. The service_id contains an ID that uniquely identifies a set of dates when service
	// is available for one or more routes. This value is referenced from the calendar.txt or calendar_dates.txt file.
	ServiceId string

	// trip_id	- Required. The trip_id field contains an ID that identifies a trip. The trip_id is dataset unique.
	Id string

	// trip_headsign - Optional. The trip_headsign field contains the text that appears on a sign that identifies the
	// trip's destination to passengers. Use this field to distinguish between different patterns of service in the
	// same route. If the headsign changes during a trip, you can override the trip_headsign by specifying values
	// for the the stop_headsign field in stop_times.txt.
	// See a Google Maps screenshot highlighting the headsign:
	// http://code.google.com/transit/spec/transit_feed_specification.html#transitTripHeadsignScreenshot
	Headsign string

	// trip_short_name	- Optional. The trip_short_name field contains the text that appears in schedules and sign boards
	// to identify the trip to passengers, for example, to identify train numbers for commuter rail trips. If riders do
	// not commonly rely on trip names, please leave this field blank.
	// A trip_short_name value, if provided, should uniquely identify a trip within a service day; it should not be used
	// for destination names or limited/express designations.
	ShortName string

	// direction_id - Optional. The direction_id field contains a binary value that indicates the direction of travel for
	// a trip. Use this field to distinguish between bi-directional trips with the same route_id. This field is not used
	// in routing; it provides a way to separate trips by direction when publishing time tables. You can specify names for
	// each direction with the trip_headsign field.
	// 		0 - travel in one direction (e.g. outbound travel)
	// 		1 - travel in the opposite direction (e.g. inbound travel)
	// For example, you could use the trip_headsign and direction_id fields together to assign a name to travel in each
	// direction on trip "1234", the trips.txt file would contain these rows for use in time tables:
	//
	// 		trip_id, ... ,trip_headsign,direction_id
	// 		1234, ... , to Airport,0
	// 		1505, ... , to Downtown,1
	//
	// See DirectionIn/Out constants
	Direction byte

	// block_id - Optional. The block_id field identifies the block to which the trip belongs. A block consists of two
	// or more sequential trips made using the same vehicle, where a passenger can transfer from one trip to the next just
	// by staying in the vehicle. The block_id must be referenced by two or more trips in trips.txt.
	BlockId string

	// shape_id - Optional. The shape_id field contains an ID that defines a shape for the trip. This value is referenced
	// from the shapes.txt file. The shapes.txt file allows you to define how a line should be drawn on the map to represent a trip.
	ShapeId string

	// wheelchair_accessible - Optional.
	// 0 (or empty) - indicates that there is no accessibility information for the trip
	// 1 - indicates that the vehicle being used on this particular trip can accommodate at least one rider in a wheelchair
	// 2 - indicates that no riders in wheelchairs can be accommodated on this trip.
	WheelchairAccessible byte

	// bikes_allowed - Optional.
	// 0 (or empty) - indicates that there is no bike information for the trip
	// 1 - indicates that the vehicle being used on this particular trip can accommodate at least one bicycle
	// 2 - indicates that no bicycles are allowed on this trip
	BikesAllowed byte
}
