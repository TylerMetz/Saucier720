
// /*
// package main

// import (
// 	"BackendPkg"
// 	//"fmt"
// 	//"time"
// )


// func main() {
// 		fmt.Println("Welcome to out Sprint 1 demo!")
// 		userPantry := BackendPkg.Pantry{
// 			TimeLastUpdated: time.Now(),
// 		}
// 		userPantry.AddToPantry()
// 		userPantry.AddToPantry()
// 		userPantry.AddToPantry()
// 		userPantry.DisplayPantry()
// 		userPantry.RemoveFromPantry()
// 		userPantry.RemoveFromPantry()
// 		userPantry.DisplayPantry()
// }
// */

// package main

// import (
// 	"fmt"
// 	"github.com/k4s/phantomgo"
// 	"io/ioutil"
// 	//. "github.com/k4s/webrowser"
// 	//"net/http"
// )

// func main() {
// 	p := phantomgo.NewPhantom()

// 	js := `
// var page = require('webpage').create(),
//   system = require('system'),
//   address;
// page.settings.userAgent = 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/37.0.2062.120 Safari/537.36';
// phantom.cookiesEnabled = true;
// phantom.addCookie({
// 	'name'     : 'store_rpt',
// 	'value'    : '1',
// 	'domain'   : 'www.publix.com',
// 	'path'     : '/',
// });
// phantom.addCookie({
// 	'name'     : 'x-ms-routing-name',
// 	'value'    : 'self',
// 	'domain'   : '.www.publix.com',
// 	'path'     : '/',
// });
// phantom.addCookie({
// 	'name'     : 'TiPMix',
// 	'value'    : '79.1668783049623',
// 	'domain'   : '.www.publix.com',
// 	'path'     : '/',
// });
// phantom.addCookie({
// 	'name'     : '_4c_',
// 	'value'    : 'jVPfT9swEP5XKj%2FwRBL%2FjO1KaCoFTZtgiHVIe6uc5KBR0yRy0gaG%2BN85tykgJk3Lg3X3%2BbvzfXeXZzKsoCZTlmpNmVbGUkNPyRqeOjJ9Jnkbzl04tr4iU7Lq%2B7abJskwDHG7zaryMc6bTdK5XVk%2FdMkAsK6eIlckuxKGyFUVOSV5UwCGMhubWKHf%2F0EvEpSiDTUmJ60v0L65%2FvVzeX45m9%2F8%2BPDSBnqPr3x4LEu6Ljn4rW8SlnxfRDzmIqbR1fz3IumsZkwIprmVwpgvs9vzM3ayKYszxqgVhqVCyFRZbihNuaCccaq0NVZrRblgJ7PbyzOGBX2dLe%2B%2BXYTSpZDcIk3Gx0YpjYS7Dvzkqnko68mid31QeVev62ao8XLRNx4mC6gg76EIWVQaJC%2Bg68oGI5qtz0PIRemRgjeoptjm%2FbJ%2FagM%2BQDbpijVeFLArc1gOZdGv9p3k9B1dQfmw6gOMowtw64OD1lDWRTN8DhvRD2E0kDPfDCgHgfnKNxuYYLMQbnAPyLXL0fRwD97vKeh15V7v%2B1xGDFfnDQ6awnw5GlWTuypEQH2c9fX80N3%2Fmgp5OSWPh0VlVFhjpNG4Sz1upUklDR8yfFmMG0sk5xJMkUegcxnJVOSRBVpE1jJ2D2CdcEDGnEZqK5QyqQ1JduUxR0qZc7l0EXWZiqRTIjJC3EcZVzpzKhVq37vH4w%2BEOjSTcqyLmWNZbTVmZO9kwawy2pgjWb6JaHcjm3%2BSTLn5W%2FJhcP%2BIsZ9jXl5eAQ%3D%3D',
// 	'domain'   : '.publix.com',
// 	'path'     : '/',
// });
// phantom.addCookie({
// 	'name'     : 's_tp',
// 	'value'    : '5652',
// 	'domain'   : '.publix.com',
// 	'path'     : '/',
// });
// phantom.addCookie({
// 	'name'     : '_uetvid',
// 	'value'    : 'd9cb4cb0b23411ed82695bfd593c245b',
// 	'domain'   : '.publix.com',
// 	'path'     : '/',
// });
// phantom.addCookie({
// 	'name'     : '_uetsid',
// 	'value'    : 'd9c9d500b23411ed80d2effd1bdedfaa',
// 	'domain'   : '.publix.com',
// 	'path'     : '/',
// });
// phantom.addCookie({
// 	'name'     : '_ga',
// 	'value'    : 'GA1.1.1434291204.1677017557',
// 	'domain'   : '.publix.com',
// 	'path'     : '/',
// });
// phantom.addCookie({
// 	'name'     : 'Store',
// 	'value'    : '{%22CreationDate%22:%222023-02-21T22:12:54.050Z%22%2C%22ForceRefreshed%22:false%2C%22Option%22:%22ACFJNORSTVY%22%2C%22ShortStoreName%22:%22University%20Vlg%20Mkt%22%2C%22StoreName%22:%22Publix%20At%20University%20Village%20Market%22%2C%22StoreNumber%22:1560}',
// 	'domain'   : '.publix.com',
// 	'path'     : '/',
// });
// phantom.addCookie({
// 	'name'     : 'ShoppingListCount',
// 	'value'    : '0',
// 	'domain'   : '.publix.com',
// 	'path'     : '/',
// });
// phantom.addCookie({
// 	'name'     : '_gid',
// 	'value'    : 'GA1.2.1549355166.1677017557',
// 	'domain'   : '.publix.com',
// 	'path'     : '/',
// });
// phantom.addCookie({
// 	'name'     : 'FPLC',
// 	'value'    : 'gMINW4sevUMps3lQLnQOa8%2F%2F9dNaED8k%2FtZyxF8pSmTrXmNTEKAF%2BAZwZx2n%2F0g2FNMKNGwOHQ300%2Fo5euhQAd6izBJakB2zzvvmoyj6Adj6Jy0nEQZOa9ezuRaGig%3D%3D',
// 	'domain'   : '.publix.com',
// 	'path'     : '/',
// });
// phantom.addCookie({
// 	'name'     : 's_ips',
// 	'value'    : '1001',
// 	'domain'   : '.publix.com',
// 	'path'     : '/',
// });
// phantom.addCookie({
// 	'name'     : '_ga_1DWX30JN6C',
// 	'value'    : 'GS1.1.1677017557.1.1.1677017582.0.0.0',
// 	'domain'   : '.publix.com',
// 	'path'     : '/',
// });
// phantom.addCookie({
// 	'name'     : '_fbp',
// 	'value'    : 'fb.1.1677017558246.1302144086',
// 	'domain'   : '.publix.com',
// 	'path'     : '/',
// });
// phantom.addCookie({
// 	'name'     : '_gcl_au',
// 	'value'    : '1.1.929578126.1677017556',
// 	'domain'   : '.publix.com',
// 	'path'     : '/',
// });
// phantom.addCookie({
// 	'name'     : 's_ecid',
// 	'value'    : 'MCMID%7C11093816334659280062302120579897750231',
// 	'domain'   : '.publix.com',
// 	'path'     : '/',
// });
// phantom.addCookie({
// 	'name'     : 'AMCVS_0B25161260B4AFD00A495E9C%40AdobeOrg',
// 	'value'    : '1',
// 	'domain'   : '.publix.com',
// 	'path'     : '/',
// });
// phantom.addCookie({
// 	'name'     : 'akavpau_VP_WWW_PRD',
// 	'value'    : '1677018188~id=198b0259def670900dbcff3ef15d5b5c',
// 	'domain'   : 'www.publix.com',
// 	'path'     : '/',
// });
// phantom.addCookie({
// 	'name'     : '_ripgplsjne',
// 	'value'    : '9acb25b3-2ad9-4213-a223-6c103788d750',
// 	'domain'   : '.publix.com',
// 	'path'     : '/',
// });
// phantom.addCookie({
// 	'name'     : 'FPID',
// 	'value'    : 'FPID2.2.A3pExBAnE%2FAoXrtCZGosG9nMPihNHtwCxDDJuWcwvb8%3D.1677017557',
// 	'domain'   : '.publix.com',
// 	'path'     : '/',
// });
// phantom.addCookie({
// 	'name'     : 's_cc',
// 	'value'    : 'true',
// 	'domain'   : '.publix.com',
// 	'path'     : '/',
// });
// phantom.addCookie({
// 	'name'     : '_pin_unauth',
// 	'value'    : 'dWlkPU16YzROREV5TURjdE1EZzNNUzAwTVdOa0xXSTFNek10WXpjd1kyWmlZbUZoTnpNMA',
// 	'domain'   : '.publix.com',
// 	'path'     : '/',
// });
// phantom.addCookie({
// 	'name'     : '_abck',
// 	'value'    : '10B92E36789FEE071DA625AE763092F6~0~YAAQLEsfRfGf+HKGAQAADyEJdgmv/uMU5VcimdNTufUHDL73+dJB1zNIXgadLOLykRGUVAVNndz/Gp8+HgAQ31bN2nWZjp1Kpi3prRRfgXxS5sIR2FuO3Yya9eFJLWKUal4wMqKn7WltcFZDREGpXXRhUmM1O6SIiXsGLplWM3y7DelsgINLEfMO6619WcC2v8xiQQSPm/aTwiqOlrFPJxcpcJgaEN+CkzJpXFEr0F2zZa/+UQPQyxMnS///5owtZmdP7ETn0Trfc7G3Ar2lwc6NkY1YW7wr0oB3OfQSNvqaD3Re6MD2QSwkdzCfm1cu++JqUkRPHohn0pXTW8snv1wV1+I/GQikbae44tXL/MFvbJDRCoHRs142IQeinkYzjxlej9nK+T8nASDv4/DMi1tbL2upzAQZ~-1~||1-LMXiTZORee-1-10-1000-2||~1677021111',
// 	'domain'   : '.publix.com',
// 	'path'     : '/',
// });
// phantom.addCookie({
// 	'name'     : '_dc_gtm_UA-34592858-5',
// 	'value'    : '1',
// 	'domain'   : '.publix.com',
// 	'path'     : '/',
// });
// phantom.addCookie({
// 	'name'     : 'AMCV_0B25161260B4AFD00A495E9C%40AdobeOrg',
// 	'value'    : '179643557%7CMCIDTS%7C19410%7CMCMID%7C11093816334659280062302120579897750231%7CMCAID%7CNONE%7CMCOPTOUT-1677024758s%7CNONE%7CMCAAMLH-1677622358%7C7%7CMCAAMB-1677622358%7Cj8Odv6LonN4r3an7LhD3WZrU1bUpAkFkkiY1ncBR96t2PTI%7CMCSYNCSOP%7C411-19417%7CvVersion%7C5.5.0',
// 	'domain'   : '.publix.com',
// 	'path'     : '/',
// });
// phantom.addCookie({
// 	'name'     : '_ga_QZBH7YJQYK',
// 	'value'    : 'GS1.1.1677017557.1.1.1677017582.0.0.0',
// 	'domain'   : '.publix.com',
// 	'path'     : '/',
// });
// phantom.addCookie({
// 	'name'     : 'ak_bmsc',
// 	'value'    : '88787A437F34446F12F4D42ADC0CB02E~000000000000000000000000000000~YAAQLEsfRfKf+HKGAQAAnSEJdhLEeAEjAacM9GrOIpZA8Mud0pou2tQK+yowuXKRVFT3v/6KNU5jFaOXkhGx8ro6JZiVbdkIgO/vtGXhY7tnhKqHMa2xAT6j0bEQY0/VGU/syBvHxz2zsSzHuJVeySH1XLm7+6iwY88uSMoLDfHSx0oTuTA06qCW3NLwcWJKEUeyr+NmSP1VqxAHg9CT+Ar7JQIc9kI1G6k5lPVYBfjpEzXCUzJb/7FhhKT9GBskgZfWMfsOj+ls1th2G6r6OfU8blLLrcFtu6O/SIjG0dNWW90+RWRzklJ5ct4yMbyPEPcUnKu/kxmpQoEDIMPN+yVcetfqDvmfytU3KQUrCrIHnpO2Ro2594hZvKG024kxJb0Ikf6A/b6F7MGaxdU18w+EwUcPBlTLmA25lSxaCDljPYXHIxGAsxSJ3Bt7KlB+SgHoIpeiRFS6yUftNwJCU/ZbRbK680mLAoLJm3wDlQ5I3VS7EpSOjzU=',
// 	'domain'   : '.publix.com',
// 	'path'     : '/',
// });
// phantom.addCookie({
// 	'name'     : 'bm_sz',
// 	'value'    : '4A3BC2CD5F582F404CFB0F6250041661~YAAQLEsfRbaf+HKGAQAACBQJdhLj16tkY63cJlQmVon+EUrqwKi3KlU3EzFBo/306zDp+oOvdWKsuRI4/okWNzC+h2XgFV0iK36EZTYTnFWlYCoH+2TdGGzIIBFyzYFWgkPK/To12WqUx0unp64CFh3QKGAKIenwX536kwSCc7zlhkm3xamMBAzPvh26ZElrRsMfCmeWDt/1h9lz98a5BsVcPQmGXQ0r0Z0TpuHTOtqM09YbSGnf5wU7l2j0AsWgpK8zMbMfbwYYQIhLYfAdhlJAhW837cWgjOWMY9S8wnPcBL0=~4535617~3687220',
// 	'domain'   : '.publix.com',
// 	'path'     : '/',
// });
// phantom.addCookie({
// 	'name'     : 's_ppv',
// 	'value'    : '%252Fsavings%252Fweekly-ad%252Fview-all%2C18%2C18%2C18%2C1001%2C5%2C1',
// 	'domain'   : '.publix.com',
// 	'path'     : '/',
// });
// phantom.addCookie({
// 	'name'     : 'dtCookie',
// 	'value'    : 'v_4_srv_9_sn_159C4F9C44085C83150A874C76F3320F_perc_27704_ol_1_app-3Aa52023d6b6739bcd_0',
// 	'domain'   : '.publix.com',
// 	'path'     : '/',
// });
// phantom.addCookie({
// 	'name'     : 'userIntent',
// 	'value'    : '{%22catering%22:{%22intent%22:false%2C%22lastIntentAction%22:%22%22%2C%22lastIntentTime%22:%22%22}}',
// 	'domain'   : '.publix.com',
// 	'path'     : '/',
// });
// phantom.addCookie({
// 	'name'     : 'bm_sv',
// 	'value'    : 'BB1B0A9EB1FD8305FC612FBEFFBA9606~YAAQLEsfRWmi+HKGAQAA44QJdhL5RRTNY1HT9bt577vtTijFfHZzL11c9DsUYJRWDeZLiSztaFGFQBBtgO/LLU3Fko6FYPTEYVjrEZ+/BZgWsjoGFA6iyRkTi0d4Q1MRoLwJxIWebvA7oqTKRQvRFGXOE6yIGgYFEi+cfUrXakJLNh+KopBD/QTZJ29TC6BRaevxezcKrFYHpglbrB1X8r1QdrCKP3VI4Vq+ot0QCO2bGX/K8S5s9yTQz6GzQ/7N~1',
// 	'domain'   : '.publix.com',
// 	'path'     : '/',
// });








