package eazye_test

import (
	"testing"

	"github.com/jprobinson/eazye"
)

func TestVisibleText(t *testing.T) {
	email := eazye.Email{
		Body: []byte(`<html>
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
