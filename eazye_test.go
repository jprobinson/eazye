package eazye

import (
	"bytes"
	"net/mail"
	"testing"
)

func TestParseQuotedBody(t *testing.T) {

	header := mail.Header(map[string][]string{
		"Content-Type": []string{"text/html; charset=\"iso-8859-1\""},
		"Content-Transfer-Encoding": []string{"quoted-printable"},
	})
	html, text, multipart, err := parseBody(header, []byte(quotedEmail))
	if err != nil {
		t.Error("error doing test! ", err)
		return
	}

	if multipart {
		t.Error("parse body returned with multipart flag set")
	}

	if len(text) > 0 {
		t.Errorf("an HTML only email returned with text: %q", string(text))
	}

	if bytes.Contains(html, []byte("=3D")) {
		t.Errorf("parseBody did not handle quoted-printable format as expected. `=3D' still found in message:\n%q", string(html))
	}

	// crazy Äpple to verify we parsed iso-8859-1 correctly AND did the quoted-printable fix
	if !bytes.HasPrefix(html, []byte("Äpple<html><head>")) || !bytes.HasSuffix(html, []byte("</html>\r\n")) {
		t.Errorf("parseBody did not pull out HTML correctly. Expected '<html><head>' prefix and '</html>\r\n' suffix. got:\n%q", string(html))
	}
}

func TestParsePlainBody(t *testing.T) {
	header := mail.Header(map[string][]string{"Content-Type": []string{"text/html; charset=UTF-8"}})
	html, text, multipart, err := parseBody(header, []byte(htmlEmail))
	if err != nil {
		t.Error("error doing test! ", err)
		return
	}

	if multipart {
		t.Error("parse body returned with multipart flag set")
	}

	if len(text) > 0 {
		t.Errorf("an HTML only email returned with text: %q", string(text))
	}

	if !bytes.HasPrefix(html, []byte("<br />")) || !bytes.HasSuffix(html, []byte("FoxNews.com.")) {
		t.Errorf("parseBody did not pull out HTML correctly. Expected '<br />' prefix and 'FoxNews.com.' suffix. got:\n%q", string(html))
	}
}

func TestParseMultiBody(t *testing.T) {
	header := mail.Header(map[string][]string{"Content-Type": []string{`multipart/alternative; boundary="eZDakj4l4DVQ=_?:"`}})
	html, text, multipart, err := parseBody(header, []byte(multipartEmail))
	if err != nil {
		t.Error("error doing test! ", err)
		return
	}

	if !multipart {
		t.Error("parse body did not return with multipart flag set")
	}

	if !bytes.HasPrefix(html, []byte("<!DOCTYPE html>")) || !bytes.HasSuffix(html, []byte("</html>\n")) {
		t.Errorf("parseBody did not pull out the HTML correctly. Expected to start with doctype and end with </html>, got:\n%q", string(html))
	}

	if !bytes.HasPrefix(text, []byte("NBA announces")) || !bytes.HasSuffix(text, []byte("approved Ballmer.\n")) {
		t.Errorf("parseBody did not pull out text correctly. Expected 'NBA announces' prefix and 'Ballmer.' suffix. got:\n%q", string(text))
	}
}

func TestParseSubject(t *testing.T) {
	tests := []struct {
		given string
		want  string
	}{
		{
			"A plain string subject",
			"A plain string subject",
		},
		{
			"=?UTF-8?B?SmFwYW7igJlzIEVjb25vbXkgU2hyaW5rcyB0aGUgTW9zdCBTaW5j?= =?UTF-8?B?ZSAyMDExIFF1YWtlIG9uIFNhbGVzIFRheA==?=",
			"Japan’s Economy Shrinks the Most Since 2011 Quake on Sales Tax",
		},
	}

	for _, test := range tests {
		got := parseSubject(test.given)
		if got != test.want {
			t.Errorf("parseSubject(%s) got:%s want:%s", test.given, got, test.want)
		}
	}

}

