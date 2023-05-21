/**
 * Created by zhouwenzhe on 2023/5/21
 */

package generate

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"regexp"
	"strings"
	"text/template"
)

var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate code",
	Long:  `Generate code to avoid duplicate development`,
	Run: func(cmd *cobra.Command, args []string) {
		s := `curl 'https://chat.openai.com/backend-api/conversation' \
  -H 'authority: chat.openai.com' \
  -H 'accept: text/event-stream' \
  -H 'accept-language: zh-CN,zh;q=0.9,en;q=0.8' \
  -H 'authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Ik1UaEVOVUpHTkVNMVFURTRNMEZCTWpkQ05UZzVNRFUxUlRVd1FVSkRNRU13UmtGRVFrRXpSZyJ9.eyJodHRwczovL2FwaS5vcGVuYWkuY29tL3Byb2ZpbGUiOnsiZW1haWwiOiJ6d3owMTIzNDYwMjE4QGljbG91ZC5jb20iLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZX0sImh0dHBzOi8vYXBpLm9wZW5haS5jb20vYXV0aCI6eyJ1c2VyX2lkIjoidXNlci1QdXZ5akNEa0FRN291YjVCbDJXa1Q3RTMifSwiaXNzIjoiaHR0cHM6Ly9hdXRoMC5vcGVuYWkuY29tLyIsInN1YiI6ImF1dGgwfDYzOGViZGNiMTQzYTFkZjQxMzk4OGU4NSIsImF1ZCI6WyJodHRwczovL2FwaS5vcGVuYWkuY29tL3YxIiwiaHR0cHM6Ly9vcGVuYWkub3BlbmFpLmF1dGgwYXBwLmNvbS91c2VyaW5mbyJdLCJpYXQiOjE2ODM4ODI4MTIsImV4cCI6MTY4NTA5MjQxMiwiYXpwIjoiVGRKSWNiZTE2V29USHROOTVueXl3aDVFNHlPbzZJdEciLCJzY29wZSI6Im9wZW5pZCBwcm9maWxlIGVtYWlsIG1vZGVsLnJlYWQgbW9kZWwucmVxdWVzdCBvcmdhbml6YXRpb24ucmVhZCBvZmZsaW5lX2FjY2VzcyJ9.a4a0YxQxtVEGwJpcUyAxEBleImqkUJBckSdhio5K1lf7Xt95Wtw3X5Gj2ytX-qm8R4xDdBtGnnq07_PcRvjF3fc19hCRsPXY_ors8nWsD5i2nlAO8aBJ8r-Uoou9VLGEQBU9ceuvj-XZ8P8_gnusQYzwr3bu7JCBzZ2sJ4Zae2-8qVo1dNfDiNuHhnSr4jkqmHiZfzAjGHrAyyWdllIIyXbeDbXtQNBStOG_Y9a8m4JiLoBM2IoY61tb8Oko6mtmtw2YxRp8IYZqoZygYC5pkut1WkubG5oiqI0bWIA3e1a50yzUYBYLqndWOncht5HdbCT40X0wgm5AZhAz5PwH1Q' \
  -H 'content-type: application/json' \
  -H 'cookie: __Host-next-auth.csrf-token=a4ec1baafa6932fcb2ce8f6157b698fbc84848390138411d8613480748c90643%7C6c431f3b2afb1d29661f5f860d684bafdf991303c5b39b5c1c48d5ed268112e5; mp_d7d7628de9d5e6160010b84db960a7ee_mixpanel=%7B%22distinct_id%22%3A%20%22184e575b95f1694-04838835645ebe-18525635-13c680-184e575b960fca%22%2C%22%24device_id%22%3A%20%22184e575b95f1694-04838835645ebe-18525635-13c680-184e575b960fca%22%2C%22%24initial_referrer%22%3A%20%22%24direct%22%2C%22%24initial_referring_domain%22%3A%20%22%24direct%22%7D; intercom-device-id-dgkjq2bp=94a30655-5588-4e4f-8b6b-4a7d2a208b7e; intercom-id-dgkjq2bp=34a41763-a9eb-42f3-9bc5-0b3151b5ccdd; _ga=GA1.1.1002464624.1670292601; cf_clearance=xCnrx2bCHKZRKb3b79oz7N8zjD.55ZmAM1SejuaAEM4-1682669600-0-1-9217aecb.40ce8e81.c516f12b-160; _ga_9YTZJE58M9=GS1.1.1683358833.3.0.1683358833.0.0.0; cf_clearance=s2IaMwEni8jr270XVb8z_G1Ngwi3A3AFovEx3OGROEo-1683775415-0-1-aeb9aa5d.213d5a36.f9b8f5bd-160; __Secure-next-auth.callback-url=https%3A%2F%2Fchat.openai.com; _cfuvid=FI0wVSYF.Vu82Xmlr84UsgrmoR3yXloMjeM6pn51fOs-1684046407047-0-604800000; intercom-session-dgkjq2bp=ZmFEUUIxN25qd1BIdnUzcWVYc1JaVCtxazRNUmplVjZhdjJNdVFtWXhTUWxVaXd4K05seUFaT0E3ZmZQZlBGOS0tN3Vaenp6WHozTDB6c0xySDJTRDlMdz09--7bba9cf2fddacccaaa258947faa73f4792d0d0c0; __Secure-next-auth.session-token=eyJhbGciOiJkaXIiLCJlbmMiOiJBMjU2R0NNIn0..SShl1YWFVkWMJzph.eJ9rRlnmBqw-E2nbDcmmhBJHtn17fNkAENtQ3jpZpGHU168jFrJMjI1OsSlKQKIK7vUtu_f4Pwe8fwassR0XtGOicqK_g8uSe9SlzuOwlcPgKQH8mDAUSFkdCVmoFvIeANmKjRdhpoA15Fq3tGdy1jfRenj0jX5T18a7Y1WtPZAqGcFj2dpjJKThx2ok0US6BjUWRasFMQ_ofVpisWtSATt6j2mTDhWdxLT_Q2ubKIkjzhOSIgk6ATVraXSYmLx-fttSrgaPVQsYtBCP3mwWQI-PN3drkuZTyY1cD7pNIXGev3bhi2uAHeZNjI9X0Q6SzvtzLsYCo5VyeXuyHGo2hB6RPSYeLKLmcVZTwUCvj73z5MCXy9DHEeSwJFEA08CxDp4Eih0OubhafX1Ebq51rgKy-tdCyB4t5PyMHog3v1RyogScDNhRENqaTItzvNI5tM0ad0ZHmd_T78g-MiJLS0rQvtV26BrVBodtE_DXJ9POCpKqxsvDTAFpBvFOMS1oYaZhSAWc6vDEOZr-JU6J_Ac1UNIF9wrb2CqYSaO8Vd0VmH5cAocfGn_owIxPfzaNh6EY6E3O71jXq-jcDJEnSjbX_MaS0UQFO1JYsJKi6tEmS_8GjuOZu-Yt4dfCvsFYuBLKmJ6Y_8q7gADbq-PNQn9WlqyEG6wLCXRlI3R3b1JeLlqUlCWZLr6nJyK-_hEodgP1LarPaqnWzOsWTLvRkAJUxwKBWvKBT7lAQ1s5F00I57Q-0rsAUar72U-lMNMtvGjT2NKzG13VMgbdKbqnssS6-V4ZLRpnh84cNfjF5z7MUGSyvgKSbWwfHznqj3HKBYeX8lQQAD6zq2KDoQOCwo7t9KBUikh_rygn_9bd5AavOXIppKYRmMXaCOxn_uRpnnvIj7ZdEbrT5NBJtpYvIYRH-q9mZ_XwCEGDcAzqtZngonvHeTMKkD2IXvT1hdXCthGKPy9YSpjlEMiFpPbvloUEBGtkpLTf_SHwW8kIXCEMr73GTFcy-YOPG5jFyNxUscPWzvONltfmgn6FcuFDvUvq2i58OMbfEIDQI3oxin2UXob30UqHZ459i_cCaTTnV_zKH8QcUkSQZeImhPY3SMS2BxoKeJA_2k9_2HEr47fWKdWR_2FWAzHcrgoDdnRohv8bi4dK4oSI6ZnVRiSPk7KdgwAbDOWyrVSbMlms3sJbh5epAIglwQh8Kitdf_zszKNT8jxa4hTPN82m_UJc_Q0hxahHES159J2jhRRYcwafkJYlL_EmsgyyQG7vPYGb0Dxz1X5CEAbaiqn1kjgETWOJTEkf8I-Jv6itGy0k0S13zL1-tTvP-pVMLLHAMVqvcc5_5cDpqV6LEl1C1tHf01kGYddzvLJtLC4GfpwZNNX8AufDVJ3nm74lmD609fTkqKtskrAJU3FasLtYPPhAjmncHRGRVgkbZaroBM--sVBruLk2V7HBFHtXrTSb-y8ZpDRb8neSipKuGsC6Oe3DtXs7FFVhXk9O9FQoaosiPoFtwUEpKHV4S8w4vEzni2-8UeULXqljoo3ay0I-R3RxBlqkrJ_UCxu_EO5_nBHbntcabWdvVwGiG_xsLiO4MyxDjPrBz59cGq4vrLQkWWpwWt8d6hdWQQIsxi8rjLs90h9cHkEnP1v2NpKxMO_gkn5n2YHq7i0IEl9U-ResdCDoivLml6p8Q1s17R_1LQXQEEvNYqdSIuJMjvzJLevcl1Ig2cNKdOTWKn6L1bavWb5DHoJRXUdPu53t0k-J8ctPKXUx48jA5wxhIIT5TpVwNh_gyVcpCmgKdu7K9w16ZvB9Q_pWH1fDl7Bkyf68WmuOdooKHDQEQJUPIGp-Gvk1TvMAigtxzFqqgJWkHR0A2czdi8lkkMgrQb-m9mFl5UTUjJPiQemguY_EGoyTyPeB_naxBn-1lid9nTelEREUhF--Wi-FklAkxi7nBMhNQLtIOeBIIMfs8V9pb36b2NsfmHzfTH3dBaketY5pyOn3Sl-Ddw_baq9lACw1_sR2xR3ZPOFSpLvGoAkaBnLPXx7fxpteholtwimQQT-8YPBV7kW_lZ4Z2eJ8nLxKI3W9ORFSmpduOqOUs3CKC1icDk2uR-7BveILZx7A2-a5zlfeMy-4z3XwrA59BzxjrsSO4f5A3xPcNwio_yxwaHsDrtqGiIBM6BuNb7seWL95NZKmnVMV44PSEgoOD2rVDl3qA32imyASs1APzadW9htqu3TF-Ml06lr8A2RnQQNkfcqz0Uawrx-aSY906Rv7LEeLrf7eTky0bODTkd910iS1Y85iJOlcQcPPgqekYy8bojW4AiF55elRmIilJe5iM2txSpdQT6Md-NKLdX29cIa-QBBr8WCOCOTUXz4kQY8y9UTgtBhFO2Ml4ruxeP6xGUx38VVjmEQoRueg2IjqTwUABygEdX562n4KMJOL2Ihdb0hBm50dlSZ3ePpS0CYdzY3bl6pTJU_vVGAUbw4Pc_qLEL0_rlxCZ1VXzABNWM0zaXQlu3gwX7DeCfbYMeLm5wdmAHlOFoVPh2tmWepte33o0ccKKf7cWKnKVST30xc8BZf-5oJyS9nrW1V2CZEvJBlM7_EhQOxkTza-YFJfoYT38jofDLjfmAIGHz9V9mEQyuUVA0MV5ESFrNrb_FMfzMZybAj7TxH139nNqVZM9GuFQwioBMGJnO1aQAxW1NPVWNHkd20VUFZnG2nixuVAuvXOQpmv7WddXlbD_YAvCC9bZbg7_7QHXiD_PeYlgCzT-VdHbDBjZP2YmBCm06-tjEdY4TSoSB_iymFFnN08SlQdVI1acqbOeAbeeZ28mge2FNgOQszGjn-4jb_WqNjioNM.zyBBLfiqz2sIo29C5J_IRA; __cf_bm=TgM7R1r5eipfcObvd7LNNgOBhNBFBRj8MTBAd9zqQkg-1684648304-0-AdQXKB4IliwibNtGxkDBmAHojirll8BAYFWEnQo9FmDh4gKfRm6i83DQIO8SIy9t6nISERMVLDuLg9Mf28xp5yayZm6FY0qPbK4vCVkj5wJq0jAxTcvRX8x0XIfXh9w3CiSfr9Hv1mdu6Majwzpk458=; _dd_s=rum=0&expire=1684649380345' \
  -H 'origin: https://chat.openai.com' \
  -H 'referer: https://chat.openai.com/?model=text-davinci-002-render-sha' \
  -H 'sec-ch-ua: "Google Chrome";v="113", "Chromium";v="113", "Not-A.Brand";v="24"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "macOS"' \
  -H 'sec-fetch-dest: empty' \
  -H 'sec-fetch-mode: cors' \
  -H 'sec-fetch-site: same-origin' \
  -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36' \
  --data-raw '{"action":"next","messages":[{"id":"aaa2644f-78f0-40c6-949b-1f14aeb502fb","author":{"role":"user"},"content":{"content_type":"text","parts":["golang判断字符串开头是不是-H"]}}],"conversation_id":"75210e39-8406-42b1-ac29-51cf13bf12b9","parent_message_id":"efcd4d41-7bec-4c69-a560-81d457f7bd3d","model":"text-davinci-002-render-sha","timezone_offset_min":-480,"history_and_training_disabled":false}' \
  --compressed`
		headerMap := make(map[string]string)
		url := ""
		payload := ""
		arr := strings.Split(s, "\n")
		for _, i := range arr {
			i = strings.TrimSpace(i)

			pattern := `'([^']*)'`
			regex := regexp.MustCompile(pattern)

			switch {
			case strings.HasPrefix(i, "curl "):
				url = regex.FindStringSubmatch(i)[1]
			case strings.HasPrefix(i, "-H "):
				parameter := regex.FindStringSubmatch(i)[1]
				kv := strings.Split(parameter, ": ")
				headerMap[kv[0]] = kv[1]
			case strings.HasPrefix(i, "--data-raw "):
				payload = regex.FindStringSubmatch(i)[1]
			default:
				fmt.Printf("未考虑到的前缀 %s\n", i)
			}
		}
		fmt.Println(url)
		headerJson, err := json.Marshal(headerMap)
		fmt.Println(err, string(headerJson))
		fmt.Println(payload)
		codeTemplate := `
package main

import "fmt"

func main() {
	fmt.Println("{{.Message}}")
}
`
		fileName := "main.py"
		file, err := os.Create(fileName)
		defer file.Close()
		tmpl, err := template.New("codeTemplate").Parse(codeTemplate)
		err = tmpl.Execute(file, struct {
			Message string
		}{
			Message: "Hello, generated code!",
		})
	},
}
