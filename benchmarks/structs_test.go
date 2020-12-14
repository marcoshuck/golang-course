package benchmarks

import (
	"testing"
	"time"
)

type Post struct {
}

type repositoryTest struct {
	Post Post
}

func (r *repositoryTest) TestA() (*Post, error) {
	time.Sleep(20 * time.Millisecond)
	return &r.Post, nil
}

func (r *repositoryTest) TestB() (*Post, error) {
	var p Post
	time.Sleep(20 * time.Millisecond)
	return &p, nil
}

func BenchmarkStructPreInitialized(b *testing.B) {
	var r repositoryTest
	for i := 0; i < b.N; i++ {
		r.TestA()
	}
}

func BenchmarkStructNoInitialized(b *testing.B) {
	var r repositoryTest
	for i := 0; i < b.N; i++ {
		r.TestB()
	}
}
