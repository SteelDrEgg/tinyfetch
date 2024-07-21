//g :build wind ws

package logos

import "strings"

//const win10 = `
//
//
//                           llllllllllll
//           lllll    lllllllllllllllllll
//   lllllllllllll    lllllllllllllllllll
//   lllllllllllll    lllllllllllllllllll
//   lllllllllllll    lllllllllllllllllll
//   lllllllllllll    lllllllllllllllllll
//   lllllllllllll    lllllllllllllllllll
//
//
//   lllllllllllll    lllllllllllllllllll
//   lllllllllllll    lllllllllllllllllll
//   lllllllllllll    lllllllllllllllllll
//   lllllllllllll    lllllllllllllllllll
//   lllllllllllll    lllllllllllllllllll
//            llll    lllllllllllllllllll
//                            lllllllllll
//
//
//`

var win10Logo = [18]string{
	"                                           ",
	"                           llllllllllll    ",
	"           lllll    lllllllllllllllllll    ",
	"   lllllllllllll    lllllllllllllllllll    ",
	"   lllllllllllll    lllllllllllllllllll    ",
	"   lllllllllllll    lllllllllllllllllll    ",
	"   lllllllllllll    lllllllllllllllllll    ",
	"   lllllllllllll    lllllllllllllllllll    ",
	"                                           ",
	"                                           ",
	"   lllllllllllll    lllllllllllllllllll    ",
	"   lllllllllllll    lllllllllllllllllll    ",
	"   lllllllllllll    lllllllllllllllllll    ",
	"   lllllllllllll    lllllllllllllllllll    ",
	"   lllllllllllll    lllllllllllllllllll    ",
	"            llll    lllllllllllllllllll    ",
	"                            lllllllllll    ",
	"                                           ",
}

//const winxp = `
//
//         @@@@@@@@@@@@@@
//         @@@@@@@@@@@@@@@
//        @@@@@@@@@@@@@@@   @@
//       @@@@@@@@@@@@@@@    @@@@@@@@@@@@@@@
//       @@@@@@@@@@@@@@@   @@@@@@@@@@@@@@@
//       @@@@@@@@@@@@@@    @@@@@@@@@@@@@@@
//      @@@@@@@@@@@@@@@   @@@@@@@@@@@@@@@@
//     @            @@    @@@@@@@@@@@@@@@
//                       @@@@@@@@@@@@@@@
//     @@@@@@@@@@@@@      @@@@@@@@@@@@@@
//    @@@@@@@@@@@@@@@         @@@@@
//   @@@@@@@@@@@@@@@    @
//   @@@@@@@@@@@@@@@   @@@@@@@@@@@@@@@
//  @@@@@@@@@@@@@@@    @@@@@@@@@@@@@@@
//  @@@@@@@@@@@@@@@   @@@@@@@@@@@@@@@
// @@@@@@@@@@@@@@@    @@@@@@@@@@@@@@@
//             @@@   @@@@@@@@@@@@@@@
//                   @@@@@@@@@@@@@@@
//                   @@@@@@@@@@@@@@
//
//`

var winxpLogo = [18]string{
	"         @@@@@@@@@@@@@@@                  ",
	"        @@@@@@@@@@@@@@@   @@              ",
	"       @@@@@@@@@@@@@@@    @@@@@@@@@@@@@@@ ",
	"       @@@@@@@@@@@@@@@   @@@@@@@@@@@@@@@  ",
	"       @@@@@@@@@@@@@@    @@@@@@@@@@@@@@@  ",
	"      @@@@@@@@@@@@@@@   @@@@@@@@@@@@@@@@  ",
	"     @            @@    @@@@@@@@@@@@@@@   ",
	"                       @@@@@@@@@@@@@@@    ",
	"     @@@@@@@@@@@@@      @@@@@@@@@@@@@@    ",
	"    @@@@@@@@@@@@@@@         @@@@@         ",
	"   @@@@@@@@@@@@@@@    @                   ",
	"   @@@@@@@@@@@@@@@   @@@@@@@@@@@@@@@      ",
	"  @@@@@@@@@@@@@@@    @@@@@@@@@@@@@@@      ",
	"  @@@@@@@@@@@@@@@   @@@@@@@@@@@@@@@       ",
	" @@@@@@@@@@@@@@@    @@@@@@@@@@@@@@@       ",
	"             @@@   @@@@@@@@@@@@@@@        ",
	"                   @@@@@@@@@@@@@@@        ",
	"                                          ",
}

func Logos(os string) [18]string {
	if strings.Contains(os, "10") || strings.Contains(os, "11") {
		return win10Logo
	} else {
		return winxpLogo
	}
}
