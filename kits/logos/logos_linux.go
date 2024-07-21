//go:build linux

package logos

import "strings"

// const ubuntuLogo = `
//
//	            )sSQQQQQQQQQQso
//	        )SQQQQQQQQQQQQQQQQQQQQo
//	      SQQQQQQQQQQQQQQQQQbCCPSQQQb
//	    SQQQQQQQQQQQSbbbbbGb     QQQQQb
//	   QQQQQQQQQbGQ        Pp  )QQQQQQQQp
//	  QQQQQQQQb    Q)spQpso     PQQQQQQQQp
//	 QQQQQQQQb    QQQQQQQQQQQo    SQQQQQQQ
//	)QQQSbbbS   (QQQQQQQQQQQQQp   )QQQQQQQb
//	)QQS     Q  QQQQQQQQQQQQQQGssspQQQQQQQb
//	)QQQp   sb  GQQQQQQQQQQQQQS   )QQQQQQQb
//	 QQQQQQQQ    GQQQQQQQQQQQb    QQQQQQQQb
//	 )QQQQQQQQ     SQQQQQQSb     QQQQQQQQb
//	  PQQQQQQQQb  Qb        sDGSQQQQQQQQb
//	   )QQQQQQQQQQQpo     (b    )QQQQQQb
//	     PQQQQQQQQQQQQQQQQQQ    sQQQQb
//	       PSQQQQQQQQQQQQQQQQQQQQQSb
//	          (PSQQQQQQQQQQQQQSbC
//	                 (CDCC
//
// `
var ubuntuLogo = [18]string{
	"              )sSQQQQQQQQQQso              ",
	"          )SQQQQQQQQQQQQQQQQQQQQo          ",
	"        SQQQQQQQQQQQQQQQQQbCCPSQQQb        ",
	"      SQQQQQQQQQQQSbbbbbGb     QQQQQb      ",
	"     QQQQQQQQQbGQ        Pp  )QQQQQQQQp    ",
	"    QQQQQQQQb    Q)spQpso     PQQQQQQQQp   ",
	"   QQQQQQQQb    QQQQQQQQQQQo    SQQQQQQQ   ",
	"  )QQQSbbbS   (QQQQQQQQQQQQQp   )QQQQQQQb  ",
	"  )QQS     Q  QQQQQQQQQQQQQQGssspQQQQQQQb  ",
	"  )QQQp   sb  GQQQQQQQQQQQQQS   )QQQQQQQb  ",
	"   QQQQQQQQ    GQQQQQQQQQQQb    QQQQQQQQb  ",
	"   )QQQQQQQQ     SQQQQQQSb     QQQQQQQQb   ",
	"    PQQQQQQQQb  Qb        sDGSQQQQQQQQb    ",
	"     )QQQQQQQQQQQpo     (b    )QQQQQQb     ",
	"       PQQQQQQQQQQQQQQQQQQ    sQQQQb       ",
	"         PSQQQQQQQQQQQQQQQQQQQQQSb         ",
	"            (PSQQQQQQQQQQQQQSbC            ",
	"                   (CDCC                   ",
}

// const archLogo = `
//
//	                 )p
//	                 QQp
//	                QQQQp
//	               QQQQQQp
//	              QQQQQQQQp
//	              GQQQQQQQQp
//	            QQppSQQQQQQQp
//	           QQQQQQQQQQQQQQp
//	         )QQQQQQQQQQQQQQQQb
//	        (QQQQQQQQQQQQQQQQQQb
//	       SQQQQQQQQbbbGQQQQQQQQS
//	      SQQQQQQQS)     GQQQQQQQQ
//	     QQQQQQQQQ        QQQQQQQQQp
//	    QQQQQQQQQb        )QQQQQbQbGp
//	  (QQQQQQQQQQb        )QQQQQQQQQp
//	 sQQQQQQSbbC            CbbbQQQQQQb
//	SQQSbD                        CPbQQQ
//
// SbC                                 PGo
// `
var archLogo = [18]string{
	"                    p                   ",
	"                   QQp                  ",
	"                  QQQQp                 ",
	"                 QQQQQQp                ",
	"                QQQQQQQQp               ",
	"                 QQQQQQQQp              ",
	"              QQp SQQQQQQQp             ",
	"             QQQQQQQQQQQQQQp            ",
	"           )QQQQQQQQQQQQQQQQb           ",
	"          (QQQQQQQQQQQQQQQQQQb          ",
	"         SQQQQQQQQbbbGQQQQQQQQS         ",
	"        SQQQQQQQS)     GQQQQQQQQ        ",
	"       QQQQQQQQQ        QQQQQQQQQp      ",
	"      QQQQQQQQQb        )QQQQQbQ  p     ",
	"    (QQQQQQQQQQb        )QQQQQQQQb      ",
	"   sQQQQQQSbbC            CbbbQQQQQQb   ",
	"  SQQSbD                        CPbQQQ  ",
	" SbC                                 PGo",
}

