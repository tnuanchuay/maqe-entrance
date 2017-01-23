package main

import "testing"

func TestItShouldBe180WhenTurnL(t *testing.T){
	deg := 90
	turn(&deg, L)
	if deg != 180 {
		t.Error("expeted 180 but got", deg)
	}
}

func TestItShouldBe0WhenTurnR(t *testing.T){
	deg := 90
	turn(&deg, R)
	if deg != 0{
		t.Error("expected 0 but got", deg)
	}
}

func TestItShouldBe270WhenTurnRAndStartAt0Deg(t *testing.T){
	deg := 0
	turn(&deg, R)
	if deg != 270{
		t.Error("expected 0 but got", deg)
	}
}

func TestItShouldBe0WhenTurnLAndStartAt270Deg(t *testing.T){
	deg := 270
	turn(&deg, L)
	if deg != 0{
		t.Error("expected 270 but got", deg)
	}
}
