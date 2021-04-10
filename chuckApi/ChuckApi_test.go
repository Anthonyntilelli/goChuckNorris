package chuckApi

import (
	"sort"
	"testing"
)

func TestRandomFact(t *testing.T) {
	result1, err1 := RandomFact()
	if err1 != nil {
		t.Errorf("Error during call1 (%v)", err1)
	}
	result2, err2 := RandomFact()
	if err2 != nil {
		t.Errorf("Error during call2 (%v)", err2)
	}
	if !result1.Valid() || !result2.Valid() {
		t.Error("result1 or result 2 are not valid")
	} else if result1.Id == result2.Id {
		// Test ID, as jokes may have been re-submitted to site.
		t.Error("result 1 and result 2 Id should not be the same.")
	}
}

func TestCategorieslist(t *testing.T) {
	result1, err1 := Categorieslist()
	if err1 != nil {
		t.Errorf("Error during call1 (%v)", err1)
	}
	result2, err2 := Categorieslist()
	if err2 != nil {
		t.Errorf("Error during call2 (%v)", err2)
	}

	if len(result1) != len(result2) {
		t.Error("results1 and results2 length are not the same.")
	} else {
		sort.Strings(result1)
		sort.Strings(result2)
		if result1[1] != result2[1] {
			t.Error("result 1 and result 2 should be the same.")
		}
	}
}

func TestRandomFactByCategory(t *testing.T) {
	categories, cErr := Categorieslist()
	if cErr != nil {
		t.Errorf("Error during category call (%v)", cErr)
	} else {
		result1, err1 := RandomFactByCategory(categories[0])
		if err1 != nil {
			t.Errorf("Error during call2 (%v)", err1)
		}
		// Invalid category
		invalidResult, err2 := RandomFactByCategory("kdjfslkfjwpoiej")
		// Error expected on invalid provided category
		if err2 == nil {
			t.Errorf("Error is expected with invalid Category")
		}
		if !result1.Valid() {
			t.Error("result1 is not valid")
		}
		if invalidResult.Valid() {
			t.Error("invalidResult is valid when it should not be")
		}
	}
}

func TestRandomFactbytext(t *testing.T) {
	result1, err1 := RandomFactbytext("potato")
	if err1 != nil {
		t.Errorf("Error during call1 (%v)", err1)
	}
	if result1.Id == "" {
		t.Error("Invalid result1 id")
	}
	if result1.Value == "" {
		t.Error("Invalid Result1 value")
	}
	// Expect error
	_, err2 := RandomFactbytext("sadahdakshdaoidlaksdjhaohdlkdowihlidhali")
	if err2 == nil {
		t.Errorf("Did not error to invalid searchTerm")
	}
	_, err3 := RandomFactbytext("potato skin")
	if err3 == nil {
		t.Errorf("Did not error to MultiWord searchTerm")
	}
}

func TestEmergencyFact(t *testing.T) {
	result := EmergencyFact()
	if result.Id != "{Emergency}" {
		t.Errorf("Emergency fact id is not `{Emergency}` instead `%v`", result.Id)
	}
	if result.Categories != nil {
		t.Errorf("Emergency fact categories should be empty instead `%v`", result.Categories)
	}
	if !result.Valid() {
		t.Error("result should valid")
	}
}

func TestValid(t *testing.T) {
	var c ChuckFact
	if c.Valid() {
		t.Error("Blank ChuckFact should not be valid")
	}
}
