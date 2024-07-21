//go:build darwin

package logos

// const appleLogo = `
//
//	                     )s
//	                  sQQQQ
//	                )QQQQQb
//	                QQQQb
//	        )oo     PP)  oppo
//	    sQQQQQQQQQQQSQQQQQQQQQQQp
//	  )QQQQQQQQQQQQQQQQQQQQQQQQQbD
//	 (QQQQQQQQQQQQQQQQQQQQQQQQSb
//	 QQQQQQQQQQQQQQQQQQQQQQQQQb
//	)QQQQQQQQQQQQQQQQQQQQQQQQQ
//	 QQQQQQQQQQQQQQQQQQQQQQQQQo
//	 QQQQQQQQQQQQQQQQQQQQQQQQQQp
//	 )QQQQQQQQQQQQQQQQQQQQQQQQQQSp
//	  )QQQQQQQQQQQQQQQQQQQQQQQQQQQb
//	   )QQQQQQQQQQQQQQQQQQQQQQQQQb
//	    (QQQQQQQQQQQQQQQQQQQQQQSb
//	      )SQQQQQQSSbSSQQQQQQSb
//	         PDb          PP)
//
// `
var appleLogo = [18]string{
	"                         )s             ",
	"                      sQQQQ             ",
	"                    )QQQQQb             ",
	"                    QQQQb               ",
	"            )oo     PP)  oppo           ",
	"        sQQQQQQQQQQQSQQQQQQQQQQQp       ",
	"      )QQQQQQQQQQQQQQQQQQQQQQQQQbD      ",
	"     (QQQQQQQQQQQQQQQQQQQQQQQQSb        ",
	"     QQQQQQQQQQQQQQQQQQQQQQQQQb         ",
	"    )QQQQQQQQQQQQQQQQQQQQQQQQQ          ",
	"     QQQQQQQQQQQQQQQQQQQQQQQQQo         ",
	"     QQQQQQQQQQQQQQQQQQQQQQQQQQp        ",
	"     )QQQQQQQQQQQQQQQQQQQQQQQQQQSp      ",
	"      )QQQQQQQQQQQQQQQQQQQQQQQQQQQb     ",
	"       )QQQQQQQQQQQQQQQQQQQQQQQQQb      ",
	"        (QQQQQQQQQQQQQQQQQQQQQQSb       ",
	"          )SQQQQQQSSbSSQQQQQQSb         ",
	"             PDb          PP)           ",
}

func Logos(os string) [18]string {
	return appleLogo
}
