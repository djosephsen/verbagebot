package bothandlers

import (
	"fmt"
	"github.com/djosephsen/hal"
	"math/rand"
	"time"
)

var Quantifyme = &hal.Handler{
	Method:  hal.RESPOND,
	Pattern: `quantify me`,
	Run: func(res *hal.Response) error {
		now:=time.Now()
		rand.Seed(int64(now.Unix()))
		states:=[]string{`passive aggressive`, `mads`, `fucks`,`horrible`}
		state:=states[rand.Intn(len(states)-1)]
		var reply string
		switch state {
		case `horrible`,`passive aggressive`:
			reply=fmt.Sprintf(`you are currently %%%d.%04d %s`,rand.Intn(int(100)),rand.Intn(int(1000)),state)
		case `mads`:
			reply=fmt.Sprintf(`you are %d.%04d %s`,rand.Intn(int(2)),rand.Intn(int(1000)),state)
		case `fucks`:
			reply=fmt.Sprintf(`you give precisely %f %s`,rand.Float64(),state)
		}

		return res.Reply(reply)
	},
}
