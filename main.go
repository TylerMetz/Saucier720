/*
package main

import (
	"BackendPkg"
	//"fmt"
	//"time"
)


func main() {
		fmt.Println("Welcome to out Sprint 1 demo!")
		userPantry := BackendPkg.Pantry{
			TimeLastUpdated: time.Now(),
		}
		userPantry.AddToPantry()
		userPantry.AddToPantry()
		userPantry.AddToPantry()
		userPantry.DisplayPantry()
		userPantry.RemoveFromPantry()
		userPantry.RemoveFromPantry()
		userPantry.DisplayPantry()
}
*/

package main

import (
	"fmt"
	"github.com/k4s/phantomgo"
	"io/ioutil"
)

func main() {
	p := phantomgo.NewPhantom()

	js := `
var page = require('webpage').create(),
  system = require('system'),
  address;
page.settings.userAgent = 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/37.0.2062.120 Safari/537.36';
phantom.cookiesEnabled = true;
phantom.addCookie({
	'name'     : 'store_rpt',
	'value'    : '1',
	'domain'   : 'www.publix.com',
	'path'     : '/',
});
phantom.addCookie({
	'name'     : 'x-ms-routing-name',
	'value'    : 'self',
	'domain'   : '.www.publix.com',
	'path'     : '/',
});
phantom.addCookie({
	'name'     : 'TiPMix',
	'value'    : '79.1668783049623',
	'domain'   : '.www.publix.com',
	'path'     : '/',
});
phantom.addCookie({
	'name'     : '_4c_',
	'value'    : 'jVPfT9swEP5XKj%2FwRBL%2FjO1KaCoFTZtgiHVIe6uc5KBR0yRy0gaG%2BN85tykgJk3Lg3X3%2BbvzfXeXZzKsoCZTlmpNmVbGUkNPyRqeOjJ9Jnkbzl04tr4iU7Lq%2B7abJskwDHG7zaryMc6bTdK5XVk%2FdMkAsK6eIlckuxKGyFUVOSV5UwCGMhubWKHf%2F0EvEpSiDTUmJ60v0L65%2FvVzeX45m9%2F8%2BPDSBnqPr3x4LEu6Ljn4rW8SlnxfRDzmIqbR1fz3IumsZkwIprmVwpgvs9vzM3ayKYszxqgVhqVCyFRZbihNuaCccaq0NVZrRblgJ7PbyzOGBX2dLe%2B%2BXYTSpZDcIk3Gx0YpjYS7Dvzkqnko68mid31QeVev62ao8XLRNx4mC6gg76EIWVQaJC%2Bg68oGI5qtz0PIRemRgjeoptjm%2FbJ%2FagM%2BQDbpijVeFLArc1gOZdGv9p3k9B1dQfmw6gOMowtw64OD1lDWRTN8DhvRD2E0kDPfDCgHgfnKNxuYYLMQbnAPyLXL0fRwD97vKeh15V7v%2B1xGDFfnDQ6awnw5GlWTuypEQH2c9fX80N3%2Fmgp5OSWPh0VlVFhjpNG4Sz1upUklDR8yfFmMG0sk5xJMkUegcxnJVOSRBVpE1jJ2D2CdcEDGnEZqK5QyqQ1JduUxR0qZc7l0EXWZiqRTIjJC3EcZVzpzKhVq37vH4w%2BEOjSTcqyLmWNZbTVmZO9kwawy2pgjWb6JaHcjm3%2BSTLn5W%2FJhcP%2BIsZ9jXl5eAQ%3D%3D',
	'domain'   : '.publix.com',
	'path'     : '/',
});
phantom.addCookie({
	'name'     : 's_tp',
	'value'    : '5652',
	'domain'   : '.publix.com',
	'path'     : '/',
});
phantom.addCookie({
	'name'     : '_uetvid',
	'value'    : 'd9cb4cb0b23411ed82695bfd593c245b',
	'domain'   : '.publix.com',
	'path'     : '/',
});
phantom.addCookie({
	'name'     : '_uetsid',
	'value'    : 'd9c9d500b23411ed80d2effd1bdedfaa',
	'domain'   : '.publix.com',
	'path'     : '/',
});
phantom.addCookie({
	'name'     : '_ga',
	'value'    : 'GA1.1.1434291204.1677017557',
	'domain'   : '.publix.com',
	'path'     : '/',
});
phantom.addCookie({
	'name'     : 'Store',
	'value'    : '{%22CreationDate%22:%222023-02-21T22:12:54.050Z%22%2C%22ForceRefreshed%22:false%2C%22Option%22:%22ACFJNORSTVY%22%2C%22ShortStoreName%22:%22University%20Vlg%20Mkt%22%2C%22StoreName%22:%22Publix%20At%20University%20Village%20Market%22%2C%22StoreNumber%22:1560}',
	'domain'   : '.publix.com',
	'path'     : '/',
});
phantom.addCookie({
	'name'     : 'ShoppingListCount',
	'value'    : '0',
	'domain'   : '.publix.com',
	'path'     : '/',
});
phantom.addCookie({
	'name'     : '_gid',
	'value'    : 'GA1.2.1549355166.1677017557',
	'domain'   : '.publix.com',
	'path'     : '/',
});
phantom.addCookie({
	'name'     : 'FPLC',
	'value'    : 'gMINW4sevUMps3lQLnQOa8%2F%2F9dNaED8k%2FtZyxF8pSmTrXmNTEKAF%2BAZwZx2n%2F0g2FNMKNGwOHQ300%2Fo5euhQAd6izBJakB2zzvvmoyj6Adj6Jy0nEQZOa9ezuRaGig%3D%3D',
	'domain'   : '.publix.com',
	'path'     : '/',
});
phantom.addCookie({
	'name'     : 's_ips',
	'value'    : '1001',
	'domain'   : '.publix.com',
	'path'     : '/',
});
phantom.addCookie({
	'name'     : '_ga_1DWX30JN6C',
	'value'    : 'GS1.1.1677017557.1.1.1677017582.0.0.0',
	'domain'   : '.publix.com',
	'path'     : '/',
});
phantom.addCookie({
	'name'     : '_fbp',
	'value'    : 'fb.1.1677017558246.1302144086',
	'domain'   : '.publix.com',
	'path'     : '/',
});
phantom.addCookie({
	'name'     : '_gcl_au',
	'value'    : '1.1.929578126.1677017556',
	'domain'   : '.publix.com',
	'path'     : '/',
});
phantom.addCookie({
	'name'     : 's_ecid',
	'value'    : 'MCMID%7C11093816334659280062302120579897750231',
	'domain'   : '.publix.com',
	'path'     : '/',
});
phantom.addCookie({
	'name'     : 'AMCVS_0B25161260B4AFD00A495E9C%40AdobeOrg',
	'value'    : '1',
	'domain'   : '.publix.com',
	'path'     : '/',
});
phantom.addCookie({
	'name'     : 'akavpau_VP_WWW_PRD',
	'value'    : '1677018188~id=198b0259def670900dbcff3ef15d5b5c',
	'domain'   : 'www.publix.com',
	'path'     : '/',
});
phantom.addCookie({
	'name'     : '_ripgplsjne',
	'value'    : '9acb25b3-2ad9-4213-a223-6c103788d750',
	'domain'   : '.publix.com',
	'path'     : '/',
});
phantom.addCookie({
	'name'     : 'FPID',
	'value'    : 'FPID2.2.A3pExBAnE%2FAoXrtCZGosG9nMPihNHtwCxDDJuWcwvb8%3D.1677017557',
	'domain'   : '.publix.com',
	'path'     : '/',
});
phantom.addCookie({
	'name'     : 's_cc',
	'value'    : 'true',
	'domain'   : '.publix.com',
	'path'     : '/',
});
phantom.addCookie({
	'name'     : '_pin_unauth',
	'value'    : 'dWlkPU16YzROREV5TURjdE1EZzNNUzAwTVdOa0xXSTFNek10WXpjd1kyWmlZbUZoTnpNMA',
	'domain'   : '.publix.com',
	'path'     : '/',
});
phantom.addCookie({
	'name'     : '_abck',
	'value'    : '10B92E36789FEE071DA625AE763092F6~0~YAAQLEsfRfGf+HKGAQAADyEJdgmv/uMU5VcimdNTufUHDL73+dJB1zNIXgadLOLykRGUVAVNndz/Gp8+HgAQ31bN2nWZjp1Kpi3prRRfgXxS5sIR2FuO3Yya9eFJLWKUal4wMqKn7WltcFZDREGpXXRhUmM1O6SIiXsGLplWM3y7DelsgINLEfMO6619WcC2v8xiQQSPm/aTwiqOlrFPJxcpcJgaEN+CkzJpXFEr0F2zZa/+UQPQyxMnS///5owtZmdP7ETn0Trfc7G3Ar2lwc6NkY1YW7wr0oB3OfQSNvqaD3Re6MD2QSwkdzCfm1cu++JqUkRPHohn0pXTW8snv1wV1+I/GQikbae44tXL/MFvbJDRCoHRs142IQeinkYzjxlej9nK+T8nASDv4/DMi1tbL2upzAQZ~-1~||1-LMXiTZORee-1-10-1000-2||~1677021111',
	'domain'   : '.publix.com',
	'path'     : '/',
});
phantom.addCookie({
	'name'     : '_dc_gtm_UA-34592858-5',
	'value'    : '1',
	'domain'   : '.publix.com',
	'path'     : '/',
});
phantom.addCookie({
	'name'     : 'AMCV_0B25161260B4AFD00A495E9C%40AdobeOrg',
	'value'    : '179643557%7CMCIDTS%7C19410%7CMCMID%7C11093816334659280062302120579897750231%7CMCAID%7CNONE%7CMCOPTOUT-1677024758s%7CNONE%7CMCAAMLH-1677622358%7C7%7CMCAAMB-1677622358%7Cj8Odv6LonN4r3an7LhD3WZrU1bUpAkFkkiY1ncBR96t2PTI%7CMCSYNCSOP%7C411-19417%7CvVersion%7C5.5.0',
	'domain'   : '.publix.com',
	'path'     : '/',
});
phantom.addCookie({
	'name'     : '_ga_QZBH7YJQYK',
	'value'    : 'GS1.1.1677017557.1.1.1677017582.0.0.0',
	'domain'   : '.publix.com',
	'path'     : '/',
});
phantom.addCookie({
	'name'     : 'ak_bmsc',
	'value'    : '88787A437F34446F12F4D42ADC0CB02E~000000000000000000000000000000~YAAQLEsfRfKf+HKGAQAAnSEJdhLEeAEjAacM9GrOIpZA8Mud0pou2tQK+yowuXKRVFT3v/6KNU5jFaOXkhGx8ro6JZiVbdkIgO/vtGXhY7tnhKqHMa2xAT6j0bEQY0/VGU/syBvHxz2zsSzHuJVeySH1XLm7+6iwY88uSMoLDfHSx0oTuTA06qCW3NLwcWJKEUeyr+NmSP1VqxAHg9CT+Ar7JQIc9kI1G6k5lPVYBfjpEzXCUzJb/7FhhKT9GBskgZfWMfsOj+ls1th2G6r6OfU8blLLrcFtu6O/SIjG0dNWW90+RWRzklJ5ct4yMbyPEPcUnKu/kxmpQoEDIMPN+yVcetfqDvmfytU3KQUrCrIHnpO2Ro2594hZvKG024kxJb0Ikf6A/b6F7MGaxdU18w+EwUcPBlTLmA25lSxaCDljPYXHIxGAsxSJ3Bt7KlB+SgHoIpeiRFS6yUftNwJCU/ZbRbK680mLAoLJm3wDlQ5I3VS7EpSOjzU=',
	'domain'   : '.publix.com',
	'path'     : '/',
});
phantom.addCookie({
	'name'     : 'bm_sz',
	'value'    : '4A3BC2CD5F582F404CFB0F6250041661~YAAQLEsfRbaf+HKGAQAACBQJdhLj16tkY63cJlQmVon+EUrqwKi3KlU3EzFBo/306zDp+oOvdWKsuRI4/okWNzC+h2XgFV0iK36EZTYTnFWlYCoH+2TdGGzIIBFyzYFWgkPK/To12WqUx0unp64CFh3QKGAKIenwX536kwSCc7zlhkm3xamMBAzPvh26ZElrRsMfCmeWDt/1h9lz98a5BsVcPQmGXQ0r0Z0TpuHTOtqM09YbSGnf5wU7l2j0AsWgpK8zMbMfbwYYQIhLYfAdhlJAhW837cWgjOWMY9S8wnPcBL0=~4535617~3687220',
	'domain'   : '.publix.com',
	'path'     : '/',
});
phantom.addCookie({
	'name'     : 's_ppv',
	'value'    : '%252Fsavings%252Fweekly-ad%252Fview-all%2C18%2C18%2C18%2C1001%2C5%2C1',
	'domain'   : '.publix.com',
	'path'     : '/',
});
phantom.addCookie({
	'name'     : 'dtCookie',
	'value'    : 'v_4_srv_9_sn_159C4F9C44085C83150A874C76F3320F_perc_27704_ol_1_app-3Aa52023d6b6739bcd_0',
	'domain'   : '.publix.com',
	'path'     : '/',
});
phantom.addCookie({
	'name'     : 'userIntent',
	'value'    : '{%22catering%22:{%22intent%22:false%2C%22lastIntentAction%22:%22%22%2C%22lastIntentTime%22:%22%22}}',
	'domain'   : '.publix.com',
	'path'     : '/',
});
phantom.addCookie({
	'name'     : 'bm_sv',
	'value'    : 'BB1B0A9EB1FD8305FC612FBEFFBA9606~YAAQLEsfRWmi+HKGAQAA44QJdhL5RRTNY1HT9bt577vtTijFfHZzL11c9DsUYJRWDeZLiSztaFGFQBBtgO/LLU3Fko6FYPTEYVjrEZ+/BZgWsjoGFA6iyRkTi0d4Q1MRoLwJxIWebvA7oqTKRQvRFGXOE6yIGgYFEi+cfUrXakJLNh+KopBD/QTZJ29TC6BRaevxezcKrFYHpglbrB1X8r1QdrCKP3VI4Vq+ot0QCO2bGX/K8S5s9yTQz6GzQ/7N~1',
	'domain'   : '.publix.com',
	'path'     : '/',
});




phantom.addCookie({
	'name'     : 'FIRST_PAGE_VISIT',
	'value'    : 'false',
});
phantom.addCookie({
	'name'     : 'OPTLY_PILOT_USER_ID',
	'value'    : '4ca5b7ce-8890-446c-a1b2-c4f202465a76',
});
phantom.addCookie({
	'name'     : 'geolocated_store_selection-test-4ca5b7ce-8890-446c-a1b2-c4f202465a76',
	'value'    : 'null',
});
phantom.addCookie({
	'name'     : 'search_product_card_aisle_location-test-4ca5b7ce-8890-446c-a1b2-c4f202465a76',
	'value'    : 'null',
});
phantom.addCookie({
	'name'     : 'savings_switch_tpr_pricing_location-4ca5b7ce-8890-446c-a1b2-c4f202465a76',
	'value'    : 'true',
});
phantom.addCookie({
	'name'     : 'search_single_click_add_to_cart-4ca5b7ce-8890-446c-a1b2-c4f202465a76',
	'value'    : 'true',
});
phantom.addCookie({
	'name'     : 'geolocated_store_selection-4ca5b7ce-8890-446c-a1b2-c4f202465a76',
	'value'    : 'true',
});
phantom.addCookie({
	'name'     : 'search_product_card_aisle_location-4ca5b7ce-8890-446c-a1b2-c4f202465a76',
	'value'    : 'true',
});
phantom.addCookie({
	'name'     : 'pbxStorStoreConfirmationModalConfirmed',
	'value'    : '{"data":false,"expires":null}',
});
phantom.addCookie({
	'name'     : 'com.adobe.reactor.core.visitorTracking.pagesViewed',
	'value'    : '3',
});
phantom.addCookie({
	'name'     : 'terms_of_use_prompt-test-4ca5b7ce-8890-446c-a1b2-c4f202465a76',
	'value'    : 'null',
});
phantom.addCookie({
	'name'     : 'getitemsbytype_all_pages-test-4ca5b7ce-8890-446c-a1b2-c4f202465a76',
	'value'    : 'null',
});
phantom.addCookie({
	'name'     : 'OPTLY_USER_ID',
	'value'    : '4ca5b7ce-8890-446c-a1b2-c4f202465a76',
});
phantom.addCookie({
	'name'     : 'search_single_click_add_to_cart-test-4ca5b7ce-8890-446c-a1b2-c4f202465a76',
	'value'    : 'null',
});
phantom.addCookie({
	'name'     : 'search_catering_feature-test-4ca5b7ce-8890-446c-a1b2-c4f202465a76',
	'value'    : 'null',
});
phantom.addCookie({
	'name'     : 'terms_of_use_prompt-4ca5b7ce-8890-446c-a1b2-c4f202465a76',
	'value'    : 'false',
});
phantom.addCookie({
	'name'     : 'search_catering_feature-4ca5b7ce-8890-446c-a1b2-c4f202465a76',
	'value'    : 'true',
});
phantom.addCookie({
	'name'     : 'savings_use_getlistitemsbytype-test-4ca5b7ce-8890-446c-a1b2-c4f202465a76',
	'value'    : 'null',
});
phantom.addCookie({
	'name'     : 'getitemsbytype_all_pages-4ca5b7ce-8890-446c-a1b2-c4f202465a76',
	'value'    : 'true',
});
phantom.addCookie({
	'name'     : 'dtDisabled',
	'value'    : 'true',
});
phantom.addCookie({
	'name'     : 'savings_switch_tpr_pricing_location-test-4ca5b7ce-8890-446c-a1b2-c4f202465a76',
	'value'    : 'null',
});
phantom.addCookie({
	'name'     : 'com.adobe.reactor.core.visitorTracking.trafficSource',
	'value'    : 'false',
});
phantom.addCookie({
	'name'     : 'pbxStor/v4/savings{"method":"get","data":null,"params":{"smImg":235,"enImg":368,"fallbackImg":false,"isMobile":false,"page":1,"pageSize":0,"includePersonalizedDeals":true},"headers":{"PublixStore":1560}}',
	'value'    : 'data:{"IsPersonalizationEnabled": false, "PersonalizedStoreNumber": 1560, "PersonalizedStoreName": null,…}',
});
phantom.addCookie({
	'name'     : 'savings_use_getlistitemsbytype-4ca5b7ce-8890-446c-a1b2-c4f202465a76',
	'value'    : 'true',
});
phantom.addCookie({
	'name'     : 'sl_count_cache_updates_enabled-test-4ca5b7ce-8890-446c-a1b2-c4f202465a76',
	'value'    : 'null',
});
phantom.addCookie({
	'name'     : 'com.adobe.reactor.core.visitorTracking.landingTime',
	'value'    : '1677020068569',
});
phantom.addCookie({
	'name'     : 'ANTI_FORGERY_TOKEN',
	'value'    : 'ZZON4tomepo7vcryOOLAC6lV_K2CkQJCbRaYTwOWQsD_gpC31vuopOF3IBVYzh_ybqIz5bEVdiaYzOxFGbFn7ztDPvE1:A5IL5MPm6fJRfunXV6MEMjlhvPj3QHB8DsqdQBBOcX_KoMpNAlLf6xF8SdmarHoRFw772KhmZ1c7nCmQODBWybHGU8U1',
});
phantom.addCookie({
	'name'     : 'sl_count_cache_updates_enabled-4ca5b7ce-8890-446c-a1b2-c4f202465a76',
	'value'    : 'true',
});
phantom.addCookie({
	'name'     : 'header_navigation_bar-test-4ca5b7ce-8890-446c-a1b2-c4f202465a76',
	'value'    : 'null',
});
phantom.addCookie({
	'name'     : 'header_navigation_bar-4ca5b7ce-8890-446c-a1b2-c4f202465a76',
	'value'    : 'false',
});
phantom.addCookie({
	'name'     : 'search_card_addtolist_v2-4ca5b7ce-8890-446c-a1b2-c4f202465a76',
	'value'    : 'false',
});
phantom.addCookie({
	'name'     : 'pbxStordisableGeoButton',
	'value'    : '{"data":true,"expires":null}',
});
phantom.addCookie({
	'name'     : 'search_card_addtolist_v2-test-4ca5b7ce-8890-446c-a1b2-c4f202465a76',
	'value'    : 'null',
});
phantom.addCookie({
	'name'     : 'com.adobe.reactor.core.visitorTracking.landingPage',
	'value'    : 'https://www.publix.com/savings/weekly-ad/view-all',
});
phantom.addCookie({
	'name'     : 'is_eu',
	'value'    : 'false',
});





if (system.args.length === 1) {
  phantom.exit(1);
} else {
  address = system.args[1];
  page.open(address, function (status) {
    console.log(page.content);
    phantom.exit();
  });
}
`
	res, _ := p.Exec(js, "https://www.publix.com/savings/weekly-ad/view-all")
	output, _ := ioutil.ReadAll(res)
	fmt.Println(string(output))

}

