package libbring

import "testing"

func TestTrack(t *testing.T) {
	trackingNumber := "TESTPACKAGE-AT-PICKUPPOINT" // "READY_FOR_PICKUP"

	if _, err := Track(trackingNumber); err != nil {
		t.Errorf("Track(%v) err = %v, want nil\n", trackingNumber, err)
	}
}

func TestPickByPostalCode(t *testing.T) {
	postalCode := "4051" // Sola, Norge

	if _, err := PickByPostalCode(postalCode); err != nil {
		t.Errorf("PickByPostalCode(%v) err = %v, want nil\n", postalCode, err)
	}
}

func TestPickByPostalCodeWithTextQuery(t *testing.T) {
	postalCode := "4051"
	query := "COOP" // As in, COOP-Mega (grocery store)

	if _, err := PickByPostalCodeWithTextQuery(postalCode, query); err != nil {
		t.Errorf("PickByPostalCodeWithQuery(%v, %v) err = %v, want nil\n", postalCode, query, err)
	}
}

func TestPickByPostalCodeWithOpeningHoursQuery(t *testing.T) {
	postalCode := "4051"
	hoursQuery := "SUNDAY"

	if _, err := PickByPostalCodeWithOpeningHoursQuery(postalCode, hoursQuery); err != nil {
		t.Errorf("PickByPostalCodeWithOpeningHoursQuery(%v, %v) err = %v, want nil\n", postalCode, hoursQuery, err)
	}
}

func TestPickById(t *testing.T) {
	id := "121110" // Oppegård, Norge

	if _, err := PickById(id); err != nil {
		t.Errorf("PickById(%v) err = %v, want nil", id, err)
	}
}

func TestPickByLocation(t *testing.T) {
	lat, lon := float64(60.39266538671475), float64(5.326359579942273) // Bergen, Norge

	if _, err := PickByLocation(lat, lon); err != nil {
		t.Errorf("PickByLocation(%v, %v) err = %v, want nil", lat, lon, err)
	}
}
