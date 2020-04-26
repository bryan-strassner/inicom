package inicom

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	unitTestDir string = `../../test/unit/`
)

var (
	basefile string = fmt.Sprintf("%sbase.ini", unitTestDir)
	addfile  string = fmt.Sprintf("%sadd.ini", unitTestDir)
	subfile  string = fmt.Sprintf("%ssub.ini", unitTestDir)
)

func TestLoadIni(t *testing.T) {
	t.Parallel()
	bf, err := loadIni(basefile)
	assert.Nil(t, err)
	assert.NotNil(t, bf)
	// just test that we used ini package correctly
	s1, err := bf.GetSection("section 1")
	assert.Nil(t, err)
	assert.NotNil(t, s1)
}

func TestParse(t *testing.T) {
	t.Parallel()
	af, err := Parse([]string{"add", addfile, "subtract", subfile})
	assert.Nil(t, err)
	assert.Equal(t, 2, len(af))
	assert.Equal(t, "add", af[0].Action)
	assert.Equal(t, "subtract", af[1].Action)
	s1, err := af[0].File.GetSection("section 1")
	assert.Nil(t, err)
	assert.Equal(t, "section 1", s1.Name())
}

func TestProcess(t *testing.T) {
	t.Parallel()
	bf, err := Basefile(basefile)
	assert.Nil(t, err)
	assert.NotNil(t, bf)
	af, err := Parse([]string{"add", addfile, "subtract", subfile})
	assert.Nil(t, err)
	Process(bf, af)
	s1, err := bf.GetSection("section 1")
	assert.Nil(t, s1)
	s2, err := bf.GetSection("section 2")
	assert.Nil(t, err)
	assert.NotNil(t, s2)
	s2v2, err := s2.GetKey("s2Value2")
	assert.Nil(t, err)
	assert.Equal(t, "case is different", s2v2.String())
}
