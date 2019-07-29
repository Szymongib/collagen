package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var testDoggos = Doggos{
	{
		Name:      "Fluffy",
		BirthDate: time.Date(2012, 1, 1, 0, 0, 0, 0, time.Local).Unix(),
		Height:    25,
	},
	{
		Name:      "Occy",
		BirthDate: time.Date(2016, 1, 1, 0, 0, 0, 0, time.Local).Unix(),
		Height:    102,
	},
	{
		Name:      "Boo",
		BirthDate: time.Date(2013, 1, 1, 0, 0, 0, 0, time.Local).Unix(),
		Height:    49,
	},
	{
		Name:      "Poo",
		BirthDate: time.Date(2009, 1, 1, 0, 0, 0, 0, time.Local).Unix(),
		Height:    17,
	},
	{
		Name:      "Rex",
		BirthDate: time.Date(2015, 1, 1, 0, 0, 0, 0, time.Local).Unix(),
		Height:    86,
	},
}

func TestDoggos_Contains(t *testing.T) {

	newDoggo := Doggo{
		Name:      "Pup",
		BirthDate: time.Date(2019, 7, 23, 18, 0, 0, 0, time.Local).Unix(),
		Height:    10,
	}

	testDoggosWithNewDoggo := append(testDoggos, newDoggo)

	for _, testCase := range []struct {
		description   string
		doggos        Doggos
		shouldContain bool
	}{
		{
			description:   "contain new doggo",
			doggos:        testDoggosWithNewDoggo,
			shouldContain: true,
		},
		{
			description:   "not contain new doggo",
			doggos:        testDoggos,
			shouldContain: false,
		},
	} {
		t.Run("should "+testCase.description, func(t *testing.T) {
			// when
			contains := testCase.doggos.Contains(newDoggo)

			// then
			assert.Equal(t, testCase.shouldContain, contains)
		})
	}

}

func TestDoggos_Drop(t *testing.T) {

	for _, testCase := range []struct {
		description string
		doggos      Doggos
		toDrop      int
		result      Doggos
	}{
		{
			description: "drop 2 Doggos",
			doggos:      testDoggos,
			toDrop:      2,
			result: Doggos{
				{Name: "Fluffy", BirthDate: time.Date(2012, 1, 1, 0, 0, 0, 0, time.Local).Unix(), Height: 25},
				{Name: "Occy", BirthDate: time.Date(2016, 1, 1, 0, 0, 0, 0, time.Local).Unix(), Height: 102},
				{Name: "Boo", BirthDate: time.Date(2013, 1, 1, 0, 0, 0, 0, time.Local).Unix(), Height: 49},
			},
		},
		{
			description: "drop 0 doggos",
			doggos:      testDoggos,
			toDrop:      0,
			result:      testDoggos,
		},
		{
			description: "drop all if dropping more doggos than exists",
			doggos:      testDoggos,
			toDrop:      len(testDoggos) + 2,
			result:      Doggos{},
		},
	} {
		t.Run("should "+testCase.description, func(t *testing.T) {
			// when
			result := testCase.doggos.Drop(testCase.toDrop)

			// then
			assert.Equal(t, testCase.result, result)
		})
	}
}

func TestDoggos_Exist(t *testing.T) {

	for _, testCase := range []struct {
		description  string
		doggos       Doggos
		existsFunc   func(doggo Doggo) bool
		shouldExists bool
	}{
		{
			description: "return true if doggo exists",
			doggos:      testDoggos,
			existsFunc: func(doggo Doggo) bool {
				return doggo.Name == "Fluffy"
			},
			shouldExists: true,
		},
		{
			description: "return false if doggo does not exist",
			doggos:      testDoggos,
			existsFunc: func(doggo Doggo) bool {
				return doggo.Height > 30000
			},
			shouldExists: false,
		},
	} {
		t.Run("should "+testCase.description, func(t *testing.T) {
			// when
			exists := testCase.doggos.Exist(testCase.existsFunc)

			// then
			assert.Equal(t, testCase.shouldExists, exists)
		})
	}

}

func TestDoggos_Filter(t *testing.T) {

	for _, testCase := range []struct {
		description string
		doggos      Doggos
		filterFunc  func(doggo Doggo) bool
		result      Doggos
	}{
		{
			description: "return doggos shorter than 50",
			doggos:      testDoggos,
			filterFunc: func(doggo Doggo) bool {
				return doggo.Height < 50
			},
			result: Doggos{
				{Name: "Fluffy", BirthDate: time.Date(2012, 1, 1, 0, 0, 0, 0, time.Local).Unix(), Height: 25},
				{Name: "Boo", BirthDate: time.Date(2013, 1, 1, 0, 0, 0, 0, time.Local).Unix(), Height: 49},
				{Name: "Poo", BirthDate: time.Date(2009, 1, 1, 0, 0, 0, 0, time.Local).Unix(), Height: 17},
			},
		},
		{
			description: "return 0 doggos if non satisfy filter func",
			doggos:      testDoggos,
			filterFunc: func(doggo Doggo) bool {
				return len(doggo.Name) > 1000
			},
			result: Doggos{},
		},
		{
			description: "return all doggos if all satisfy filter func",
			doggos:      testDoggos,
			filterFunc: func(doggo Doggo) bool {
				return true
			},
			result: testDoggos,
		},
	} {
		t.Run("should "+testCase.description, func(t *testing.T) {
			// when
			result := testCase.doggos.Filter(testCase.filterFunc)

			// then
			assert.Equal(t, testCase.result, result)
		})
	}
}

func TestDoggos_Take(t *testing.T) {

	for _, testCase := range []struct {
		description string
		doggos      Doggos
		toTake      int
		result      Doggos
	}{
		{
			description: "take 2 Doggos",
			doggos:      testDoggos,
			toTake:      2,
			result: Doggos{
				{Name: "Fluffy", BirthDate: time.Date(2012, 1, 1, 0, 0, 0, 0, time.Local).Unix(), Height: 25},
				{Name: "Occy", BirthDate: time.Date(2016, 1, 1, 0, 0, 0, 0, time.Local).Unix(), Height: 102},
			},
		},
		{
			description: "take 0 doggos",
			doggos:      testDoggos,
			toTake:      0,
			result:      Doggos{},
		},
		{
			description: "take all all doggos if take more than exists",
			doggos:      testDoggos,
			toTake:      len(testDoggos) + 2,
			result:      testDoggos,
		},
	} {
		t.Run("should "+testCase.description, func(t *testing.T) {
			// when
			result := testCase.doggos.Take(testCase.toTake)

			// then
			assert.Equal(t, testCase.result, result)
		})
	}
}
