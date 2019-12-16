package tests

import (
	"testing"

	"github.com/pjvds/go-cqrs/sourcing"
	"github.com/pjvds/go-cqrs/tests/domain"
	"github.com/pjvds/go-cqrs/tests/events"
	. "github.com/smartystreets/goconvey/convey"
)

func TestStateChangesAreRepresentedByEvents(t *testing.T) {

	Convey("Given we create a new user 'pjvds'", t, func() {
		user := domain.NewUser("pjvds")

		Convey("Then the username should be 'pjvds'", func() {
			So(user.Username, ShouldEqual, "pjvds")
		})

		Convey("And the creation should be captured by an event", func() {
			So(len(user.Events()), ShouldEqual, 1)
		})

		Convey("When we change the username to 'wwwouter'", func() {
			user.ChangeUsername("wwwouter")

			Convey("Then this should also be captured by an event", func() {
				So(len(user.Events()), ShouldEqual, 2)
			})
		})
	})
}

func TestDomainObjectCanBeBuildFromHistory(t *testing.T) {

	Convey("Given the user and the name changes from the last test", t, func() {

		// The id of our event source that we will rebuild from history.
		sourceId, _ := sourcing.ParseEventSourceId("0791d279-664d-458e-bf60-567ade140832")

		// The full history for the User domain object
		history := []sourcing.Event{
			// It was first created
			events.UserCreated{
				Username: "pjvds",
			},
			// Then the username was changed
			events.UsernameChanged{
				OldUsername: "pjvds",
				NewUsername: "wwwouter",
			},
		}

		Convey("When we create a new User domain object from the event history", func() {
			user := domain.NewUserFromHistory(sourceId, history)

			Convey("Then the username should not be 'pjvds'", func() {
				So(user.Username, ShouldNotEqual, "pjvds")
			})

			Convey("But it should be 'wwwouter'", func() {
				So(user.Username, ShouldEqual, "wwwouter")
			})
		})
	})
}