func TestVisibleText(t *testing.T) {
	email := Email{
		HTML: []byte(`<html>
<head>
 <title></title>
 <style type="text/css">a {color:#004276;text-decoration:none;}
        h6 {font-size: 18px; margin:0;}
 </style>
</head>
<body>
<div align="center">
<table border="0" cellpadding="0" cellspacing="0" style="margin:0 45px;" width="500">
 <tbody>
  <tr>
   <td align="center" style="padding-bottom:11px; font-family: arial,helvetica,sans serif; font-size:11px; color:#666">To ensure delivery to your inbox, please add <a href="http://p.nytimes.com/email/re?location=4z5Q7LhI+KVBjmEgFdYACHCpQxp3Fo5Z&amp;campaign_id=132&amp;instance_id=45669&amp;segment_id=62821&amp;user_id=104ac336ab7205252982f8c4bea84287&amp;regi_id=66509827" style="color:#004276; text-decoration:none;">nytdirect@nytimes.com</a> to your address book.</td>
  </tr>
 </tbody>
</table>

<div id="NYTEmail" style="margin:0 45px; width:500px;"><!--Header Begins-->
<table border="0" cellpadding="0" cellspacing="0" id="customHeader" width="500">
 <tbody>
  <tr>
   <td>
   <table border="0" cellpadding="0" cellspacing="0" valign="bottom" width="500">
    <tbody>
     <tr>
      <td id="MainHeader" style="font-size: 14px; padding: 26px 0 4px 0;" width="350">
      <table border="0" style="font-size: 14px;">
       <tbody>
        <tr>
         <td><a href="http://p.nytimes.com/email/re?location=4z5Q7LhI+KVBjmEgFdYACKqxdqYf5TwZRlcqaJDmNgZR6U1YR0ohawHboqRm57jG&amp;campaign_id=132&amp;instance_id=45669&amp;segment_id=62821&amp;user_id=104ac336ab7205252982f8c4bea84287&amp;regi_id=66509827" style="color: #004276; text-decoration: none;"><img alt="The New York Times" border="0" height="19" src="http://graphics8.nytimes.com/images/logos/nyt/nyt-logo-122x18.png" style="vertical-align: bottom;" width="122" /> </a></td>
         <td align="right" width="5"><span style="color: #000001; font-size: 14px; line-height: 18px; vertical-align: top;">|</span></td>
         <td><a href="http://p.nytimes.com/email/re?location=4z5Q7LhI+KVBjmEgFdYACKqxdqYf5TwZRlcqaJDmNgZR6U1YR0ohawHboqRm57jG&amp;campaign_id=132&amp;instance_id=45669&amp;segment_id=62821&amp;user_id=104ac336ab7205252982f8c4bea84287&amp;regi_id=66509827" style="color: #000001; line-height: 18px; font-family: arial,helvetica,sans serif; font-size: 14px; font-weight: bold; text-decoration: none;">BREAKING NEWS ALERT</a></td>
        </tr>
       </tbody>
      </table>
      </td>
      <td align="right" id="subNavigation" style=" font-size:11px; padding:26px 0 4px 0;" width="160">
      <table border="0" cellpadding="0" cellspacing="0" style="font-size: 11px; -webkit-text-size-adjust: none;" width="160">
       <tbody>
        <tr>
         <td align="center" width="90"><a href="http://p.nytimes.com/email/re?location=4z5Q7LhI+KVBjmEgFdYACKqxdqYf5TwZRlcqaJDmNgZR6U1YR0ohawHboqRm57jG&amp;campaign_id=132&amp;instance_id=45669&amp;segment_id=62821&amp;user_id=104ac336ab7205252982f8c4bea84287&amp;regi_id=66509827" style="color: #808080; font-family: arial,helvetica,sans serif; font-size:11px; text-decoration:none;">NYTimes.com </a></td>
         <td align="center" width="2"><span style=" color: #808080; font-size: 11px;">|</span></td>
         <td align="center" width="88"><span><a href="http://www.nytimes.com/gst/unsub.html?email=@gmail.com&id=66509827&segment=62821&group=nl&product=NA" style="color: #808080; font-family: arial,helvetica,sans serif; font-size:11px; text-decoration:none;"></a></span></td>
        </tr>
       </tbody>
      </table>
      </td>
     </tr>
    </tbody>
   </table>
   </td>
  </tr>
 </tbody>
</table>

<table cellpadding="0" cellspacing="0" id="headerExtra" width="500">
 <tbody>
  <tr>
   <td style="border-top: 4px solid #000001; padding-top:15px;">
   <table style="text-align:left; -webkit-text-size-adjust: none;">
    <tbody>
     <tr>
      <td align="left"><span style="color:#A81817; font-family: arial,helvetica,sans-serif; font-size: 11px; font-weight:bold;">BREAKING NEWS</span></td>
      <td align="left"><span style="color:#666; font-family: arial,helvetica,sans-serif; font-size: 11px;">Sunday, August 24, 2014 2:11 PM EDT</span></td>
     </tr>
    </tbody>
   </table>
   </td>
  </tr>
 </tbody>
</table>
<!--Header Ends--><!--aColumn Begins-->

<table cellpadding="0" cellspacing="0" id="AssetBody" width="500">
 <tbody>
  <tr>
   <td style="color:#333; font-family: georgia,serif;" valign="top">
   <table cellpadding="0" cellspacing="0">
    <tbody>
     <tr>
      <td style="padding-bottom:11px; width: 490px !important;">
      <table cellpadding="0" cellspacing="0" style="font-size: 15px; line-height:22px; margin:0; -webkit-text-size-adjust: none;">
       <tbody>
        <tr>
         <td><a href="http://p.nytimes.com/email/re?location=4z5Q7LhI+KVBjmEgFdYACPLKh239P3pgQLjG2m3jp+WcpvlGpgZEKT3q73Mn5RDyQ5ptpoEAUoNNFw7jyLJpkQCGONPVoIcAR6NTKX3jPD/Y6M61NUtbzP4m8j4mnAa/uJz6slrI6OQsDpEjqlyV7HcTqxgI9zdnbEjTRgatCSYHfrMImtDDEhNsAeYWRR8Df2Feo8Cua0s=&amp;campaign_id=132&amp;instance_id=45669&amp;segment_id=62821&amp;user_id=104ac336ab7205252982f8c4bea84287&amp;regi_id=66509827" style="color:#000001; font-size: 24px; line-height:30px; text-align: left; text-decoration:none;">U.S. Journalist Held by Qaeda Affiliate in Syria Is Freed After Nearly 2 Years</a>

         <table cellpadding="0" cellspacing="0" id="summary" style="color: #333; line-height:22px; font-size: 15px; margin-top: 11px;">
          <tbody><tr><td style="color: #333; line-height:22px; font-size: 15px;  padding:4px 0 12px 0;">An American journalist held captive for nearly two years by Al Qaeda&#8217;s
official branch in Syria has been freed, according to a representative of the
journalist&#8217;s family and a report on Sunday by the Al Jazeera
network.</td></tr>
<tr><td style="color: #333; line-height:22px; font-size: 15px;  padding:4px 0 12px 0;">The journalist, Peter Theo Curtis, was abducted near the
Syria-Turkey border in October 2012. He was held by the Nusra Front, the Qaeda
affiliate in Syria, which has broken with the more radical Islamic State in Iraq
and Syria, or ISIS. Another American journalist, James W. Foley, who was
kidnapped in Syria the following month, was beheaded last week by ISIS, which
posted images of his execution on YouTube.</td></tr>
<tr><td style="color: #333; line-height:22px; font-size: 15px;  padding:4px 0 12px 0;">A family friend confirmed on
Sunday that Mr. Curtis, originally from Boston, had been handed over to a United
Nations representative.</td></tr>

           <tr>
            <td style="color:#000001; font-family: arial,helvetica,sans serif; font-size: 11px; padding:3px 0 11px 0;">
            <h4 style="color:#000001; font-weight: bold; margin: 5px 0 2px 0;">READ MORE &raquo;</h4>
            <a href="http://p.nytimes.com/email/re?location=4z5Q7LhI+KVBjmEgFdYACPLKh239P3pgQLjG2m3jp+WcpvlGpgZEKT3q73Mn5RDyQ5ptpoEAUoNNFw7jyLJpkQCGONPVoIcAR6NTKX3jPD/Y6M61NUtbzP4m8j4mnAa/uJz6slrI6OQsDpEjqlyV7HcTqxgI9zdnbEjTRgatCSYHfrMImtDDEhNsAeYWRR8Df2Feo8Cua0s=&amp;campaign_id=132&amp;instance_id=45669&amp;segment_id=62821&amp;user_id=104ac336ab7205252982f8c4bea84287&amp;regi_id=66509827" style="color:#004276; text-decoration: none;">http://www.nytimes.com/2014/08/25/world/middleeast/peter-theo-curtis-held-by-qaeda-affiliate-in-syria-is-freed-after-2-years.html?emc=edit_na_20140824</a></td>
           </tr>
          </tbody>
         </table>
         </td>
        </tr>
       </tbody>
      </table>
      </td>
     </tr>
    </tbody>
   </table>
   </td>
  </tr>
 </tbody>
</table>
<!--aColumn Ends--><!--aColumn Extra Begins-->

<div><!--begin centered ad -->
<div style="text-align:center;">
<div style="text-align:center">
<table cellpadding="0" cellspacing="0" id="adContent" style="padding-bottom: 24px;" width="500">
 <tbody>
  <tr>
   <td style="border-top: 1px solid #cbcbcb; color:#333; font-family: georgia,serif; width:498px;" valign="top">
   <table align="center" border="0" cellpadding="0" cellspacing="0" style="padding-top: 21px;" width="358">
    <tbody>
     <tr>
      <td align="center" style="color:#909090; font-family: arial,helvetica,sans-serif; font-size: 10px; padding-bottom: 5px; text-transform: uppercase;">ADVERTISEMENT</td>
     </tr>
     <tr>
      <td align="center"><!-- ADXINFO classification="Big_Ad_-_Email" campaign="nl-2013remnant_html_BreakingNewsAlert_LiveIntent" priority="1500" isInlineSafe="N" width="300" height="250" --><table border="0" cellpadding="0" cellspacing="0" align="center"><tr><td colspan="2"><a style="display: block; width: 300px; height: 250px;" href="http://Z.GLAKA.COM/click?s=62659&t=newsletter&sz=300x250&li=BreakingNews&m=104AC336AB7205252982F8C4BEA84287&p=2014.08.24.18.14.53" rel="nofollow"><img src="http://Z.GLAKA.COM/imp?s=62659&t=newsletter&sz=300x250&li=BreakingNews&m=104AC336AB7205252982F8C4BEA84287&p=2014.08.24.18.14.53" border="0" width="300" height="250"/></a></td></tr><tr style="display:block; height:1px; line-height:1px;"><td><img src="http://Z.GLAKA.COM/imp?s=62660&t=newsletter&sz=1x1&li=BreakingNews&m=104AC336AB7205252982F8C4BEA84287&p=2014.08.24.18.14.53" height="1" width="10" /></td><td><img src="http://Z.GLAKA.COM/imp?s=62661&t=newsletter&sz=1x1&li=BreakingNews&m=104AC336AB7205252982F8C4BEA84287&p=2014.08.24.18.14.53" height="1" width="10" /></td></tr><tr><td align="left"><a href="http://Z.GLAKA.COM/click?s=2297&t=newsletter&sz=116x15&li=BreakingNews&m=104AC336AB7205252982F8C4BEA84287&p=2014.08.24.18.14.53" rel="nofollow"><img src="http://Z.GLAKA.COM/imp?s=2297&t=newsletter&sz=116x15&li=BreakingNews&m=104AC336AB7205252982F8C4BEA84287&p=2014.08.24.18.14.53" border="0"/></a></td><td align="right"><a href="http://Z.GLAKA.COM/click?s=2298&t=newsletter&sz=69x15&li=BreakingNews&m=104AC336AB7205252982F8C4BEA84287&p=2014.08.24.18.14.53" rel="nofollow"><img src="http://Z.GLAKA.COM/imp?s=2298&t=newsletter&sz=69x15&li=BreakingNews&m=104AC336AB7205252982F8C4BEA84287&p=2014.08.24.18.14.53" border="0"/></a></td></tr></table>
<table cellpadding="0" cellspacing="0" border="0" width="24" height="6"><tbody><tr><td><img src="http://Z.GLAKA.COM/imp?s=62662&t=newsletter&sz=2x1&li=BreakingNews&m=104AC336AB7205252982F8C4BEA84287&p=2014.08.24.18.14.53" width="2" height="6" border="0" /></td><td><img src="http://Z.GLAKA.COM/imp?s=62663&t=newsletter&sz=2x1&li=BreakingNews&m=104AC336AB7205252982F8C4BEA84287&p=2014.08.24.18.14.53" width="2" height="6" border="0" /></td><td><img src="http://Z.GLAKA.COM/imp?s=62664&t=newsletter&sz=2x1&li=BreakingNews&m=104AC336AB7205252982F8C4BEA84287&p=2014.08.24.18.14.53" width="2" height="6" border="0" /></td><td><img src="http://Z.GLAKA.COM/imp?s=62665&t=newsletter&sz=2x1&li=BreakingNews&m=104AC336AB7205252982F8C4BEA84287&p=2014.08.24.18.14.53" width="2" height="6" border="0" /></td><td><img src="http://Z.GLAKA.COM/imp?s=62666&t=newsletter&sz=2x1&li=BreakingNews&m=104AC336AB7205252982F8C4BEA84287&p=2014.08.24.18.14.53" width="2" height="6" border="0" /></td><td><img src="http://Z.GLAKA.COM/imp?s=62667&t=newsletter&sz=2x1&li=BreakingNews&m=104AC336AB7205252982F8C4BEA84287&p=2014.08.24.18.14.53" width="2" height="6" border="0" /></td><td><img src="http://Z.GLAKA.COM/imp?s=62668&t=newsletter&sz=2x1&li=BreakingNews&m=104AC336AB7205252982F8C4BEA84287&p=2014.08.24.18.14.53" width="2" height="6" border="0" /></td><td><img src="http://Z.GLAKA.COM/imp?s=62669&t=newsletter&sz=2x1&li=BreakingNews&m=104AC336AB7205252982F8C4BEA84287&p=2014.08.24.18.14.53" width="2" height="6" border="0" /></td><td><img src="http://Z.GLAKA.COM/imp?s=62670&t=newsletter&sz=2x1&li=BreakingNews&m=104AC336AB7205252982F8C4BEA84287&p=2014.08.24.18.14.53" width="2" height="6" border="0" /></td><td><img src="http://Z.GLAKA.COM/imp?s=62671&t=newsletter&sz=2x1&li=BreakingNews&m=104AC336AB7205252982F8C4BEA84287&p=2014.08.24.18.14.53" width="2" height="6" border="0" /></td><td><img src="http://Z.GLAKA.COM/imp?s=62672&t=newsletter&sz=2x1&li=BreakingNews&m=104AC336AB7205252982F8C4BEA84287&p=2014.08.24.18.14.53" width="2" height="6" border="0" /></td><td><img src="http://Z.GLAKA.COM/imp?s=62673&t=newsletter&sz=2x1&li=BreakingNews&m=104AC336AB7205252982F8C4BEA84287&p=2014.08.24.18.14.53" width="2" height="6" border="0" /></td></tr></tbody></table></td>
     </tr>
    </tbody>
   </table>
   </td>
  </tr>
 </tbody>
</table>
</div>
</div>
<!--end centered ad --></div>
<!--aColumn Extra Ends--><img height="1" src="http://pixel.monitor1.returnpath.net/pixel.gif?r=9dbaad520ccf62ad66a806f22f1f3b4689e74e0c&amp;s=@gmail.com&amp;id=66509827&amp;gender=0" width="1" /> <!--Footer Extra Begins-->
<table cellpadding="0" cellspacing="0" id="footerExtra" style="color: #333; font-family:arial,helvetica,sans-serif; font-size: 13px;" width="500">
 <tbody>
  <tr>
   <td style="padding-bottom:10px; text-align: center;">
   <table border="0" cellpadding="0" cellspacing="0" style="border-top: 1px solid #CCC; font-size: 10px; padding:14px 0 3px 0; -webkit-text-size-adjust: none;" width="100%">
    <tbody>
     <tr>
      <td align="right" style="color: #333; font-family:arial,helvetica,sans-serif; padding-right:10px;">FOLLOW US:</td>
      <td align="center" width="10"><img alt="Twitter" class="twitterIcon" src="http://graphics8.nytimes.com/images/article/functions/twitter.gif" style="vertical-align: bottom;" /></td>
      <td align="center" style="font-size: 12px; margin-left: -1px; vertical-align: bottom;" width="62"><a href="http://p.nytimes.com/email/re?location=4z5Q7LhI+KWkqACYTXC750bgRZRyRovPJIEXYIwmphMhsceygzyJu29NpehnSEEo6LFfBsWzU88ZuR9TBid9Iw==&amp;campaign_id=132&amp;instance_id=45669&amp;segment_id=62821&amp;user_id=104ac336ab7205252982f8c4bea84287&amp;regi_id=66509827" style="color:#004276; font-family:arial,helvetica,sans-serif; text-decoration: none;">@NYTimes</a></td>
      <td align="center" width="9"><span style="color: #004276; padding: 0 3px 0 2px;">|</span></td>
      <td align="center" width="14"><img alt="Facebook" height="14" src="http://graphics8.nytimes.com/images/icons/sharetools/classic/facebook.gif" style="vertical-align: bottom;" width="14" /></td>
      <td align="left" style="font-size: 12px; padding-left: 2px; vertical-align: bottom;"><a href="http://p.nytimes.com/email/re?location=4z5Q7LhI+KWmuacvKaClnScR5XC3Nq7OzJ3aY8D5xuztsiw8wiydwih+6YpVRsJc+1yYzu7dP2XDVJCMCnMSSQ==&amp;campaign_id=132&amp;instance_id=45669&amp;segment_id=62821&amp;user_id=104ac336ab7205252982f8c4bea84287&amp;regi_id=66509827" style="color:#004276; font-family:arial,helvetica,sans-serif; text-decoration: none;">Facebook</a></td>
     </tr>
    </tbody>
   </table>
   </td>
  </tr>
  <tr>
   <td align="left" style="font-family: Arial,Helvetica,sans-serif; font-size: 12px; padding-bottom: 27px;">
   <table style="background-color: rgb(244, 244, 244); border-bottom: 1px solid #CCC; border-top: 1px solid #CCC; padding: 20px 0 0 20px;  width: 500px;">
    <tbody>
     <tr>
      <td>
      <table style="background-color: rgb(244, 244, 244); color: #333; font-family: arial,helvetica,sans-serif; font-size: 13px; height: 100px; line-height: 15px; padding-bottom: 1px; width: 500px;">
       <tbody>
        <tr>
         <td>
         <table border="0" cellpadding="0" cellspacing="0" style="background-color: rgb(244, 244, 244); padding-bottom: 15px; padding-left: 10px; width: 500px; -webkit-text-size-adjust: none;">
          <tbody>
           <tr>
            <td colspan="2" style="color: #333; font-family: arial,helvetica,sans-serif; font-size: 13px; line-height: 15px; padding-bottom: 18px; ">For breaking news on your mobile device, go to <a href="http://p.nytimes.com/email/re?location=XzzrGg+iKs9YRj28EbS+JZqSgyyIS3CdRlcqaJDmNgZR6U1YR0ohawHboqRm57jG&amp;campaign_id=132&amp;instance_id=45669&amp;segment_id=62821&amp;user_id=104ac336ab7205252982f8c4bea84287&amp;regi_id=66509827" style="color:#004276; font-weight: bold; text-decoration:none;">m.nyt.com &raquo;</a></td>
           </tr>
           <tr>
            <td style="color: #004276; text-align: left; text-decoration: none;"><a href="http://p.nytimes.com/email/re?location=4z5Q7LhI+KVBjmEgFdYACKqxdqYf5TwZIbHHsoM8ibtvTaXoZ0hBKOixXwbFs1PPGbkfUwYnfSM=&amp;campaign_id=132&amp;instance_id=45669&amp;segment_id=62821&amp;user_id=104ac336ab7205252982f8c4bea84287&amp;regi_id=66509827" style="color:#004276; float:left; text-decoration: none;"><img alt="NYT" border="0" height="34" src="http://graphics8.nytimes.com/images/logos/nyt/t-logo-50x50-666666-151515.gif" style="padding-right:12px;" width="35" /> </a></td>
            <td><span style="color: #333; font-family: arial,helvetica,sans-serif; font-size: 13px;  line-height: 15px; margin-bottom: 4px;">Access The New York Times from anywhere with our suite of apps:</span><br />
            <span style="margin-left: 0; font-family: arial,helvetica,sans-serif; font-size: 13px;  line-height: 15px;"><a href="http://p.nytimes.com/email/re?location=pMJKdIFVI6rQ1A83aF/Jg7/jh3WgNi7QC3ByiIlXfeyrM7pHW6Nf0YR2ZYH3dsmtEsAl7dga78yL6JFK19p6o+hhB8ykGYZz&amp;campaign_id=132&amp;instance_id=45669&amp;segment_id=62821&amp;user_id=104ac336ab7205252982f8c4bea84287&amp;regi_id=66509827" style="color: #004276; text-decoration: none;">iPhone&reg;</a> </span> <span style="color: #004276; padding: 0 1px 0 2px; ">|</span> <span style="font-family: arial,helvetica,sans-serif; font-size: 13px; margin-left: 0; line-height: 15px;"> <a href="http://p.nytimes.com/email/re?location=pMJKdIFVI6rQ1A83aF/Jg7/jh3WgNi7QC3ByiIlXfeyrM7pHW6Nf0aeI6OX4PrLsH+x+BGVRjPrTK+6nwkWAZmxECYpgphj+&amp;campaign_id=132&amp;instance_id=45669&amp;segment_id=62821&amp;user_id=104ac336ab7205252982f8c4bea84287&amp;regi_id=66509827" style="color: #004276; text-decoration: none;">iPad&reg;</a> </span> <span style="color: #004276; padding: 0 2px;">|</span> <span style="font-family: arial,helvetica,sans-serif; font-size: 13px; margin-left: 0; line-height: 15px;"> <a href="http://p.nytimes.com/email/re?location=pMJKdIFVI6rUh2PW8hSxDWaBY9XU8GQKtWNzHdH3en9ZHcSRhJ2hEmDwszquUW3CIOwqcCRo5H9ukhmkhh6HKYf91At3OSX4XuXVqLUBJbs61squNiQonUgEfZzB7Yje0tm2Ooiy/S5dJFqVEKZ7GBc7Bpqh/PZz&amp;campaign_id=132&amp;instance_id=45669&amp;segment_id=62821&amp;user_id=104ac336ab7205252982f8c4bea84287&amp;regi_id=66509827" style="color:#004276; text-decoration: none;">Android</a> </span> <span style="color: #004276; padding: 0 2px;">|</span> <span style="font-family: arial,helvetica,sans-serif; font-size: 13px; margin-left: 0; line-height: 15px;"> <a href="http://p.nytimes.com/email/re?location=4z5Q7LhI+KVBjmEgFdYACDuqzkg7rwCI6o0MKWD0L+R/dy2cD6PGiMA+x7BjGIifmXtrOBE/zshdbDPu26DFLCIIgdKzVWca&amp;campaign_id=132&amp;instance_id=45669&amp;segment_id=62821&amp;user_id=104ac336ab7205252982f8c4bea84287&amp;regi_id=66509827" style="color:#004276; text-decoration: none;">All</a> </span></td>
           </tr>
          </tbody>
         </table>
         </td>
        </tr>
       </tbody>
      </table>
      </td>
     </tr>
    </tbody>
   </table>
   </td>
  </tr>
 </tbody>
</table>
<!--Ends Footer Extra--> <!--Footer Begins-->

<table cellpadding="0" cellspacing="0" id="customFooter" width="500">
 <tbody>
  <tr>
   <td colspan="2" style="font-family: arial,helvetica,sans-serif; font-size: 11px; ">
   <div>
   <h4 style="color:#333; font-size: 11px; font-weight: bold; margin: 0 0 11px; text-transform: none;">About This Email</h4>

   <p style="color:#333; margin-bottom: 4px;">This is an automated email. Please do not reply directly to this email.</p>

   <p style="color:#333; margin: 0;">You received this message because you signed up for NYTimes.com's breaking news alerts. As a member of the TRUSTe privacy program, we are committed to protecting your privacy.</p>
   </div>

   <div style="margin-top: 11px"><span><span style="color: #666; font-size:11px; padding: 0 3px 0 3px;"><span><a href="http://www.nytimes.com/gst/unsub.html?email=@gmail.com&id=66509827&segment=62821&group=nl&product=NA" style="color:#004276; text-decoration: none;"></a></span> |</span> </span> <span style="margin-left: 0"> <a href="http://p.nytimes.com/email/re?location=4z5Q7LhI+KVBjmEgFdYACHjdnNiNT4YKoxod5S8PmJHLwSaYHKQ3OoW+Ockj4gmLUzeq86lskkg=&amp;campaign_id=132&amp;instance_id=45669&amp;segment_id=62821&amp;user_id=104ac336ab7205252982f8c4bea84287&amp;regi_id=66509827" style="color:#004276; text-decoration: none;">Manage Subscriptions</a> <span style="color: #666; font-size:11px; padding: 0 3px 0 3px;">|</span> </span> <span style="margin-left: 0"> <a href="http://p.nytimes.com/email/re?location=pMJKdIFVI6rcMVYYXia0CswGfPcU3iP2BbboDbCax1YCgTSJ+ur9HWzXi1VU9xZiy8EmmBykNzqFvjnJI+IJi1M3qvOpbJJI&amp;campaign_id=132&amp;instance_id=45669&amp;segment_id=62821&amp;user_id=104ac336ab7205252982f8c4bea84287&amp;regi_id=66509827" style="color:#004276; text-decoration: none;">Change Your Email</a> <span style="color: #666; font-size:11px; padding: 0 3px 0 3px; ">|</span> </span> <span style="margin-left: 0"> <a href="http://p.nytimes.com/email/re?location=4z5Q7LhI+KVBjmEgFdYACMlEhIhWVuPIxganfKahJGpDcKtdpfztygcexoigFj8dX+EMGUu62OTaaiHAHzlw2EZXKmiQ5jYGUelNWEdKIWsB26KkZue4xg==&amp;campaign_id=132&amp;instance_id=45669&amp;segment_id=62821&amp;user_id=104ac336ab7205252982f8c4bea84287&amp;regi_id=66509827" style="color:#004276; text-decoration: none;">Privacy Policy</a> <span style="color: #666; font-size:11px; padding: 0 3px 0 3px; ">|</span> </span> <span style="margin-left: 0"> <a href="http://p.nytimes.com/email/re?location=4z5Q7LhI+KVBjmEgFdYACDE6h1nWlLhfgdX2Y/OzTy3J7c0AmM08e7ziiNbylE13sP5ku+sbDLgjEe9PMPAFMUJzhTEwqtoPG/ft8gXI4cr6RBpQS6Iucw==&amp;campaign_id=132&amp;instance_id=45669&amp;segment_id=62821&amp;user_id=104ac336ab7205252982f8c4bea84287&amp;regi_id=66509827" style="color:#004276; text-decoration: none;">Contact</a> </span><span style="margin-left: 0;">&nbsp;</span></div>
   </td>
  </tr>
  <tr>
   <td align="left" style="font-family: arial,helvetica,sans-serif; font-size: 11px; padding: 21px 0 17px 0; display:block; ">
   <div style="border-top: 1px solid rgb(226, 226, 226);">
   <div style="color:#000001; font-family: arial,helvetica,sans-serif; font-size: 12px; height: 20px; padding-top:7px;"><span style="color: #909090; font-family: arial,helvetica,sans serif; font-size:10px;">Copyright 2014</span> <span style="color: #909090; font-size:10px; padding-right:2px;">|</span> <span style="color: #909090; font-family: arial,helvetica,sans serif; font-size:10px;">The New York Times Company</span> <span style="color: #909090; font-size:10px; padding: 0 5px 0 2px; ">|</span><span style="color: #909090; font-family: arial,helvetica,sans serif; font-size:10px;">NYTimes.com 620 Eighth Avenue New York, NY 10018 </span></div>
   </div>
   </td>
  </tr>
 </tbody>
</table>
<!--Footer Ends--></div>
</div>
</body>
</html>
<img src="http://p.nytimes.com/email/re?location=hdaNaYedr2/IomeWRKt0nffrak8aSGLbvtkkq/r7ihwOf5XePlpJ1w==&amp;campaign_id=132&amp;instance_id=45669&amp;segment_id=62821&amp;user_id=104ac336ab7205252982f8c4bea84287&amp;regi_id=66509827"/>`),
	}

	want := []string{
		"To ensure delivery to your inbox, please add",
		"nytdirect@nytimes.com",
		"to your address book.",
		"|",
		"BREAKING NEWS ALERT",
		"NYTimes.com",
		"|",
		"BREAKING NEWS",
		"Sunday, August 24, 2014 2:11 PM EDT",
		"U.S. Journalist Held by Qaeda Affiliate in Syria Is Freed After Nearly 2 Years",
		`An American journalist held captive for nearly two years by Al Qaeda’s
official branch in Syria has been freed, according to a representative of the
journalist’s family and a report on Sunday by the Al Jazeera
network.`,
		`The journalist, Peter Theo Curtis, was abducted near the
Syria-Turkey border in October 2012. He was held by the Nusra Front, the Qaeda
affiliate in Syria, which has broken with the more radical Islamic State in Iraq
and Syria, or ISIS. Another American journalist, James W. Foley, who was
kidnapped in Syria the following month, was beheaded last week by ISIS, which
posted images of his execution on YouTube.`,
		`A family friend confirmed on
Sunday that Mr. Curtis, originally from Boston, had been handed over to a United
Nations representative.`,
		"READ MORE »",
		"http://www.nytimes.com/2014/08/25/world/middleeast/peter-theo-curtis-held-by-qaeda-affiliate-in-syria-is-freed-after-2-years.html?emc=edit_na_20140824",
		"ADVERTISEMENT",
		"FOLLOW US:",
		"@NYTimes",
		"|",
		"Facebook",
		"For breaking news on your mobile device, go to",
		"m.nyt.com »",
		"Access The New York Times from anywhere with our suite of apps:",
		"iPhone®",
		"|",
		"iPad®",
		"|",
		"Android",
		"|",
		"All",
		"About This Email",
		"This is an automated email. Please do not reply directly to this email.",
		"You received this message because you signed up for NYTimes.com's breaking news alerts. As a member of the TRUSTe privacy program, we are committed to protecting your privacy.",
		"|",
		"Manage Subscriptions",
		"|",
		"Change Your Email",
		"|",
		"Privacy Policy",
		"|",
		"Contact",
		"Copyright 2014",
		"|",
		"The New York Times Company",
		"|",
		"NYTimes.com 620 Eighth Avenue New York, NY 10018",
	}

	got, err := email.VisibleText()
	if err != nil {
		t.Errorf("VisibleText() encountered unexpected error: %s", err)
		return
	}
	for i, _ := range got {
		if want[i] != string(got[i]) {
			t.Errorf("VisibleText() want[%d](%s) != got[%d](%s)", i, want[i], i, string(got[i]))
		}
	}
}