// if (system.args.length === 1) {
//   phantom.exit(1);
// } else {
//   address = system.args[1];
//   page.open(address, function (status) {
//     console.log(page.content);
//     phantom.exit();
//   });
// }
// `
// 	res, _ := p.Exec(js, "https://www.publix.com/savings/weekly-ad/view-all")
// 	output, _ := ioutil.ReadAll(res)
// 	fmt.Println(string(output))

// 		/*
//     pTwo := &Param{
//         Method:       "GET",
//         Url:          "https://www.publix.com/savings/weekly-ad/view-all",
//         Header:       http.Header{"Cookie": []string{"your cookie"}},
//         UsePhantomJS: true,
//     }
//     brower := NewWebrowse()
//     resp, err := brower.Download(pTwo)
//     if err != nil {
//         fmt.Println(err)
//     }
//     body, err := ioutil.ReadAll(resp.Body)
//     fmt.Println(string(body))
//     fmt.Println(resp.Cookies())
// 	*/
// }

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"bytes"
)

func main() {
	client := &http.Client{}
	var data = []byte(`{
    "url": "https://www.publix.com/savings/weekly-ad/view-all",
    "type": "chrome",
    "ignoreHTTPStatusErrCodes": false,
    "proxy": "country-us",
    "waitDelay": 1,
    "initialCookies": [
        {
            "domain": ".publix.com",
            "expirationDate": 1711600801.924212,
            "hostOnly": false,
            "httpOnly": false,
            "name": "_4c_",
            "path": "/",
            "sameSite": "unspecified",
            "secure": false,
            "session": false,
            "storeId": "0",
            "value": "jVRtb9owEP4rKB%2F6qSF24rcgVROl1bSprOpopX1DwTZgEeLIMaRdxX%2FfGQJlbTUtQsJ399z5uecueY3apa6iAWacI4IEwnmKLqOVfmmiwWvkjAp%2F22gQ6RmXVOc6RjmjMeFSxgVNRUzwnEhMeMZkGl1Gz6GWIBgJkuaA3F1Gsu5qvEYbV0Kppfd1M0iStm379WZWmue%2BtOukKbamWjRJq%2FWqfIkLlWyNbuOiLKGutEpDKs77ok%2FB9r%2FBijOE4KyrQLB2Cs7348ef0%2Bvb4ej%2Bx9lNa%2B0d3HJ22SxpmuRg184mOPk%2BidN%2BmvVRfDf6NUkaTDKMUyZEluVUfBk%2BXF%2Fhi7VRV5QxiBHCGUsznnOSUYSyPAf1Mo5yQmhGSX4xfLi9wkDo63D69O0GuAjGsUgpEn0Qm2VZCgJB%2FKnRrndnF6bqTXzhQ5NP1aqybQXBibdO9ya61NJrFfqnLHQ80U1jLGTYjZMh5cY4gEAEmlEb6af%2BpQ7%2BVs96jVpBQOmtkXraGuWXoRAh6M271Gax9ODO94Kq2sE5jLM1lbLt%2B6zOe8oSaeh05mwLzYA9Wjq71j2MA9jCJkXjQsLR6bl2bg8BqzH7bt%2BG0vlg%2BU7u0FEYbhh5aWVRhgxdHQc9Hu2l%2Fb%2BRQNKjM4uFdmPtlzbI%2BegKZTwoWYQdU%2FAiAEelG7MIV6jQG1grb%2BuTe9ftOOc45bD3BDHYRg97LRhB4dkdOO9Xnv6NZox8gj4I948c%2FskNZYfGJzQioEHGKO3QYVgdemuOr7GYwSPmOEap4DFRRRbPCBXxPKcFkTJXQhTRWcmUISzIqaQ4VlR63pVMP9AVH%2BkGaQ8EzvTtvhUCcgn8BOBOPIt3cS4Qgbhrj20fAqJj9%2FzOs9vt%2FgA%3D",
            "id": 1
        },
        {
            "domain": ".publix.com",
            "expirationDate": 1708576704.883027,
            "hostOnly": false,
            "httpOnly": false,
            "name": "_abck",
            "path": "/",
            "sameSite": "unspecified",
            "secure": true,
            "session": false,
            "storeId": "0",
            "value": "6C9650C414A2AED8A380FB4FCEF91375~0~YAAQNew4F12oakSGAQAAuF1qdwkO62uDx7W8RZAGXXqXAV+oV1GfgZO1PDkNaHzDTBaVqplWKFmtN83utvIs9M0HelSjCxmLnlUxI4SeMowMonqnjP2/Lx0VWe5bwLHoxngTU3lB6sORdEmLSosTmdHW7XGBsnWbrzJO2pzTmuiWw2gqoHU8qNGeAUyBBZbx5UTuab9hByZQ/SHZDkBVAe5Jpa5mOVhFejiTzZnTtHCTQQllN38b9+f6o7EBxj3Qc6jT16Yr7+Pb+GQ9v1bIHXHjZhrz7edXwkm/+IYXUiInkGd60tPlwicuqbiJpbxq8cKKNU7Dqc1IXh5m2gA3w4j/lRxPC1gnf3elcfGNMRbf4FCFU2CgXt/X3ncE5oz1H5wyM7NbWXnmcmONG7mpZnaaEfV2YVMj~-1~-1~1677044285",
            "id": 2
        },
        {
            "domain": ".publix.com",
            "expirationDate": 1711600743.239398,
            "hostOnly": false,
            "httpOnly": false,
            "name": "_ga",
            "path": "/",
            "sameSite": "unspecified",
            "secure": false,
            "session": false,
            "storeId": "0",
            "value": "GA1.1.867182508.1676332429",
            "id": 3
        },
        {
            "domain": ".publix.com",
            "expirationDate": 1711600743.185042,
            "hostOnly": false,
            "httpOnly": false,
            "name": "_ga_1DWX30JN6C",
            "path": "/",
            "sameSite": "unspecified",
            "secure": false,
            "session": false,
            "storeId": "0",
            "value": "GS1.1.1677040701.10.1.1677040743.0.0.0",
            "id": 4
        },
        {
            "domain": ".publix.com",
            "expirationDate": 1711600743.237359,
            "hostOnly": false,
            "httpOnly": false,
            "name": "_ga_QZBH7YJQYK",
            "path": "/",
            "sameSite": "unspecified",
            "secure": false,
            "session": false,
            "storeId": "0",
            "value": "GS1.1.1677040701.10.1.1677040743.0.0.0",
            "id": 5
        },
        {
            "domain": ".publix.com",
            "expirationDate": 1710896066,
            "hostOnly": false,
            "httpOnly": false,
            "name": "_ga_T6YDMRV44X",
            "path": "/",
            "sameSite": "unspecified",
            "secure": false,
            "session": false,
            "storeId": "0",
            "value": "GS1.1.1676336052.1.1.1676336066.0.0.0",
            "id": 6
        },
        {
            "domain": ".publix.com",
            "expirationDate": 1684108427,
            "hostOnly": false,
            "httpOnly": false,
            "name": "_gcl_au",
            "path": "/",
            "sameSite": "unspecified",
            "secure": false,
            "session": false,
            "storeId": "0",
            "value": "1.1.1422772654.1676332427",
            "id": 7
        },
        {
            "domain": ".publix.com",
            "expirationDate": 1677127143,
            "hostOnly": false,
            "httpOnly": false,
            "name": "_gid",
            "path": "/",
            "sameSite": "unspecified",
            "secure": false,
            "session": false,
            "storeId": "0",
            "value": "GA1.2.892197158.1677012623",
            "id": 8
        },
        {
            "domain": ".publix.com",
            "expirationDate": 1708576745,
            "hostOnly": false,
            "httpOnly": false,
            "name": "_pin_unauth",
            "path": "/",
            "sameSite": "unspecified",
            "secure": false,
            "session": false,
            "storeId": "0",
            "value": "dWlkPU9URm1PVFl4TTJZdFpqSTBZaTAwTW1VeUxUaG1NVFV0WkRBM1lXWTFOemt6WldOag",
            "id": 9
        },
        {
            "domain": ".publix.com",
            "expirationDate": 1710892425,
            "hostOnly": false,
            "httpOnly": false,
            "name": "_ripgplsjne",
            "path": "/",
            "sameSite": "unspecified",
            "secure": true,
            "session": false,
            "storeId": "0",
            "value": "49662ea6-65ae-4d10-86b2-def022c3e40e",
            "id": 10
        },
        {
            "domain": ".publix.com",
            "expirationDate": 1677127143,
            "hostOnly": false,
            "httpOnly": false,
            "name": "_uetsid",
            "path": "/",
            "sameSite": "unspecified",
            "secure": false,
            "session": false,
            "storeId": "0",
            "value": "5e566b40b22911edb808c59be35e4fe4",
            "id": 11
        },
        {
            "domain": ".publix.com",
            "expirationDate": 1710736743,
            "hostOnly": false,
            "httpOnly": false,
            "name": "_uetvid",
            "path": "/",
            "sameSite": "unspecified",
            "secure": false,
            "session": false,
            "storeId": "0",
            "value": "9dde4b802ecb11edb7ee2f3e31713387",
            "id": 12
        },
        {
            "domain": ".publix.com",
            "expirationDate": 1677047896.79727,
            "hostOnly": false,
            "httpOnly": true,
            "name": "ak_bmsc",
            "path": "/",
            "sameSite": "unspecified",
            "secure": false,
            "session": false,
            "storeId": "0",
            "value": "1591DF7AB0406281C43158148AD3D9C8~000000000000000000000000000000~YAAQNew4F/qlakSGAQAAv01qdxIGhSExv/dTCuQ9mroV7TvSyMFv4VT9U+uTEVy72TTMlqO04iVdpQFe//GuIbFkwy9nQmZ2B5BT1S+WxTrq87bwMk8q3ZOtEluboxeIYIHLwHDFTLd17jOABLMfpsOmMRwtqUiKDbAojUVWLqhvGcwax/1w2na2ljZIXEv60AjQxudGgjIDg/juaQzEHycGanyq6JnwfbDKLzVStC8/VyoI1qmndCve86RfH7YwFzOt30NphsftfMLbFbWcJg5nA/aCElk3KWWFwJ+M9wHugd1CkRlZPfQxxB68YG00Sx1aXl8ASHwzdjI4YDtD45BLii5svPouBShBrarZB5FXQEMXkXWPEz2jcaBPgKHkkK4T9OwJONuMc6XXH70iR1mBXTIsDUg4D1qxqsncCfc/KerlqIgVXSdSTOyK6xJOfckKWBnM5IIDMaSn6Vuxz239tgW243cjfYOlaZkewSBltT6UsHj+AlU=",
            "id": 13
        },
        {
            "domain": ".publix.com",
            "expirationDate": 1711600703.639695,
            "hostOnly": false,
            "httpOnly": false,
            "name": "AMCV_0B25161260B4AFD00A495E9C%40AdobeOrg",
            "path": "/",
            "sameSite": "unspecified",
            "secure": false,
            "session": false,
            "storeId": "0",
            "value": "179643557%7CMCIDTS%7C19410%7CMCMID%7C56614344766237974350039940837094453549%7CMCAAMLH-1677645503%7C7%7CMCAAMB-1677645503%7CRKhpRz8krg2tLO6pguXWp5olkAcUniQYPHaMWWgdJ3xzPWQmdj0y%7CMCOPTOUT-1677047903s%7CNONE%7CMCAID%7CNONE%7CvVersion%7C5.5.0",
            "id": 14
        },
        {
            "domain": ".publix.com",
            "hostOnly": false,
            "httpOnly": false,
            "name": "AMCVS_0B25161260B4AFD00A495E9C%40AdobeOrg",
            "path": "/",
            "sameSite": "unspecified",
            "secure": false,
            "session": true,
            "storeId": "0",
            "value": "1",
            "id": 15
        },
        {
            "domain": ".publix.com",
            "expirationDate": 1677047904.90671,
            "hostOnly": false,
            "httpOnly": false,
            "name": "bm_sv",
            "path": "/",
            "sameSite": "unspecified",
            "secure": true,
            "session": false,
            "storeId": "0",
            "value": "367DE6BC04FDFCFEDBAD5240C1EC5DDD~YAAQNew4F5y9akSGAQAAuNpqdxL9oCJIe3ZYbnw8yvVnYbLslgC+YYIwd35wFmsNRoDxHVi1560IczMVecJN/SLEJa9ZEGdu9VOC5FoP9KFbRXjOanx9hkA6NmsZJyh/qum7bhLjYzLiu/II7xV+kjjhC6/09P1mUxGru97J34MU8Ybif/Ni8rXcfBANebuWoXmC3lvMrmNP3PAhsg0KAIrBF+WERrQiqkD/0qSI83nM1Ms0GToWa3XJjcm4FkbU~1",
            "id": 16
        },
        {
            "domain": ".publix.com",
            "expirationDate": 1677044326,
            "hostOnly": false,
            "httpOnly": false,
            "name": "bm_sz",
            "path": "/",
            "sameSite": "unspecified",
            "secure": false,
            "session": false,
            "storeId": "0",
            "value": "07DD1D112BD5DB112B8B532CE4A5A479~YAAQDqvbFxYMKFGGAQAA8OnFdhJhdKkFnilQ/GSpOZXAK38UFQXo+bmC+61i30h7aAuFr/bMDJ4iRT3+o6khrWYMZlxCshBQ2e/BuictVFbWX7bYIRjH5PfQFBOysgOgeXvUEZz9lXS7s+lqLnjQV4Al5lHk03cY1SRwyIgPTw40KcQBJgGhf4fryY/LCCFocYmHVaDCMhkF4Wg7pxEVMidxGMo5PrCuhZwQ8khFT3KGnNIaSIwG2sDBvK2CIRnE5ImBT5JqxZYAg5ZAz4zSc5mp3/1De/G4PkNaIk49A+yoTwo=~3425860~4274224",
            "id": 17
        },
        {
            "domain": ".publix.com",
            "hostOnly": false,
            "httpOnly": false,
            "name": "dtCookie",
            "path": "/",
            "sameSite": "unspecified",
            "secure": false,
            "session": true,
            "storeId": "0",
            "value": "v_4_srv_1_sn_A2FEC3E9DC8D0B2AD3EBA25A4961FDA3_perc_32485_ol_1_app-3Aa52023d6b6739bcd_0",
            "id": 18
        },
        {
            "domain": ".publix.com",
            "expirationDate": 1711600744.507761,
            "hostOnly": false,
            "httpOnly": true,
            "name": "FPID",
            "path": "/",
            "sameSite": "unspecified",
            "secure": true,
            "session": false,
            "storeId": "0",
            "value": "FPID2.2.Pxv5D6d6geydyDN0ey4xH99isAbm8vpEr2o6gmBMsFw%3D.1676332429",
            "id": 19
        },
        {
            "domain": ".publix.com",
            "expirationDate": 1677084628,
            "hostOnly": false,
            "httpOnly": false,
            "name": "FPLC",
            "path": "/",
            "sameSite": "unspecified",
            "secure": true,
            "session": false,
            "storeId": "0",
            "value": "gE%2BSmTS2vT0lS8k0Qp1wZzAUtHaZDXlq2o%2BJSlAe8Vqaxth5lZFAXqXGo8LACUKaPbe%2FAYKNWTH%2FGaCngXD3k8w6TZDL%2Fnjq73joDYuBHR8AdCTSPMLwCGtpNh5sJg%3D%3D",
            "id": 20
        },
        {
            "domain": ".publix.com",
            "hostOnly": false,
            "httpOnly": false,
            "name": "s_cc",
            "path": "/",
            "sameSite": "unspecified",
            "secure": false,
            "session": true,
            "storeId": "0",
            "value": "true",
            "id": 21
        },
        {
            "domain": ".publix.com",
            "expirationDate": 1711600742.980616,
            "hostOnly": false,
            "httpOnly": false,
            "name": "s_ecid",
            "path": "/",
            "sameSite": "unspecified",
            "secure": false,
            "session": false,
            "storeId": "0",
            "value": "MCMID%7C56614344766237974350039940837094453549",
            "id": 22
        },
        {
            "domain": ".publix.com",
            "hostOnly": false,
            "httpOnly": false,
            "name": "s_ips",
            "path": "/",
            "sameSite": "unspecified",
            "secure": false,
            "session": true,
            "storeId": "0",
            "value": "821",
            "id": 23
        },
        {
            "domain": ".publix.com",
            "hostOnly": false,
            "httpOnly": false,
            "name": "s_ppv",
            "path": "/",
            "sameSite": "unspecified",
            "secure": false,
            "session": true,
            "storeId": "0",
            "value": "%252Fsavings%252Fweekly-ad%252Fview-all%2C15%2C15%2C15%2C821%2C6%2C1",
            "id": 24
        },
        {
            "domain": ".publix.com",
            "hostOnly": false,
            "httpOnly": false,
            "name": "s_tp",
            "path": "/",
            "sameSite": "unspecified",
            "secure": false,
            "session": true,
            "storeId": "0",
            "value": "5532",
            "id": 25
        },
        {
            "domain": ".publix.com",
            "expirationDate": 1677041339,
            "hostOnly": false,
            "httpOnly": false,
            "name": "ShoppingListCount",
            "path": "/",
            "sameSite": "unspecified",
            "secure": false,
            "session": false,
            "storeId": "0",
            "value": "0",
            "id": 26
        },
        {
            "domain": ".publix.com",
            "expirationDate": 1692564613,
            "hostOnly": false,
            "httpOnly": false,
            "name": "Store",
            "path": "/",
            "sameSite": "unspecified",
            "secure": false,
            "session": false,
            "storeId": "0",
            "value": "%7B%22CreationDate%22%3A%222023-02-21T20%3A50%3A13.6406522%2B00%3A00%22%2C%22StoreName%22%3A%22Publix%20At%20University%20Village%20Market%22%2C%22StoreNumber%22%3A1560%2C%22Option%22%3A%22A%2CC%2CF%2CJ%2CN%2CO%2CR%2CS%2CT%2CV%2CY%22%2C%22ShortStoreName%22%3A%22University%20Vlg%20Mkt%22%2C%22ForceRefreshed%22%3Atrue%7D",
            "id": 27
        },
        {
            "domain": ".publix.com",
            "expirationDate": 1684108427,
            "hostOnly": false,
            "httpOnly": false,
            "name": "userIntent",
            "path": "/",
            "sameSite": "unspecified",
            "secure": false,
            "session": false,
            "storeId": "0",
            "value": "{%22catering%22:{%22intent%22:false%2C%22lastIntentAction%22:%22%22%2C%22lastIntentTime%22:%22%22}}",
            "id": 28
        },
        {
            "domain": ".www.publix.com",
            "expirationDate": 1677044297.165462,
            "hostOnly": false,
            "httpOnly": true,
            "name": "TiPMix",
            "path": "/",
            "sameSite": "no_restriction",
            "secure": true,
            "session": false,
            "storeId": "0",
            "value": "78.45588263342414",
            "id": 29
        },
        {
            "domain": ".www.publix.com",
            "expirationDate": 1677044297.165526,
            "hostOnly": false,
            "httpOnly": true,
            "name": "x-ms-routing-name",
            "path": "/",
            "sameSite": "no_restriction",
            "secure": true,
            "session": false,
            "storeId": "0",
            "value": "self",
            "id": 30
        },
        {
            "domain": "www.publix.com",
            "hostOnly": true,
            "httpOnly": true,
            "name": "akavpau_VP_WWW_PRD",
            "path": "/",
            "sameSite": "no_restriction",
            "secure": true,
            "session": true,
            "storeId": "0",
            "value": "1677041336~id=3c866912637c4579f410cda5f640d92d",
            "id": 31
        },
        {
            "domain": "www.publix.com",
            "hostOnly": true,
            "httpOnly": true,
            "name": "ASP.NET_SessionId",
            "path": "/",
            "sameSite": "unspecified",
            "secure": true,
            "session": true,
            "storeId": "0",
            "value": "kjd0brg1jlxqbmowtftj1e5x",
            "id": 32
        },
        {
            "domain": "www.publix.com",
            "hostOnly": true,
            "httpOnly": false,
            "name": "store_rpt",
            "path": "/",
            "sameSite": "unspecified",
            "secure": false,
            "session": true,
            "storeId": "0",
            "value": "1",
            "id": 33
        }
    ],
    "actions": [
        {
            "jsclick": {
                "selector": ".button"
            }
        }
    ]
}`)
	req, err := http.NewRequest("POST", "https://api.dataflowkit.com/v1/fetch?api_key=34ffe1c536505c5168ae76212d27861cd67457cc4c8a8ff88aaca7c48e2401dc", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}

