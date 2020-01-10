package hufid

import (
	"bytes"
	"fmt"
	"testing"
)

func ExampleUseUniqID() {
	id := NewUniqID(5)
	fmt.Println(id)
}

func ExampleUseID() {
	id := NewID(2, bytes.NewReader([]byte{0x11, 0x22, 0x33, 0x44}))
	fmt.Println(id)
	// Output: KJ863-2100A
}

func ExampleIDValidation() {
	var id ID = "KJ863-2100A"
	fmt.Println(id.Validate())
	id += "5"
	fmt.Println(id.Validate())
	// Output:
	// true
	// false
}

func ExampleIDNormalization() {
	var id ID = "kJB63-2IO0A"
	fmt.Println(id.Normalize())
	fmt.Println(id)
	// Output:
	// true
	// KJ863-2100A
}

func TestBasic(t *testing.T) {
	id := NewUniqID(5)
	if len(id) != 4*(symbolsPerGroup+1)+symbolsPerGroup {
		t.Errorf("Generated id does not have expected length (%d)", len(id))
	}
	if id.Normalize() != true {
		t.Errorf("Generated id must be valid during normalizing (%s)", id)
	}
	if id.Validate() != true {
		t.Errorf("Generated id must be valid (%s)", id)
	}
}

func TestBasicFromData(t *testing.T) {
	id := NewID(2, bytes.NewReader([]byte{0x28}))
	if len(id) != 1*(symbolsPerGroup+1)+symbolsPerGroup {
		t.Errorf("Generated id does not have expected length (%d)", len(id))
	}
	if id != "81000-0000Q" {
		t.Errorf("Generated id must be equal to expected id (%s)", id)
	}
	if id.Normalize() != true {
		t.Errorf("Generated id must be valid during normalizing (%s)", id)
	}
	if id.Validate() != true {
		t.Errorf("Generated id must be valid (%s)", id)
	}
}

func TestNormalization(t *testing.T) {
	data := [...]struct {
		id    ID
		norm  ID
		valid bool
	}{
		{id: "", norm: "", valid: false},
		{id: "SVXVC-2T9E2-75VKJ-YCL3E-GX7C6", norm: "SVXVC-2T9E2-75VKJ-YCL3E-GX7C6", valid: true},
		{id: "SVXVc-2T9e2-75VKJ-YCL3E-GX7C6", norm: "SVXVC-2T9E2-75VKJ-YCL3E-GX7C6", valid: true},
		{id: "SVXVC2T9E275VKJYCL3EGX7C6", norm: "SVXVC-2T9E2-75VKJ-YCL3E-GX7C6", valid: true},
		{id: "svXvC2T9E2-75VKJ-YCL3E-Gx7C6", norm: "SVXVC-2T9E2-75VKJ-YCL3E-GX7C6", valid: true},
		{id: "SVXVC-2T9E2-75VKJ-YCL3E-GX7C5", norm: "SVXVC-2T9E2-75VKJ-YCL3E-GX7C6", valid: false},
		{id: "svXvC2T9A2-75VKJ-YCL3E-Gx7C6", norm: "SVXVC-2T9E2-75VKJ-YCL3E-GX7C6", valid: false},
		{id: "svDvC2T9E2-75VKJ-YCL3E-Gx7C6", norm: "SVXVC-2T9E2-75VKJ-YCL3E-GX7C6", valid: false},
		{id: "svXvC2T9E2-75V#%J-YCL3E-Gx7C6", norm: "SVXVC-2T9E2-75VKJ-YCL3E-GX7C6", valid: false},
		{id: "81000-00000-0000C", norm: "81000-00000-0000C", valid: true},
		{id: "b1000-00000-0000C", norm: "81000-00000-0000C", valid: true},
		{id: "8i000-00000-0000C", norm: "81000-00000-0000C", valid: true},
		{id: "8iOOO-00000-0000C", norm: "81000-00000-0000C", valid: true},
		{id: "biooo-dDdDd-ooooc", norm: "81000-00000-0000C", valid: true},
		{id: "VSDKVsdsdv", norm: "", valid: false},
		{id: "81001", norm: "81001", valid: true},
		{id: "BiDoI", norm: "81001", valid: true},
		{id: "8-1-0-0-1", norm: "81001", valid: true},
		{id: "--8--1--0--0--1", norm: "81001", valid: true},
		{id: "--8--1--0--0--1-", norm: "81001", valid: false},
	}

	data[2].id.Normalize()

	for i, cs := range data {
		id := cs.id
		if id.Validate() != cs.valid {
			t.Errorf("Validation case failed for %s (case %d)", id, i)
		}
		isnorm := id.Normalize()
		if isnorm != cs.valid {
			t.Errorf("Normalization failed for %s (case %d) by validation", id, i)
		}
		if isnorm == true {
			if id != cs.norm {
				t.Errorf("Normalization failed for %s (case %d) as expected %s", id, i, cs.norm)
			}
		}
	}
}