var centosLogo = [18]string{
	"                    GGQo                   ",
	"                  SGGGGGQo                 ",
	"        oooooo  oooo Gb)ssss esssssp       ",
	"        GGGG (QGGGGG Gb)QQQQQbcSQQQb       ",
	"        GD GQo GGGGG Gb)QQQSb)QSbPSb       ",
	"         GGo GGo GGG Gb)QSbsQSbsQQp(       ",
	"        GGGGGs )Gs ) Gb)bsQbbsQQQQQQ       ",
	"    sQb GGGGGGGGc)QG ) )Qb)sQQQQQQQb)Qb    ",
	"  sQQQQQQQQQQQQQQQp     )QQQQQQQQQQQQQQQbc ",
	"   PQQb)QQQQQQQQpcpp     p)pppppppppdQQbb  ",
	"     Pb)QQQQQQSbsQSb Go PQs PGGGGGGb)bP    ",
	"        QQQQbbsQbbsQ GO)s )Gs )GGGGC       ",
	"       )pPbbsQb)sQQQ GO)GGSc)QSc)C)G       ",
	"       )QQQQb)QQQQQQ GO)GGGGQo(GQQGb       ",
	"       )QQQQQpPQQQQQ GO)GGGGb (QGGGb       ",
	"        ((((((  cspp(GDooo                 ",
	"                   GGGGGC                  ",
	"                     GC                    ",
}

//var genericLogo =[21]string{
//	"               sQQQQQQSQQ               ",
//	"              QQQQQQQQQQQQQ             ",
//	"             )QQQQQQQSQSQQQp            ",
//	"             )bQQPQbCQQPQQQb            ",
//	"             )bGSSbbdQS SQQb            ",
//	"             )bQGGSGGGQQQQQQ            ",
//	"             dQGQSSSQQSQQQGSb           ",
//	"             SbOGbbbbC  (QQQQb          ",
//	"           sQb            QQQQQp        ",
//	"         )QQb            o)QQQQQQc      ",
//	"        )SQbC             )CQQQQQQp     ",
//	"        QQb                  QQQQQQ     ",
//	"       QSQ                   SQQQQQQ    ",
//	"      QQSb                   )QQQQQQ    ",
//	"     )bbSb                   SQQQQbb    ",
//	"    sQGGSGQQ               QGSQQQQbGp   ",
//	" bGGGGGGGGGSQQQ           GQGGGbbGSGS   ",
//	" SGGGGGGGGGGGSS          (sQGGGGGGGGGGQo",
//	")GGGGGGGGGGGGSp        sQQQbGGGGGGGQQQSb",
//	"bQQQSQQQGGGGGQQQQQQQQQQQQQQQQQQQQQSbb   ",
//	"    (PbbGQQQQQbc      (((PPQQQQQbb      ",
//}

var genericLogo = [18]string{
	"             sQQQQQSe             ",
	"            QQQQQQQQQQQ           ",
	"           )SSQQQSbGSQQb          ",
	"           )bQpQb)Qp)QQb          ",
	"           )bSGGGGGbQQQb          ",
	"           )QQQQQSQSSQbQ          ",
	"           sbGbbbbC  bQQQc        ",
	"         sQb          QQQQb       ",
	"       )QQb          oCQQQQQQ     ",
	"       SSb              PQQQQb    ",
	"      SQb                QQSQQo   ",
	"     QSQ                 QQQQQb   ",
	"     bbQ                 QQQQSb   ",
	" sssSGGGGQp            bQSQQSbS   ",
	" GGGGGGGGQQQb         GbGGGGGGGb  ",
	" GGGGGGGGGGQb       )sQbGGGGGGGQQS",
	"QQQQGGGGGGGQSQeeseQQQQQbGGGQQQbbP ",
	"  (PbbGQQQQQbb)((( (PbGQQQQSb     ",
}

//`
//               sQQQQQQSQQ
//              QQQQQQQQQQQQQ
//             )QQQQQQQSQSQQQp
//             )bQQPQbCQQPQQQb
//             )bGSSbbdQS SQQb
//             )bQGGSGGGQQQQQQ
//             dQGQSSSQQSQQQGSb
//             SbOGbbbbC  (QQQQb
//           sQb            QQQQQp
//         )QQb            o)QQQQQQc
//        )SQbC             )CQQQQQQp
//        QQb                  QQQQQQ
//       QSQ                   SQQQQQQ
//      QQSb                   )QQQQQQ
//     )bbSb                   SQQQQbb
//    sQGGSGQQ               QGSQQQQbGp
// bGGGGGGGGGSQQQ           GQGGGbbGSGS
// SGGGGGGGGGGGSS          (sQGGGGGGGGGGQo
//)GGGGGGGGGGGGSp        sQQQbGGGGGGGQQQSb
//bQQQSQQQGGGGGQQQQQQQQQQQQQQQQQQQQQSbb
//    (PbbGQQQQQbc      (((PPQQQQQbb
//`

func Logos(os string) [18]string {
	if strings.Contains(strings.ToLower(os), "ubuntu") {
		return ubuntuLogo
	} else if strings.Contains(strings.ToLower(os), "arch") {
		return archLogo
	} else if strings.Contains(strings.ToLower(os), "centos") {
		return centosLogo
	} else {
		return genericLogo
	}
}
