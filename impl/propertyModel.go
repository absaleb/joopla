package impl

// import (
// 	"github.com/globalsign/mgo/bson"
// 	"time"
// )
//
// type PropertyAvailabilitySlotModel struct {
// 	WeekDay   int32  `bson:"weekday"`
// 	TimeStart int32  `bson:"time_start"`
// 	TimeEnd   int32  `bson:"time_end"`
// 	Type      string `bson:"type"`
// 	Day       int32  `bson:"day"`
// 	Month     int32  `bson:"month"`
// 	Year      int32  `bson:"year"`
// }
//
// type PropertyModel struct {
// 	Id               bson.ObjectId                    `bson:"_id,omitempty"`
// 	City             *string                          `bson:"city,omitempty"`
// 	Address          *string                          `bson:"address,omitempty"`
// 	Postcode         *string                          `bson:"postcode,omitempty"`
// 	CountryCode      *string                          `bson:"country_code,omitempty"`
// 	BuildingId       *string                          `bson:"building_id,omitempty"`
// 	LandlordId       *string                          `bson:"landlord_id,omitempty"`
// 	SubmitterId      *string                          `bson:"submitter_id,omitempty"`
// 	TenantIds        []string                         `bson:"tenant_ids,omitempty"`
// 	ThreadId         *string                          `bson:"thread_id,omitempty"`
// 	CreatedAt        *time.Time                       `bson:"created_at,omitempty"`
// 	Status           string                           `bson:"status,omitempty"`
// 	Type             string                           `bson:"type,omitempty"`
// 	Sqm              *float64                         `bson:"sqm,omitempty"`
// 	UkTaxBand        string                           `bson:"uk_tax_band,omitempty"`
// 	Tags             []string                         `bson:"tags,omitempty,omitempty"`
// 	StartedListingAt *time.Time                       `bson:"started_listing_at,omitempty"`
// 	EndedListingAt   *time.Time                       `bson:"ended_listing_at,omitempty"`
// 	Availability     []*PropertyAvailabilitySlotModel `bson:"availability,omitempty"`
// 	ImageB64         *string                          `bson:"image_b64,omitempty"`
// 	Location         []float64                        `bson:"location,omitempty"`
// }
