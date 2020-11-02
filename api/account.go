package api

import (
	"encoding/json"
	"fmt"
	"github.com/louisevanderlith/funds/core"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/records"
	"io/ioutil"
	"net/http"
)

func FetchAccount(web *http.Client, host string, k hsk.Key) (core.Account, error) {
	url := fmt.Sprintf("%s/accounts/%s", host, k.String())
	resp, err := web.Get(url)

	if err != nil {
		return core.Account{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bdy, _ := ioutil.ReadAll(resp.Body)
		return core.Account{}, fmt.Errorf("%v: %s", resp.StatusCode, string(bdy))
	}

	result := core.Account{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}

func FetchAllAccounts(web *http.Client, host, pagesize string) (records.Page, error) {
	url := fmt.Sprintf("%s/accounts/%s", host, pagesize)
	resp, err := web.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bdy, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("%v: %s", resp.StatusCode, string(bdy))
	}

	result := records.NewResultPage(core.Account{})
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}
