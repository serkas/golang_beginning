package main

func NewA() *ComponentA {
	return &ComponentA{
		b: &ComponentB{}, // 1) "hard" binding 2) each new ComponentA creates another ComponentB
	}
}

type ComponentA struct {
	b *ComponentB
}

type ComponentB struct {
}