const quotedEmail ="Delivered-To: an.email.address@gmail.com\r\nReceived: by 10.220.224.7 with SMTP id im7csp165179vcb;\r\n        Mon, 11 Aug 2014 15:14:17 -0700 (PDT)\r\nX-Received: by 10.66.240.140 with SMTP id wa12mr524751pac.99.1407795256741;\r\n        Mon, 11 Aug 2014 15:14:16 -0700 (PDT)\r\nReturn-Path: <bo-b65ymr9bfbugyhauy2x7bbykuhtky7@b.e.latimes.com>\r\nReceived: from mta852.e.latimes.com (mta852.e.latimes.com. [63.232.236.160])\r\n        by mx.google.com with ESMTP id kd14si14437848pbb.64.2014.08.11.15.14.16\r\n        for <an.email.address@gmail.com>;\r\n        Mon, 11 Aug 2014 15:14:16 -0700 (PDT)\r\nReceived-SPF: pass (google.com: domain of bo-b65ymr9bfbugyhauy2x7bbykuhtky7@b.e.latimes.com designates 63.232.236.160 as permitted sender) client-ip=63.232.236.160;\r\nAuthentication-Results: mx.google.com;\r\n       spf=pass (google.com: domain of bo-b65ymr9bfbugyhauy2x7bbykuhtky7@b.e.latimes.com designates 63.232.236.160 as permitted sender) smtp.mail=bo-b65ymr9bfbugyhauy2x7bbykuhtky7@b.e.latimes.com;\r\n       dkim=pass header.i=@e.latimes.com\r\nDKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed; d=e.latimes.com;\r\n\ts=20120316; t=1407795256; x=1423692856;\r\n\tbh=CUzkYJbRqeJ0BB67gF474DY+T+fY0fLYMvAR3aPdlow=; h=From:Reply-To;\r\n\tb=gIPnif1mtRvQ/8DG4nqqCvaq6sNBJPAA8syDte/LQsMgPaXZBF4vEDc0t1ThWQtJF\r\n\t Yz8cDtWLfHv3fKx52DXhfTczzRGnOmpZYM1Z9DaYnIBzLcCxKwls/KYAjHmkSEZgeQ\r\n\t 13jIuKIu3eqVBRy5KypzJwPz9Ao5i0YSwOZYN/RM=\r\nDomainKey-Signature: a=rsa-sha1; q=dns; c=nofws;\r\n  s=200505; d=e.latimes.com;\r\n  b=WZ2Hpj3Ke741wPIrt7DXtfArp9aUrk63jUhl9Px7st2cUVj/dDxQM6F+jqdJmuyg6LgTBQ0gtSWU1VhXcYjgF1+t3y2CNap8lNi7+FYaWo9T2TZi2CnOLTfq5vc1i8uuTTqTginraNmYu1w+oj07GaKD5P2pTVwfJVdsAB6jNRk=;\r\n h=Date:Message-ID:List-Unsubscribe:From:To:Subject:MIME-Version:Reply-To:Content-type:Content-Transfer-Encoding;\r\nDate: Mon, 11 Aug 2014 22:14:16 -0000\r\nMessage-ID: <b65ymr9bfbugyhauy2x7bbykuhtky7.8225590.5365@mta852.e.latimes.com>\r\nList-Unsubscribe: <mailto:rm-0b65ymrykuhtky7@e.latimes.com>\r\nFrom: \"Los Angeles Times\" <news@e.latimes.com>\r\nTo: an.email.address@gmail.com\r\nSubject: Breaking News: Hostage in Stockton bank robbery was killed by officers, not suspects\r\nMIME-Version: 1.0\r\nReply-To: \"Los Angeles Times\" <support-b65ymr9bfbugyhauy2x7bbykuhtky7@e.latimes.com>\r\nContent-type: text/html; charset=\"iso-8859-1\"\r\nContent-Transfer-Encoding: quoted-printable\r\n\r\n\xC4pple<html><head></head><body onload=3D''><div class=3D\"module blurb clearfix\">\r\n\t<style><![CDATA[\r\n\t#email-wrapper {font-family: Georgia,Times,serif; font-size: 14px; width: =\r\n630px; padding: 10px;}\r\n\t#breaking-news-banner {width: 100%;}\r\n\t#emailadbox {width: 300px; height: 250px; margin-top: 3px;}\r\n#banner-graphic {margin-bottom:5px;}\r\n        #storyslug {width:280px; font-family: Georgia,Times,serif; font-siz=\r\ne: 14px;}\r\n\tp#emailad  {font-family: Arial,Helvetica,sans-serif; font-size: 10px; colo=\r\nr: #999; letter-spacing: 1px; text-align: center; margin-bottom:0px; margin=\r\n-top:5px;}\r\n\t.bottom-text {font-size: .85em; border-top: 1px solid #ccc; margin-top: 24=\r\npx; padding-top: 3px;}\r\n\t.bottom-text p {margin-top:4px; margin-bottom:1px;}\r\n\tp.email-head {margin-bottom:10px;}=09\r\n.email-date\t{color: #930000 ; font-style: italic; font-size: 11px; }\r\n.email-graph { margin-bottom: 12px;}\r\n]]></style><div id=3D\"email-wrapper\">\r\n\r\n<div id=3D\"banner-graphic\"><img src=3D\"http://www.latimes.com/media/graphic=\r\n/2010-02/52101671.png\" alt=3D\"Los Angeles Times\" /></div>\r\n\r\n<div id=3D\"breaking-news-banner\"><img src=3D\"http://www.latimes.com/media/a=\r\nlternatethumbnails/blurb/2012-07/47391835-16074348.gif\" alt=3D\"Breaking new=\r\ns\" border=3D\"0\" /></div>\r\n\r\n<table width=3D\"630\"><tr><td>\r\n<div id=3D\"storyslug\">\r\n<h1><a style=3D\"font-size: 20px; color: black\">Hostage in Stockton bank rob=\r\nbery was killed by officers, not suspects</a></h1>\r\n\r\n<!--<div style=3D\"margin-bottom: 12px;\" class=3D\"email-date\">Los Angeles Ti=\r\nmes | May 22, 2012 | 11:47 a.m.</div>-->\r\n<div style=3D\"margin-bottom: 12px;\" class=3D\"email-date\">Los Angeles Times =\r\n| August 11, 2014 |  3:09 PM</div>\r\n=20\r\n\r\n<p><p>Officials Monday said a Stockton woman taken hostage and used as a hu=\r\nman shield during a bank robbery turned police chase last month was killed =\r\nby gunfire from officers, not the suspects.</p>&#13;\r\n<p>A preliminary ballistics report indicates it was bullets from the police=\r\n that killed Misty Jean Holt-Singh during the chaotic July 16 gun battle, S=\r\ntockton Police Chief Eric Jones said. Initial reports suggest she was shot =\r\nabout 10 times, he added.</p>&#13;\r\n<p>The three suspects in the case -- two of whom were also killed -- fired =\r\nmore than 100 bullets during the one-hour incident, Jones said. Preliminary=\r\n reports show 33 police officers fired an estimated 600 bullets, he added.<=\r\n/p>&#13;\r\n<p>For the latest information go to <a href=3D\"http://e.latimes.com/a/hBT6T=\r\n8rB8hLWGB87vhDAAfYM2RC/exmp1\">www.latimes.com</a>.</p></p>\r\n</div>\r\n\r\n</td>\r\n<td width=3D\"330\" align=3D\"center\">\r\n\t\t<div style=3D\"margin: 30px 0;\"><center><span style=3D\"color: #c2c2c2;font=\r\n-family: arial; font-size:8px; line-height: 22px;\">ADVERTISEMENT</span></ce=\r\nnter>\r\n\t\t\t\t<table border=3D\"0\" cellpadding=3D\"0\" cellspacing=3D\"0\"><tr><td colspan=\r\n=3D\"2\"><a style=3D\"display: block; width: 300px; height: 250px;\" href=3D\"ht=\r\ntp://li.latimes.com/click?s=3D73326&t=3Dnewsletter&sz=3D300x250&li=3DLATime=\r\ns&e=3Dan.email.address@gmail.com&cm=3D001_hBT6T8rB8hLWGB87vhDAAfYM2RC\"=\r\n rel=3D\"nofollow\"><img src=3D\"http://li.latimes.com/imp?s=3D73326&t=3Dnewsl=\r\netter&sz=3D300x250&li=3DLATimes&e=3Dan.email.address@gmail.com&cm=3D00=\r\n1_hBT6T8rB8hLWGB87vhDAAfYM2RC\" border=3D\"0\" width=3D\"300\" height=3D\"250\" />=\r\n</a></td></tr><tr style=3D\"display:block; height:1px; line-height:1px;\"><td=\r\n><img src=3D\"http://li.latimes.com/imp?s=3D73327&t=3Dnewsletter&sz=3D1x1&li=\r\n=3DLATimes&e=3Dan.email.address@gmail.com&cm=3D001_hBT6T8rB8hLWGB87vhD=\r\nAAfYM2RC\" height=3D\"1\" width=3D\"10\" /></td><td><img src=3D\"http://li.latime=\r\ns.com/imp?s=3D73328&t=3Dnewsletter&sz=3D1x1&li=3DLATimes&e=3D@gmail.com&cm=3D001_hBT6T8rB8hLWGB87vhDAAfYM2RC\" height=3D\"1\" width=\r\n=3D\"10\" /></td></tr><tr><td align=3D\"left\"><a href=3D\"http://li.latimes.com=\r\n/click?s=3D49864&t=3Dnewsletter&sz=3D116x15&li=3DLATimes&e=3D@gmail.com&cm=3D001_hBT6T8rB8hLWGB87vhDAAfYM2RC\" rel=3D\"nofollow\"><i=\r\nmg src=3D\"http://li.latimes.com/imp?s=3D49864&t=3Dnewsletter&sz=3D116x15&li=\r\n=3DLATimes&e=3Dan.email.address@gmail.com&cm=3D001_hBT6T8rB8hLWGB87vhD=\r\nAAfYM2RC\" border=3D\"0\" /></a></td><td align=3D\"right\"><a href=3D\"http://li.=\r\nlatimes.com/click?s=3D49865&t=3Dnewsletter&sz=3D69x15&li=3DLATimes&e=3Db=\r\nkingr@gmail.com&cm=3D001_hBT6T8rB8hLWGB87vhDAAfYM2RC\" rel=3D\"no=\r\nfollow\"><img src=3D\"http://li.latimes.com/imp?s=3D49865&t=3Dnewsletter&sz=\r\n=3D69x15&li=3DLATimes&e=3Dan.email.address@gmail.com&cm=3D001_hBT6T8rB=\r\n8hLWGB87vhDAAfYM2RC\" border=3D\"0\" /></a></td></tr></table><br /><table cell=\r\npadding=3D\"0\" cellspacing=3D\"0\" border=3D\"0\" width=3D\"24\" height=3D\"6\"><tbo=\r\ndy><tr><td><img src=3D\"http://li.latimes.com/imp?s=3D77983&t=3Dnewsletter&s=\r\nz=3D2x1&li=3DLATimes&e=3Dan.email.address@gmail.com&cm=3D001_hBT6T8rB8=\r\nhLWGB87vhDAAfYM2RC\" width=3D\"2\" height=3D\"6\" border=3D\"0\" /></td><td><img s=\r\nrc=3D\"http://li.latimes.com/imp?s=3D77984&t=3Dnewsletter&sz=3D2x1&li=3DLATi=\r\nmes&e=3Dan.email.address@gmail.com&cm=3D001_hBT6T8rB8hLWGB87vhDAAfYM2R=\r\nC\" width=3D\"2\" height=3D\"6\" border=3D\"0\" /></td><td><img src=3D\"http://li.l=\r\natimes.com/imp?s=3D77985&t=3Dnewsletter&sz=3D2x1&li=3DLATimes&e=3Dr@gmail.com&cm=3D001_hBT6T8rB8hLWGB87vhDAAfYM2RC\" width=3D\"2\" he=\r\night=3D\"6\" border=3D\"0\" /></td><td><img src=3D\"http://li.latimes.com/imp?s=\r\n=3D77986&t=3Dnewsletter&sz=3D2x1&li=3DLATimes&e=3Dan.email.address@gma=\r\nil.com&cm=3D001_hBT6T8rB8hLWGB87vhDAAfYM2RC\" width=3D\"2\" height=3D\"6\" borde=\r\nr=3D\"0\" /></td><td><img src=3D\"http://li.latimes.com/imp?s=3D77987&t=3Dnews=\r\nletter&sz=3D2x1&li=3DLATimes&e=3Dan.email.address@gmail.com&cm=3D001_h=\r\nBT6T8rB8hLWGB87vhDAAfYM2RC\" width=3D\"2\" height=3D\"6\" border=3D\"0\" /></td><t=\r\nd><img src=3D\"http://li.latimes.com/imp?s=3D77988&t=3Dnewsletter&sz=3D2x1&l=\r\ni=3DLATimes&e=3Dan.email.address@gmail.com&cm=3D001_hBT6T8rB8hLWGB87vh=\r\nDAAfYM2RC\" width=3D\"2\" height=3D\"6\" border=3D\"0\" /></td><td><img src=3D\"htt=\r\np://li.latimes.com/imp?s=3D77989&t=3Dnewsletter&sz=3D2x1&li=3DLATimes&e=3Db=\r@gmail.com&cm=3D001_hBT6T8rB8hLWGB87vhDAAfYM2RC\" width=\r\n=3D\"2\" height=3D\"6\" border=3D\"0\" /></td><td><img src=3D\"http://li.latimes.c=\r\nom/imp?s=3D77990&t=3Dnewsletter&sz=3D2x1&li=3DLATimes&e=3D&cm=3D001_hBT6T8rB8hLWGB87vhDAAfYM2RC\" width=3D\"2\" height=3D\"=\r\n6\" border=3D\"0\" /></td><td><img src=3D\"http://li.latimes.com/imp?s=3D77991&=\r\nt=3Dnewsletter&sz=3D2x1&li=3DLATimes&e=3Dan.email.address@gmail.com&cm=\r\n=3D001_hBT6T8rB8hLWGB87vhDAAfYM2RC\" width=3D\"2\" height=3D\"6\" border=3D\"0\" /=\r\n></td><td><img src=3D\"http://li.latimes.com/imp?s=3D77992&t=3Dnewsletter&sz=\r\n=3D2x1&li=3DLATimes&e=3Dan.email.address@gmail.com&cm=3D001_hBT6T8rB8h=\r\nLWGB87vhDAAfYM2RC\" width=3D\"2\" height=3D\"6\" border=3D\"0\" /></td><td><img sr=\r\nc=3D\"http://li.latimes.com/imp?s=3D77993&t=3Dnewsletter&sz=3D2x1&li=3DLATim=\r\nes&e=3Dan.email.address@gmail.com&cm=3D001_hBT6T8rB8hLWGB87vhDAAfYM2RC=\r\n\" width=3D\"2\" height=3D\"6\" border=3D\"0\" /></td><td><img src=3D\"http://li.la=\r\ntimes.com/imp?s=3D77994&t=3Dnewsletter&sz=3D2x1&li=3DLATimes&e=3DD001_hBT6T8rB8hLWGB87vhDAAfYM2RC\" width=3D\"2\" hei=\r\nght=3D\"6\" border=3D\"0\" /></td></tr></tbody></table></div>\r\n</td>\r\n\r\n</tr></table><div class=3D\"bottom-text\">\r\n\t<p class=3D\"email-head\">Text \"=\r\news text alerts. You will receive 2 msgs/week. Msg&amp;data rates may apply=\r\n. Text HELP for help. Text STOP to cancel.</p>\r\n=09\r\n\t\t<p>California and the world: Visit <a href=3D\"http://e.latimes.com/a/hBT6=\r\nT8rB8hLWGB87vhDAAfYM2RC/exmp1\">http://www.latimes.com</a> for up-to-the-min=\r\nute news.</p>\r\n\t\t<p><strong>Follow</strong> @LATimes on Twitter: <a href=3D\"http://e.latim=\r\nes.com/a/hBT6T8rB8hLWGB87vhDAAfYM2RC/exmp2\">http://twitter.com/latimes</a><=\r\n/p>\r\n\t\t<p><strong>Connect</strong> with the L.A. Times on Facebook: <a href=3D\"h=\r\nttp://e.latimes.com/a/hBT6T8rB8hLWGB87vhDAAfYM2RC/exmp3\">http://facebook.co=\r\nm/latimes</a></p>\r\n\t\t<p><strong>Sign up</strong> for more email newsletters: <a href=3D\"http:/=\r\n/e.latimes.com/a/hBT6T8rB8hLWGB87vhDAAfYM2RC/exmp4\">http://latimes.com/news=\r\nletters</a></p>\r\n=09\r\n</div>\r\n\r\n\r\n\r\n<div class=3D\"bottom-text\">\r\n\t\t<p class=3D\"email-head\"><i>About this communication:</i></p>\r\n\t\t<p>You are receiving this email because you opted to receive Breaking New=\r\ns Alerts from the Los Angeles Times.</p>\r\n\t\t<p>You're currently subscribed to Los Angeles Times Breaking News with th=\r\ne address an.email.address@gmail.com. If you'd like to unsubscribe, pl=\r\nease click here: <a href=3D\"http://e.latimes.com/a/hBT6T8rB8hLWGB87vhDAAfYM=\r\n2RC/exmp5?email=3Dan.email.address@gmail.com\">http://ebm.cheetahmail.c=\r\nom/r/webunsub?t=3DBT6M2RC&amp;email=3Dcom&amp;n=3D1</a></p>\r\n\t\t<p>You can also unsubscribe by modifying your profile on latimes.com at <=\r\na href=3D\"http://e.latimes.com/a/hBT6hDAAfYM2RC/exmp6\">http://=\r\nwww.latimes.com/newsletters</a></p>\r\n\t\t<p>For information on how we protect your information, please read our pr=\r\nivacy policy at <a href=3D\"http://e.latimes.com/a/hBT6T8rB8hLWGB87vhDAAfYM2=\r\nRC/exmp7\">http://www.latimes.com/privacypolicy</a></p>\r\n=09\r\n</div>\r\n\r\n=09\r\n\r\n</div>\r\n</div>\r\n\r\n<!--x-Instance-Name: i5latisrapp08--><img src=3D\"http://e.latimes.com/a/hBT=\r\n6T8rB8hLWGB87vhDAAfYM2RC/spacer.gif\">\r\n</body></html>=\r\n\r\n" 
const htmlEmail = "Delivered-To: an.email.address@gmail.com\r\nReceived: by 10.220.224.7 with SMTP id im7csp226521vcb;\r\n        Tue, 12 Aug 2014 10:49:55 -0700 (PDT)\r\nX-Received: by 10.236.81.243 with SMTP id m79mr20366101yhe.28.1407865795556;\r\n        Tue, 12 Aug 2014 10:49:55 -0700 (PDT)\r\nReturn-Path: <foxnews_BADE9CA9E0AAA50AF208A0917BB3264E@response.foxnews.com>\r\nReceived: from vmta.response.foxnews.com ([216.87.167.12])\r\n        by mx.google.com with ESMTP id t94si33703331yhp.75.2014.08.12.10.49.55\r\n        for <an.email.address@gmail.com>;\r\n        Tue, 12 Aug 2014 10:49:55 -0700 (PDT)\r\nReceived-SPF: pass (google.com: domain of foxnews_BADE9CA9E0AAA50AF208A0917BB3264E@response.foxnews.com designates 216.87.167.12 as permitted sender) client-ip=216.87.167.12;\r\nAuthentication-Results: mx.google.com;\r\n       spf=pass (google.com: domain of foxnews_BADE9CA9E0AAA50AF208A0917BB3264E@response.foxnews.com designates 216.87.167.12 as permitted sender) smtp.mail=foxnews_BADE9CA9E0AAA50AF208A0917BB3264E@response.foxnews.com;\r\n       dkim=policy (weak key) header.i=@newsletters.foxnews.com\r\nDKIM-Signature: v=1; a=rsa-sha1; c=relaxed/relaxed; s=key1; d=newsletters.foxnews.com;\r\n h=Date:From:Reply-To:To:Message-ID:Subject:MIME-Version:Content-Type:Content-Transfer-Encoding:List-Unsubscribe; i=foxnews@newsletters.foxnews.com;\r\n bh=HkiAm/nPEUtzltXrGl+KGJnYQHE=;\r\n b=Gmo0MhqaZfy5xMJpIOvxM1kuJE+7j+viAp8Y7WkZzQgiaN1zrYfpbBabKMxjW/U5ri9r67/\r\n   rkL6YnMUYQ==\r\nDomainKey-Signature: a=rsa-sha1; c=nofws; q=dns; s=key1; d=newsletters.foxnews.com;\r\n b=qHcGch0YgOPIUsD8ggeQ8Sly0+QxGJ/xJ3JozcbeVLk5JcQOAbmmBP0Rj/9bS5q+EX1KkktNB65A\r\n   9I3Ina5nlQ==;\r\nReceived: from wc-robot.tpa.foxnews.com (192.168.193.69) by vmta.response.foxnews.com (PowerMTA(TM) v3.5r16) id ht99s60q7b88 for <an.email.address@gmail.com>; Tue, 12 Aug 2014 13:49:55 -0400 (envelope-from <foxnews_BADE9CA9E0AAA50AF208A0917BB3264E@response.foxnews.com>)\r\nDate: Tue, 12 Aug 2014 13:49:54 -0400\r\nFrom: \"FoxNews.com\" <foxnews@newsletters.foxnews.com>\r\nReply-To: foxnews_BADE9CA9E0AAA50AF208A0917BB3264E@newsletters.foxnews.com\r\nTo: an.email.address@gmail.com\r\nMessage-ID: <BADE9CA9E0AAA50AF208A0917BB3264E-d67ac706448a4e119df2ac91045dc209@response.foxnews.com>\r\nSubject: WATCH LIVE: News conference on death of Robin Williams\r\nMIME-Version: 1.0\r\nContent-Type: text/html; charset=UTF-8\r\nContent-Transfer-Encoding: 7bit\r\nX-Mailer: WhatCounts\r\nENVID: WC-1407865794717-BADE9CA9E0AAA50AF208A0917BB3264E-d67ac706448a4e119df2ac91045dc209\r\nList-Unsubscribe: <http://email.foxnews.com/u?id=BADE9A50AF208A0917BB3264E>\r\nX-Unsubscribe-Web: <http://email.foxnews.com/u?id=BADE9CA9E0AAA50AF208A0917BB3264E>\r\n\r\n<br />\r\n<br />\r\n <a href=\"http://email.foxnews.com/t?r=5&c=29452&l=35&ctl=57F1B:BADE9CA9E0AAA50AF208A0917BB3264E&\">http://video.foxnews.com/v/2553193403001/#sp=watch-live</a>\r\n<br />\r\n<br />\r\n <a href=\"\"></a>\r\n<br />\r\n<br />\r\nFor more news, please go to <a href=\"http://email.foxnews.com/t?r=5&c=29452&l=35&ctl=57F1C:BADE9CA9E0AAA50AF208A0917BB3264E&\">FoxNews.com</a> and watch Fox News Channel.\r\n<br />\r\n<br />\r\n<br />\r\n<tr>\r\n\r\n<td valign=\"top\" align=\"center\">\t\t\t                \t\r\n\r\n<p style=\"margin-top: 0; margin-bottom: 10px; color: #999999; font-family: arial; font-size:11px; font-weight:bold;\"><a style=\"color: #183A52; font-family: arial; font-size:11px; font-weight:bold; text-decoration:none\" href=\"http://email.foxnews.com/t?r=5&c=29452&l=35&ctl=57F1D:BADE9CA9E0AAA50AF208A0917BB3264E&\"><span style=\"color: #183a52;\">More Newsletters</span></a> | <a style=\"color: #183A52; font-family: arial; font-size:11px; font-weight:bold; text-decoration:none\" href=\"http://email.foxnews.com/u?id=BADE9CA9E0AAA50AF208A0917BB3264E\"><span style=\"color: #183a52;\">Unsubscribe</span></a> | <a style=\"color: #183A52; font-family: arial; font-size: 11px; text-decoration: none\" href=\"http://email.foxnews.com/t?r=5&c=29452&l=35&ctl=57F1E:BADE9CA9E0AAA50AF208A0917BB3264E&\"><span style=\"color: #183a52;\">Privacy Policy</span></a></p>\r\n\r\n<p style=\"margin-top: 0; margin-bottom: 10px; color: #666666; font-family: arial; font-size: 11px;\">&#169;2014 Fox News Network, LLC. All Rights Reserved.</p>\r\n\r\n</td>\r\n\r\n</tr>\r\n<p style=\"margin-top: 0; margin-bottom: 10px; color: #666666; font-family: arial; font-size: 11px;\">Fox News never sends unsolicited email. You received this email because you requested a subscription to Breaking Alerts from FoxNews.com."

