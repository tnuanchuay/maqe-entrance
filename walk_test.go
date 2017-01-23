package main

import "testing"

func TestItShouldBe1X0YWhenWalk0DegAnd1Step(t *testing.T){
	var p position
	walk(&p, 0, 1)
	if p.x != 1 || p.y != 0{
		t.Error("expected 1, 0 but got", p.x, p.y)
	}
}

func TestItShouldBe0X1YWhenWalk90DegAnd1Step(t *testing.T){
	var p position
	walk(&p, 90, 1)
	if p.x != 0 || p.y != 1{
		t.Error("expected 0, 1 but got", p.x, p.y)
	}
}

func TestItShouldBe_1X0YWhenWalk180DegAnd1Step(t *testing.T){
	var p position
	walk(&p, 180, 1)
	if p.x != -1 || p.y != 0{
		t.Error("expected -1, 0 but got", p.x, p.y)
	}
}

func TestItShouldBe0X_1YWhenWalk270DegAnd1Step(t *testing.T){
	var p position
	walk(&p, 270, 1)
	if p.x != 0 || p.y != -1{
		t.Error("expected 0, -1 but got", p.x, p.y)
	}
}

func TestItShouldBe15X0YWhenWalk0DegAnd15Step(t *testing.T){
	var p position
	walk(&p, 0, 15)
	if p.x != 15 || p.y != 0{
		t.Error("expected 15, 0 but got", p.x, p.y)
	}
}