const multipartEmail = `Delivered-To: an.email.address@gmail.com
Received: by 10.220.224.7 with SMTP id im7csp224598vcb;
        Tue, 12 Aug 2014 10:20:06 -0700 (PDT)
X-Received: by 10.50.143.65 with SMTP id sc1mr35630igb.19.1407864005741;
        Tue, 12 Aug 2014 10:20:05 -0700 (PDT)
Return-Path: <bounce-331977_HTML-1601718049-49633151-16472-0@bounce.e.usatoday.com>
Received: from mta.e.usatoday.com (mta.e.usatoday.com. [207.250.68.11])
        by mx.google.com with ESMTP id j15si39020373icg.10.2014.08.12.10.20.05
        for <an.email.address@gmail.com>;
        Tue, 12 Aug 2014 10:20:05 -0700 (PDT)
Received-SPF: pass (google.com: domain of bounce-331977_HTML-1601718049-49633151-16472-0@bounce.e.usatoday.com designates 207.250.68.11 as permitted sender) client-ip=207.250.68.11;
Authentication-Results: mx.google.com;
       spf=pass (google.com: domain of bounce-331977_HTML-1601718049-49633151-16472-0@bounce.e.usatoday.com designates 207.250.68.11 as permitted sender) smtp.mail=bounce-331977_HTML-1601718049-49633151-16472-0@bounce.e.usatoday.com;
       dkim=pass header.i=@e.usatoday.com
DKIM-Signature: v=1; a=rsa-sha1; c=relaxed/relaxed; s=200608; d=e.usatoday.com;
 h=From:To:Subject:Date:List-Unsubscribe:MIME-Version:Reply-To:Message-ID:Content-Type; i=newsletters@e.usatoday.com;
 bh=02BurYsjiH7bUVcml+R/S8Sw5Mk=;
 b=NzSpRT57HgSklAT9ngTV9WE4Bo4MXlgy63dX8yUsYTvqcWrHC2CkCea1
   M/NCOqkHZHeeZWyo8GdKrgJrIBjDLvUdbv2dbN9vjrfbYcIvIBHb/EEp/PKTQG01YIMkl982n+eA
   8mB2Oqg9hRXR1pCFu6s=
Received: by mta.e.usatoday.com id ht96ca163hsc for <an.email.address@gmail.com>; Tue, 12 Aug 2014 11:11:28 -0600 (envelope-from <bounce-331977_HTML-1601718049-49633151-16472-0@bounce.e.usatoday.com>)
From: "USATODAY.com" <newsletters@e.usatoday.com>
To: <an.email.address@gmail.com>
Subject: BREAKING: Clippers sale to Steve Ballmer finalized
Date: Tue, 12 Aug 2014 11:11:26 -0600
List-Unsubscribe: <mailto:leave-fcadc342838-fdf7d14747c-fe5f7077c7015-fefb1576716306-ffcf14@leave.e.usatoday.com>
MIME-Version: 1.0
Reply-To: "USATODAY.com" <reply-fe5f10797367077c7015-331977_HTML-1601718049-16472-0@reply.e.usatoday.com>
x-job: 16472_49633151
Message-ID: <6a356406-f972-49b6-915b-90d7cd32fbb2@xtinmta496.xt.local>
Content-Type: multipart/alternative;
	boundary="eZDakj4l4DVQ=_?:"

This is a multi-part message in MIME format.

--eZDakj4l4DVQ=_?:
Content-Type: text/plain;
	charset="iso-8859-1"
Content-Transfer-Encoding: 7bit

NBA announces the $2 billion sale of team previously owned by Sterling Family Trust is complete; owners had already approved Ballmer.

--eZDakj4l4DVQ=_?:
Content-Type: text/html;
	charset="iso-8859-1"
Content-Transfer-Encoding: 7bit

<!DOCTYPE html>
<html>
<head xmlns="http://www.w3.org/1999/xhtml">
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
<meta name="HandheldFriendly" content="True" />
<meta name="format-detection" content="telephone=no" />
<title>USA Today  -  Breaking News</title>
<div id = "scoped-content">
<style type = "text/css" scoped>
</style>
</div>
</head>
<body yahoo="fix">
</body>
</html>

--eZDakj4l4DVQ=_?:--

`
