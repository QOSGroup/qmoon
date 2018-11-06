package tmmock

var mockdata = map[string]string{
	"abci_info.json": `
{
  "jsonrpc": "2.0",
  "id": "",
  "result": {
    "response": {
      "data": "GaiaApp",
      "last_block_height": "870165",
      "last_block_app_hash": "5Brz2uO+MNE2EQg1DjqXJXnPNxw="
    }
  }
}`,
	"abci_query.json": `
{
  "jsonrpc": "2.0",
  "id": "",
  "result": {
    "response": {
      "code": 65542,
      "log": "=== ABCI Log ===\nCodespace: 1\nCode:      6\nABCICode:  65542\nError:     --= Error =--\nData: common.FmtError{format:\"no query path provided\", args:[]interface {}(nil)}\nMsg Traces:\n--= /Error =--\n\n=== /ABCI Log ===\n"
    }
  }
}`,
	"block.json": `
{
  "jsonrpc": "2.0",
  "id": "",
  "result": {
    "block_meta": {
      "block_id": {
        "hash": "5AE4E524C5291B555B1D07748BB53831B0E5152B",
        "parts": {
          "total": "1",
          "hash": "B0D552846C47268223ACD4660CFDA893A3EBF8BC"
        }
      },
      "header": {
        "chain_id": "gaia-8001",
        "height": "870167",
        "time": "2018-10-26T03:52:43.640425408Z",
        "num_txs": "0",
        "last_block_id": {
          "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
          "parts": {
            "total": "1",
            "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
          }
        },
        "total_txs": "99824",
        "last_commit_hash": "F4DD87D7B9E41B12C5FC6425E5A7EFF84949FBAA",
        "data_hash": "",
        "validators_hash": "FB6D785A83CB40632A74CB05E08B6121D8276156",
        "consensus_hash": "D6B74BB35BDFFD8392340F2A379173548AE188FE",
        "app_hash": "9979166E98DE9BA1A401E7EF4CDCF0F6A6D030B9",
        "last_results_hash": "",
        "evidence_hash": ""
      }
    },
    "block": {
      "header": {
        "chain_id": "gaia-8001",
        "height": "870167",
        "time": "2018-10-26T03:52:43.640425408Z",
        "num_txs": "0",
        "last_block_id": {
          "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
          "parts": {
            "total": "1",
            "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
          }
        },
        "total_txs": "99824",
        "last_commit_hash": "F4DD87D7B9E41B12C5FC6425E5A7EFF84949FBAA",
        "data_hash": "",
        "validators_hash": "FB6D785A83CB40632A74CB05E08B6121D8276156",
        "consensus_hash": "D6B74BB35BDFFD8392340F2A379173548AE188FE",
        "app_hash": "9979166E98DE9BA1A401E7EF4CDCF0F6A6D030B9",
        "last_results_hash": "",
        "evidence_hash": ""
      },
      "data": {
        "txs": null
      },
      "evidence": {
        "evidence": null
      },
      "last_commit": {
        "block_id": {
          "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
          "parts": {
            "total": "1",
            "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
          }
        },
        "precommits": [
          {
            "validator_address": "0077F68359FE12C678980AC3442460CB1CC1ABC4",
            "validator_index": "0",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.189950373Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "xRjc59V0YNEG8AE2TKQUZO6Kvc3WtATNcvwlVvRn2sYMISWTWXBCmJU3OBd8kJ9MRWhFXQQc2PRNvthur51rAQ=="
          },
          {
            "validator_address": "0081A939ABBB733E8DC36E406EC51E4309B9AE3A",
            "validator_index": "1",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.067800178Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "naaMnmwm08vvpNl70s3iEEDX6coVV/mGYEaQ+EOntHmGNpI9t5wNkMbgl81hrihMHXwc9xUyWWAov/aStgipAw=="
          },
          {
            "validator_address": "01F78669F9515FD83DF9250F5C0EE143D3DAD65C",
            "validator_index": "2",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.101701501Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "UBa9ablHj2lL3iTORS1cb0ycgTS/1cYq9AR3Kc1KgUu1/bwCCjUZAwur57rtNYKkvqPaGQDHwU5rOyw4akDVAw=="
          },
          {
            "validator_address": "0548261C50222FD710EB5EBDF03A0E6567B21D20",
            "validator_index": "3",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.130751353Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "uMQUOEoVhj3Wqa4Yb6yutibB0gliI2pVZo01qqERHSYdAwh6mwNW1UYbN673CuVNV6JSNU9cvI+tNwtoKj9pBA=="
          },
          {
            "validator_address": "066C70AF46E5E1B473BD125FD8937EF90E555641",
            "validator_index": "4",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.200738426Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "EMjPPQrTA27PvU5P0WZvVnDSAganGiqJOvJdxTyNC9EsIuV7txvV/ikrxh9KNyApLJ2n0ZiP6Zki6wywJ4YlDg=="
          },
          {
            "validator_address": "0A692DBFFE9E5DD28E7DECC33CED1AEB4C0D014E",
            "validator_index": "5",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.282420641Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "+Wea4LpFXMclMAPDECPDKBNAJkC18T82B/QONpCsjyEtkzczCtK8ZxFhWREA/dw+2ynNvAoo4CXzoZ8bvx5nAQ=="
          },
          {
            "validator_address": "0DDF97111D73DB825138EC15EC27DE5FE8536901",
            "validator_index": "6",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:41.976166483Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "2HaSBIvDMdKljevffA+kJjrvIWisnQPUU2y8IbjuzT2XNW/OafHlzlslBB5QWd2KEvMhrbvqTl/ExSr+9TC3DA=="
          },
          {
            "validator_address": "1ECD8C5FE1341D74CDC633A245DD56858DE2F9A2",
            "validator_index": "7",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.235497128Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "jwPj47g4i8ei7iG1WfmVvdj9O4e3Ugl9ojLlAvN99umBoEtTP+Ec7ItOOSdJ9p/oS6X2/bGccnLcH4QlcEHaCg=="
          },
          null,
          {
            "validator_address": "296869AB66F7003A5AADA1A8DEEEA2913486668F",
            "validator_index": "9",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.005874479Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "4GRjT/dh6GngUd490YXiW5KGhRqSNCsKB80kFoAA0GnSz11hLFZ1y3qc9EniyT4qsRAl21oNwOYT1leeQqmvAQ=="
          },
          {
            "validator_address": "29B93E4EECD8C591AB76C2D52FD7C43CBEEEC50B",
            "validator_index": "10",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.057300716Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "qrhOV6fVPTk68JwzP14LdZK9OO2oX0MSNZRf3goewO5w8HTefAHyKicgKOFGk+/h8FfXWgQfEsaiFCQvJsAnDg=="
          },
          {
            "validator_address": "2E6AF0D5B1A85E173F5EBEDA93D7E7D97A88D06C",
            "validator_index": "11",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.057249847Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "z9jI3Ip2E6HLg4U8mjm8T4Jb0TDZc2baFzHLPGcVu4pHBZzjhllZLMXP8nOfH7AzkgiPYf5+e7DtUEtuoqXIBw=="
          },
          {
            "validator_address": "33412C9A69C60CB0FAA24A5C5B3A906DE24B47C2",
            "validator_index": "12",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.128096872Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "o8noHlLObtT6f92UtCBEsemJCFg9BwXOtkgdg5pYACcCErRIoBxOhvNuwNt1MVZO/XGWuTDlfQk2D76pihORBQ=="
          },
          {
            "validator_address": "3363E8F97B02ECC00289E72173D827543047ACDA",
            "validator_index": "13",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.100909943Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "343kMqgT444LRv1LzLNu/g+F2tkY8dp0RJ6pejwWHWepjoXtYMvBdLeyrURaADWXQt791DB4XkcC67iW80YzDQ=="
          },
          {
            "validator_address": "36E9B7FAEC964F97D84C028ED62493E373AD0B38",
            "validator_index": "14",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.128502084Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "nDafL0XAt/RC/Gs/il2rECqG9i0pvDvsJy4fQnAsLUmU9xdXT+ZcVckXAl5COjCB/PoVr02oTIuCT1ZrlvfoDQ=="
          },
          {
            "validator_address": "398610C4CF11C84C89AC6975470972EE75DA17E4",
            "validator_index": "15",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:37.293349454Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "hEYkSaTsSK+AICtdAJ7WUddHoXQFDXB8ObWGVD3P+gzE2YvjUpf5E5NfAOUXNsoRm7msAb79njxRSbea36NZAg=="
          },
          {
            "validator_address": "3AC99B708404452366E0953A4BE849226DC86CBC",
            "validator_index": "16",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.202258147Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "DsxcRivUh7LCBgoGtgsYcKy2vqxe/WCNA49M0Cp4K1e9OjvX8Q5HR1ElBtotFTw7Oe9IF+IttSHwVQpLRAGJCA=="
          },
          {
            "validator_address": "3E899BA36A93C7370347BD5FE85909F2ECD93F8C",
            "validator_index": "17",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.115752723Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "IxsG9Ll0A4fjEuNyRiy9UsT914bMVeNkRmj8d6fNfSp6YUSws5y9icvUuMJi10n0GiB4BFU58Qlw7cDMkNOgAw=="
          },
          {
            "validator_address": "411C430038318FE4DB32AFC508A6C12782161294",
            "validator_index": "18",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.100518689Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "EY16j0s4VSgJWc4+1Ny8QRubWRHrBP6oQMyhX1kaqRFdGvV6tCKcE0o3UXD7eorQqxbGg7X9WTR8IeKEedlgBQ=="
          },
          {
            "validator_address": "4944F8BBBB5303E17BC0607DF332320EDBC750EA",
            "validator_index": "19",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:41.999978795Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "FSDiabh6EpczotlMXLAq2NplgFPlWH34AmJfBB/fTTc7ggmaktQrvdBlJ0acnjbjV3J32g4XLIc3BLh41PniDg=="
          },
          {
            "validator_address": "4B9C6FF8B36EE6134AA8D822772C0162C471186D",
            "validator_index": "20",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.100321876Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "I3vvAk/Zhb/k2UvDy2IaG6MncxDfUsDDg8aOuxRa3sppr1jY9CjApZ2Mwg2rbIcpBjn5qO8O/L9Sy8b8SyOiBQ=="
          },
          {
            "validator_address": "50C18305F49BE332E4C6BCDE0DEECD4A22B1D6A4",
            "validator_index": "21",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.139732103Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "WHV1hoo52mvEuEn6jmPBfwCfIg5NcXEOi7RlYENz78x+9yY+u28Jpp3sChjNWaOYD6KgGhFm4x3ln6F4zP97Dw=="
          },
          {
            "validator_address": "5104F488EA6B1E2C33CBDA3C18389C88104DBC72",
            "validator_index": "22",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.029139417Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "jFCTELShgHjP7YoMJa8oue0w7gknl8YQYX/HN2bpdHpsZCKDNhtVu8Qs8sprL13RxGuiHGkiae9Qcunho9kxCg=="
          },
          {
            "validator_address": "570B6A755ED339EAC56BB3390A48069F6AC9449B",
            "validator_index": "23",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:44.06663216Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "p99vXUa/fRwrOo95pu/NWGBxm48VLAMrQI3Lxn06z6v5EBtT3r4j5K9JbKoEq2QNxa4x2/+Xg9xUNfZKiVr+DA=="
          },
          {
            "validator_address": "5B9628695CDFC6025857578114E3D3126687EAE1",
            "validator_index": "24",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:12.614433636Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "7QQ1eHpQfe17dbpDY/rB2AGo0Zq/aNRfGfiPv5r4/ICcKaDaT38rG1V0EBOSRRBmrsUM3oCGCt2memHQzkEVDA=="
          },
          {
            "validator_address": "5BDF8EA41B1835FBAD865B181A2A4B23F511B5E4",
            "validator_index": "25",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:41.990302733Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "LDg/I9XREXlbRNSc3JpZ7JpNOVtrI0jlRWsModEx1p0bLK7X4C4PzCJnSPzssBDo5LDL66rmv1BUl/lwNk9iAg=="
          },
          {
            "validator_address": "5DE9674DA2D535D5B1547D90990E9087EAE15E39",
            "validator_index": "26",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.07235729Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "7hQHHu9cuhHYfVNSlz8anndkQ2bdyFD8xw+Frp0ssdxqyy3VYTzsRq7h+BQ2WoKzb0Thb0bxDvR5C0RqB23tBw=="
          },
          {
            "validator_address": "5E06FB65FD0E0EBB11D0D74A1E480B4A8D7D0FE7",
            "validator_index": "27",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:41.965453549Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "xScWQobpDi0lxC/FJHpAtWHumQxizcUax/L6k5yw3C9qV0oRAtaaYplwRvZZLoRyhRRxb4d5jekl15R3bSX7Bw=="
          },
          {
            "validator_address": "5E8673673E37450F01B64138FBF4B172BDB52D68",
            "validator_index": "28",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.081514045Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "JoLjWHPBJMc1uwLH0qOjJfmeFI94/BOYGvbdps2hNTkAHiDrvUT9w6X75mpyAGTvmdeO61fULoh8l5+gYGpTAg=="
          },
          {
            "validator_address": "651995D8471D07A22C149265F5237A89149795C1",
            "validator_index": "29",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.018934378Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "UiYCrvNjvDFm5HV+Vu1nRaEYzdWtOKfkPPpafy8g4xL6w4hm2zyXGW1breb0dVubIB2bwT813lOYmQTFEufaBw=="
          },
          {
            "validator_address": "668309BE1A6F5317197B0BC6E47153D249DD5DE9",
            "validator_index": "30",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.098763494Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "XYU7bPKqeit5ZgvKKCsjKFCYDQNGFo+Amv8gM1+camySy/bi+mkIXWjDgAUXQ62D/1X5fle6uW3KmFNfxZj5Bw=="
          },
          {
            "validator_address": "69D99B2C66043ACBEAA8447525C356AFC6408E0C",
            "validator_index": "31",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:41.996341936Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "OB3vB574218TMn7dRMn7KzJjYIUOhMOsMndxFPGs2+KN9rENKUQ0yk2VSrA4gVeafkaWJaRwTMkF+2tON0zLCg=="
          },
          {
            "validator_address": "69EDADFA4A803E9DBF548C900D4A465E7D154262",
            "validator_index": "32",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.223120502Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "M7uNIMHh4px4Ncsj9g2cseVWCjvgwBgERFXEPpASOXxyxM/d9wsq+UOlCOc9TJ1VOfdOzucmpYoYKPjItpINCg=="
          },
          {
            "validator_address": "6B275FFA0571EEFC4482CD7A6CAE836F5785CB4E",
            "validator_index": "33",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.03667375Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "vr1KHNuV2q3QZLDwjLJqR01dsUQjQuHvWbL4HTe9qeJ7lbkRTBVmF1lhMuYGpCsnr4BLDTGq6mptSnrtR2/eBw=="
          },
          {
            "validator_address": "6D3AD380E398D8592B03AA124BAFED5F125DF863",
            "validator_index": "34",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:50:54.681101074Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "0ZNEJWyBqsxAepp3TX+xzvchbf9a0Vzr6c2O4C7fQSo9yK4ONFqi7e/JzYO3tGPEiCz68F6OPy0bdntN5AbHAA=="
          },
          {
            "validator_address": "6F6C93B8BF3495063FEC9B5C771158607F4AE9D6",
            "validator_index": "35",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.190635843Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "hOrsNKvcM7F3tqFoJXIfBySela1uK63jEqQPVBANPqIes6TThPWC3n7pDxKtKql2NmpsD6XndoNcfgP7p/iDAQ=="
          },
          {
            "validator_address": "71BC71D28E4599BCA22214BBB2CE0DFEE90A9EE4",
            "validator_index": "36",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.041132395Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "U//xPD1HePtO10BJh/4mzWfevNUAxyrN/vq/ze0GQCip/fvfhhQEDs7tVINUErFy4wK2Z4oJJjWIhIZGCqEhDw=="
          },
          {
            "validator_address": "732CEEF54C374DDC6ADECBFD707AEFD07FEDC143",
            "validator_index": "37",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.016726928Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "qQCG30vTzkgkC2BMVj7Bw657i+Xgw05daoHIQVLaNvokVYiBckwe/R8w1TfVMpoBvrLvdOs6E5mVVJs1BJdnCg=="
          },
          {
            "validator_address": "7A20421F99FC4362B13C2FD43548F3E75B057017",
            "validator_index": "38",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:44.706679123Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "bkS9M25iVKhi14rG7sVX9UUPJ817t4aMB0Frf7AwTidsqBMDAll01xK5JIUOYaWTdduH2se0dWTKJqw5i3MgAA=="
          },
          {
            "validator_address": "7A32D912F7A790EEE5610653DC4DDBC0DB4F2497",
            "validator_index": "39",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.080837884Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "Sz0wNrcSej2S0Y+93wLfz/rA19DKyv3FUQs3aa0qKKvs5TeWFocIy8yj20tJfSST3Ssi3ngUOvk4Gq+pp1b0Dw=="
          },
          {
            "validator_address": "7EB7593F519F1180C7E872D0494611464EDF03AE",
            "validator_index": "40",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.197659052Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "MdX5MdYXqJ/1+XgwI+aKLweaAcbYbBEji/OzZTSjKVfIrxl/hb7SjSZZUi+f0oRozxrXsNwjmgmXTeSdKiQRBw=="
          },
          {
            "validator_address": "807CF675BA8479FC4FEC1047215DF07970040F1F",
            "validator_index": "41",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.269727839Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "TNUXDqr7jULG6bfHiNw+XL8n6+xxoOUvkJoi1yTtczf261S78o5EiJNzG8NLDcTG+9LomMNxt4hkeZrBwSfzBw=="
          },
          {
            "validator_address": "8340F82D32C4A221C57FC33FCBA44309C4DF1BBD",
            "validator_index": "42",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.131325444Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "YW1me/jfIgTbDsp+GXCI9x+WArhvbuEuxkGE46dVP8+9WSr2u4KErQJGToNZXV+mFW+ATi/u222TyNY0Zz77Bg=="
          },
          null,
          {
            "validator_address": "850D387A41ADB21F7D67EAC5076D594DF3D861C7",
            "validator_index": "44",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.165995379Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "RtU8kUArqL7OBgAqka1WGQ1P9jtdmD+OvIoDo6T19aN+whyQyGvfk8cSK9diA3i0wDrzGfg6BZ9J8CmzvnGMBQ=="
          },
          {
            "validator_address": "85335B5077FAE8E2926AD338634079124EE1C4A1",
            "validator_index": "45",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.116412347Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "xM/CG9az41pARqLk4xGegDDtKMT031SyvAqU7AiHZy32q4EFgH1vqX75vYLr6uNeiW58yqfO82o+msr47y/BDg=="
          },
          {
            "validator_address": "86DDBEEF35D35786A68834B34869174AC314BB74",
            "validator_index": "46",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:41.993818104Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "fmw9t97V6ebjrGYU0esV5S5e6ZAzmSFCcut/0jhma8RaFe7M6kCkGXbK/s8erZYIu0JHGe6uqvKPUEEMltlYBw=="
          },
          {
            "validator_address": "89B357643A945BD433CB5B2B9CBD9B90CB9CCD5E",
            "validator_index": "47",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.038224523Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "ilNyX2m/gtfPG2QNi/dbo/1ROcDTJ0dpAgsIrmHCrmIHpChLSAczN+BT3UUJXdxs56KahVBTrPCEJuHAVbhcAw=="
          },
          {
            "validator_address": "8DF81F7913F1EBF856618B1475C12EEEC1214BB0",
            "validator_index": "48",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.056282558Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "96wb0VQWJnevrcpYsFc952zXkC/521zEmgJXpMp2/EATa0sMnQZ7LXgYc/yg/qKRbj+SdzkPakPx4QFLP3sMDg=="
          },
          {
            "validator_address": "8EA62D00960B7557560E0D78278B6B36165AD0A8",
            "validator_index": "49",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.199660098Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "yCilkJGljGmQymG4VmNBDu809Sjh+IROAsewuzr5/kBd9wazHH8K+f1iKAdTJ1xE9rmtrLYYeEst9ctAxG6zDQ=="
          },
          {
            "validator_address": "8F06FE1FD042683C426137FBC4BBAA4288924217",
            "validator_index": "50",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.01093748Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "OAVJHz8mhD1UdnHgYUbz5726c5gE8v2+AWitwmLnwnlFnUrhJvDYqrn7CHiBjKVVQnhfwqTJBkhHSFnvU5MxCg=="
          },
          {
            "validator_address": "8F302C08318444DDC43EABDE96E869D0B1F5A421",
            "validator_index": "51",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.026036164Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "XJfQ1TXFRYjQN4S2mZF+AMe/SHc5uJ3CFNOTjN2vyzXXMv5iAbx2dJ+GTkpkYrQ73DgP9xA+oodFoWuprSWUBg=="
          },
          {
            "validator_address": "927CC759310F4DF5C8936EF6D666ED3135D609B7",
            "validator_index": "52",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:44:41.919412991Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "sROtH4UZy9uT9yhNCt3nnHIzlnXfWxyIov1BmIMQq5qU9CSbbh/OAMN41dwVrrUKdskIEpFeIsWNHl14+zYnCw=="
          },
          {
            "validator_address": "9524C2BBF5C9BCC7AFEE121AF03E31D8480DD5E5",
            "validator_index": "53",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.012608011Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "QD3zI7Q8eCQF7pgJyy3/+3kEi9a2T+u7AUKediUryHpmewMgUqXhP46XgJX+DszIg2Slt2duU+J3RPzKX5seAA=="
          },
          {
            "validator_address": "96BDE45EC52606CF5CB59853097871AD60680549",
            "validator_index": "54",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.007833207Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "0A+eZ8V2ZtRyXlCmC9uF/TgvIDih0NzCOkZGJOB2Woe1RNfM/cHNn4VqUBTuJIyf/3ziKj9COFppVIgN1DHfAw=="
          },
          {
            "validator_address": "989E892A603E6BA7F325BC52555DFB47A8751E66",
            "validator_index": "55",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.164627682Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "+IZFjjbzM7re3NyC4M1t8g0vAy7fGbvaS5uUXZ0mzOW8c6pqobDB0EYaS6+akhqik8cL+xGy95+tmWcTftzhCA=="
          },
          {
            "validator_address": "9D93AD02FE1BD7E76B9E58D7EB461088EAD011E5",
            "validator_index": "56",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:41.993763921Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "2E1QNaNl50ck1qa48Au8txo5z96REwEmJOblJcuX9amRMAcpLWVlMnNaXZ3q2653CzJ2B9d/rMzvil70FNeBCw=="
          },
          {
            "validator_address": "A0ED8A257D7862094FE8015BCE2135D111722DF6",
            "validator_index": "57",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.100837729Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "b5KQ7I98LI9XkD2l4WKsAq6qWeG7VOeDbc8aH1gwny1+3GPeEYsAxq+wPAbfW0JtUXgZwqjiqAew8KnFutV6Dg=="
          },
          {
            "validator_address": "A5868971300F155020634D7BB49C7634F7AACCD3",
            "validator_index": "58",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:41.977351828Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "iGOEd4yTr+X0ozlTMLlUsszhN1SeihUuiSKZSyPgdfwJHSDkzJ7AK0fh2TJeS6+iUnhIwbBM5pOHRlzzPw4yAg=="
          },
          {
            "validator_address": "AC96444A9083890B7C7B430DB17BF1D3EA5EED04",
            "validator_index": "59",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.25660217Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "gUshjhAFn8xIqOdgpGYPiTlRVkosoebdi1sxKEF6/BP1XjMIgDaPAiDr8bx8JlFKDXKZhrIe6kQ8rRi/DtDWAQ=="
          },
          {
            "validator_address": "AFCFB9B8DF23559CB5EB97C7CF1D77B411B2F47F",
            "validator_index": "60",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.080357596Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "P+odoXKwSgi1ak2wVZevvKBdN1Rcy5HvZT2sKBjHrxqmHsS0oJRQHzFeNOVhc6RA7W4u0dPYX4NzFkErGy5lCQ=="
          },
          {
            "validator_address": "B0155252D73B7EEB74D2A8CC814397E66970A839",
            "validator_index": "61",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.006185788Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "BRj7PJs9OdC+9Eul6XqA6vMqz7Al4Tj4739j511d5+teiUrKRyNsc3SljUd/jppewX7aSosvZdX+duDtytwTAA=="
          },
          {
            "validator_address": "C39A8A2B85198DCB8F23AD1F6C198E2BA7AB5D60",
            "validator_index": "62",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:41.985158559Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "xbvq41fZMrvRF12RslvLIiNK48pWDfJAn8MRRmQxkjqUzo/TANj5AEgG1Rd80PVrY34ABS+srTWatDkVqdkHCA=="
          },
          {
            "validator_address": "C65E1E696C038A9A4A3565014ED168E97E7E63DF",
            "validator_index": "63",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.055016597Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "0tnHvNeQbqWtq+sLyOYijyfoAGpypZoyCMJQvKKWLx5N+QoqpJ0mE3lYrKHfg0fo7VBWOMDWPFqxPlbFW1VvAA=="
          },
          {
            "validator_address": "D31B1198AE055BC52C600C19751011D95A64298A",
            "validator_index": "64",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.168394334Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "IjPnnq7O9/WXQD1GHz+0n83lGCNwKBxvUh3F+qL12qSkFQkRBVtEm4eI0bqQegcKZAzkaUPAPf4mZslrBR9mCw=="
          },
          {
            "validator_address": "D6285C3422C8886445D53757CBE12D2B91B5A9D0",
            "validator_index": "65",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.094849005Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "p4SP9+qeMj7psLXDG0I34uxeS8iJyf/43xv2rmlPhS++1XteiM0CAE0/gM+0RTGkjY96r6C2p9LCa+K/+wbpCw=="
          },
          {
            "validator_address": "D77FB84A729972CD3E757C9BEB61561855832B92",
            "validator_index": "66",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:41.97533167Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "+dybZUeCsFYgpHw+F6XGFozHQIGy5tkJF6J8p9lvRBuUK5rHyqq4rzg0QuVI4a4s7swFIRPldoj62J1SAvG6CQ=="
          },
          {
            "validator_address": "D981CB6DDF336E311A33C585C23C10F9F033A305",
            "validator_index": "67",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.05154154Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "TuVRt1wq6698MBvLRPKU7EmsLOUbt6uN1oBHi0Tjs+ViNf6yQtGNCJ9pM9qvIdwJmtDI45b+e8s008chQ/BiCg=="
          },
          {
            "validator_address": "DA38013D7BCC3EC46109F8CDDF14F8D080CC4A9D",
            "validator_index": "68",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.005390795Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "MAMLlWTGRTusSwmus2DG6w0FHte97OUnqqCngqxw6phzN4OCK79+uakCfwhXSQ4a8q3pGeOy3CYTiMA9T8XKBA=="
          },
          {
            "validator_address": "DB71F2DC8C967F0BABE9EF797760F17EF3D99ACE",
            "validator_index": "69",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.077737592Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "Rl8gjUlCgT24gcjpHW+McXe4Q4Z0AYUBTCEx/6FcdcbqP7zB4N4pm7NpXt7pABPol4HkY8I+qXkiq3S7k3ETAQ=="
          },
          {
            "validator_address": "DBA70FA7E9D55E035AD87B41C4DC0C38511FD09A",
            "validator_index": "70",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.149066786Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "fbFI8pvSgJibWoEEcx53vc5HxLxmAstAyr281mhM5PIOXV/dECMrapzME0uU4eObX6okk+uL6J0ZR5H9+saNDw=="
          },
          {
            "validator_address": "DDDC63E5B9C7D1509805601E2EF6934BAEAC9115",
            "validator_index": "71",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:41.965200538Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "cAoDrF6YpYOZ7niWVpk7EixfjLolyZllB7/PXZ1txSIQVoqPTbn/6V+ByBOBOicDWHdzolBrsOXUSBy/kViZDw=="
          },
          {
            "validator_address": "E39B5778E0D4297A46F45CF6940B0A163A430FF8",
            "validator_index": "72",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.006840979Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "WqmhEKhJh+xfHkQsvDX+jS12z6V/ETc6xaLfJQ1o2ftmlM4M//D/6VIMlsKThm+CbN7TtrXbILrhjg+y8C9qBg=="
          },
          {
            "validator_address": "E3B19738F7132F950668504D326C7660F159D170",
            "validator_index": "73",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:41.95015111Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "yZ4dcc60x8iGhKfbubvCNGD5cnQ5yz7cxOzIwdi3k7t7dqxNsMB1SwXQ6hxfqrptK4Nl93wt9HysJ9j7Aw/SDw=="
          },
          {
            "validator_address": "E42F8AEABB685219F7027F54D2B0DBC7846F633A",
            "validator_index": "74",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.021793181Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "nY1nl6sE0VT1uuuzwZp3nZh7Guf65t7dFOlvZjwBgm/N+0RjJdzbvKqZ2KL1Cep3GC6jqtY5HQjLu9D6yYmaCQ=="
          },
          {
            "validator_address": "E4E6C2E81269DFFA196C17F98240E491479A6008",
            "validator_index": "75",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.06499155Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "oS4vk8HnL0sDG0AEuD1ILrSF9o0yQLN0Ck3UYsTfc+sE9aj5JLhs16gOKfd3FCOb//qQ45hpoNR+2CF7eLbvDQ=="
          },
          {
            "validator_address": "E50ADECD5FD27FA2CD610C07F8AED36A2FC4A6A6",
            "validator_index": "76",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.033090607Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "CeIaESD9TN/Cm79yukwjHPKRMP0oOS90f70Z7xzD1sZdyL1vCFvEd3lHOeODlpxlspdCDOnKEurj3/jKfi4bCQ=="
          },
          {
            "validator_address": "E56E00A6B6AC1BE546CE9A9F34C7C1BB7585A8C1",
            "validator_index": "77",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.038447217Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "AlmrqN48+tJqU3sirgqq39+yysIjJVcnVqN8qVkgnfx9rqpK3ZdUUhIVAOhkbFIYTUzPpm8eh6u+OS7xxSq2DQ=="
          },
          {
            "validator_address": "EE67A5CA3269434AFAA41CE43E17B467488B5210",
            "validator_index": "78",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:41.988588971Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "vN/gteKMq8Mq6H1tUkPWSyasPUbzpxf2Bmh/9D8xR28ksgSPpFKQD/dcKvBrXutvrmSvhU9hQRmzE3bW+fhyBQ=="
          },
          {
            "validator_address": "F23FF36BD5B90C33CE3A03ED72DBDCF5EC07D6AF",
            "validator_index": "79",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.059432589Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "5TLNZGC+I7KsUTaPnatZ3RA8Q5zZ6RUcVtE70AXlZcGyv3bY9AuQUWbra130tzLLoAmYUFnjl+B7CbtNOoeSCw=="
          },
          {
            "validator_address": "F4268597DE93E3E2CED38845010F8AA34E2A879E",
            "validator_index": "80",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.021015532Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "FOZUqXNsKy7diEH4dGVJ1/+d7b697p1NfLFDVHQVZz6/uy3k6tRW+VbD0nh7smUac+IhyLdO+B+ZmyZ0nP9MDw=="
          },
          {
            "validator_address": "F6738260186D33D9C14FC6E7017AFE6BB952A63D",
            "validator_index": "81",
            "height": "870166",
            "round": "1",
            "timestamp": "2018-10-26T03:52:42.002597132Z",
            "type": 2,
            "block_id": {
              "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
              "parts": {
                "total": "1",
                "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
              }
            },
            "signature": "ESLWfcavlEtTRq7qKA5pwHeh77CqTbpLZvebDocwm95FxKjuquGtD/GD+cVHMxEIyvHQXSt4XkqMLQjcrM7XBA=="
          }
        ]
      }
    }
  }
}`,
	"block_results.json": `
{
  "jsonrpc": "2.0",
  "id": "",
  "result": {
    "height": "870168",
    "results": {
      "DeliverTx": null,
      "EndBlock": {
        "validator_updates": null
      }
    }
  }
}`,
	"blockchain.json": `
{
  "jsonrpc": "2.0",
  "id": "",
  "result": {
    "last_height": "870168",
    "block_metas": [
      {
        "block_id": {
          "hash": "27F56F86AB02B26B77011FBFBFDBD185555D79CE",
          "parts": {
            "total": "1",
            "hash": "208F863FB187966AA38A154AFAEBB42B18886B2E"
          }
        },
        "header": {
          "chain_id": "gaia-8001",
          "height": "870168",
          "time": "2018-10-26T03:52:52.942673199Z",
          "num_txs": "0",
          "last_block_id": {
            "hash": "5AE4E524C5291B555B1D07748BB53831B0E5152B",
            "parts": {
              "total": "1",
              "hash": "B0D552846C47268223ACD4660CFDA893A3EBF8BC"
            }
          },
          "total_txs": "99824",
          "last_commit_hash": "5837E1B6DC16D5948ABD179F8273518494D6258F",
          "data_hash": "",
          "validators_hash": "FB6D785A83CB40632A74CB05E08B6121D8276156",
          "consensus_hash": "D6B74BB35BDFFD8392340F2A379173548AE188FE",
          "app_hash": "5B62944ED04CC460BE9A3945B665508E5287537A",
          "last_results_hash": "",
          "evidence_hash": ""
        }
      },
      {
        "block_id": {
          "hash": "5AE4E524C5291B555B1D07748BB53831B0E5152B",
          "parts": {
            "total": "1",
            "hash": "B0D552846C47268223ACD4660CFDA893A3EBF8BC"
          }
        },
        "header": {
          "chain_id": "gaia-8001",
          "height": "870167",
          "time": "2018-10-26T03:52:43.640425408Z",
          "num_txs": "0",
          "last_block_id": {
            "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
            "parts": {
              "total": "1",
              "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
            }
          },
          "total_txs": "99824",
          "last_commit_hash": "F4DD87D7B9E41B12C5FC6425E5A7EFF84949FBAA",
          "data_hash": "",
          "validators_hash": "FB6D785A83CB40632A74CB05E08B6121D8276156",
          "consensus_hash": "D6B74BB35BDFFD8392340F2A379173548AE188FE",
          "app_hash": "9979166E98DE9BA1A401E7EF4CDCF0F6A6D030B9",
          "last_results_hash": "",
          "evidence_hash": ""
        }
      },
      {
        "block_id": {
          "hash": "18463C2EAE377394E0CF136CED0F6E8D34ACBD6E",
          "parts": {
            "total": "1",
            "hash": "F6E3260F13C395E2D3950F1990ABA9D8D4060179"
          }
        },
        "header": {
          "chain_id": "gaia-8001",
          "height": "870166",
          "time": "2018-10-26T03:52:41.135084143Z",
          "num_txs": "0",
          "last_block_id": {
            "hash": "68B72AD48FD67C747068ECE69E86B8CE8C0D81B2",
            "parts": {
              "total": "1",
              "hash": "246A04488A62075A03D4D46015C2F99650A304E0"
            }
          },
          "total_txs": "99824",
          "last_commit_hash": "19C62149631FB47031F668E6636FC197ACC8B021",
          "data_hash": "",
          "validators_hash": "FB6D785A83CB40632A74CB05E08B6121D8276156",
          "consensus_hash": "D6B74BB35BDFFD8392340F2A379173548AE188FE",
          "app_hash": "E41AF3DAE3BE30D1361108350E3A972579CF371C",
          "last_results_hash": "",
          "evidence_hash": ""
        }
      },
      {
        "block_id": {
          "hash": "68B72AD48FD67C747068ECE69E86B8CE8C0D81B2",
          "parts": {
            "total": "1",
            "hash": "246A04488A62075A03D4D46015C2F99650A304E0"
          }
        },
        "header": {
          "chain_id": "gaia-8001",
          "height": "870165",
          "time": "2018-10-26T03:52:31.833769183Z",
          "num_txs": "0",
          "last_block_id": {
            "hash": "FFFA1EA15ED5973DE27004F9957D4DC5DDED1F8D",
            "parts": {
              "total": "1",
              "hash": "A3AC7A181D8B79FA421EE8164B5B16DC2356CF2F"
            }
          },
          "total_txs": "99824",
          "last_commit_hash": "6151EE9FD9E09F45BD41AB0BD61E3B313C339534",
          "data_hash": "",
          "validators_hash": "FB6D785A83CB40632A74CB05E08B6121D8276156",
          "consensus_hash": "D6B74BB35BDFFD8392340F2A379173548AE188FE",
          "app_hash": "99FEB47B99EC8DE52C860BFC25245FEA110CAD5A",
          "last_results_hash": "",
          "evidence_hash": ""
        }
      },
      {
        "block_id": {
          "hash": "FFFA1EA15ED5973DE27004F9957D4DC5DDED1F8D",
          "parts": {
            "total": "1",
            "hash": "A3AC7A181D8B79FA421EE8164B5B16DC2356CF2F"
          }
        },
        "header": {
          "chain_id": "gaia-8001",
          "height": "870164",
          "time": "2018-10-26T03:52:22.754732623Z",
          "num_txs": "0",
          "last_block_id": {
            "hash": "8C1557DCD51412D92A1B8EDD3431B295B0427AFF",
            "parts": {
              "total": "1",
              "hash": "67F05931E635CE0FA72E3B0E5ACC8C7ED8860CAC"
            }
          },
          "total_txs": "99824",
          "last_commit_hash": "E465BB3359202ECB7A188CCE71EEEDC7A77A668F",
          "data_hash": "",
          "validators_hash": "FB6D785A83CB40632A74CB05E08B6121D8276156",
          "consensus_hash": "D6B74BB35BDFFD8392340F2A379173548AE188FE",
          "app_hash": "58B0F4D9CB1E73E86697DFB72D787B710B11D811",
          "last_results_hash": "",
          "evidence_hash": ""
        }
      },
      {
        "block_id": {
          "hash": "8C1557DCD51412D92A1B8EDD3431B295B0427AFF",
          "parts": {
            "total": "1",
            "hash": "67F05931E635CE0FA72E3B0E5ACC8C7ED8860CAC"
          }
        },
        "header": {
          "chain_id": "gaia-8001",
          "height": "870163",
          "time": "2018-10-26T03:52:13.744996626Z",
          "num_txs": "0",
          "last_block_id": {
            "hash": "C4417BA8A6F0E97D5E88389CFD244204D807D48B",
            "parts": {
              "total": "1",
              "hash": "607242E8C3E07E6276743AEFF6F426652ACEB3DE"
            }
          },
          "total_txs": "99824",
          "last_commit_hash": "C5DDE1E9127036BF7965FD21459D29F04BCDF725",
          "data_hash": "",
          "validators_hash": "FB6D785A83CB40632A74CB05E08B6121D8276156",
          "consensus_hash": "D6B74BB35BDFFD8392340F2A379173548AE188FE",
          "app_hash": "098AE8C0C39060EB227D8F5809D007A04A8E5428",
          "last_results_hash": "",
          "evidence_hash": ""
        }
      },
      {
        "block_id": {
          "hash": "C4417BA8A6F0E97D5E88389CFD244204D807D48B",
          "parts": {
            "total": "1",
            "hash": "607242E8C3E07E6276743AEFF6F426652ACEB3DE"
          }
        },
        "header": {
          "chain_id": "gaia-8001",
          "height": "870162",
          "time": "2018-10-26T03:52:04.388883032Z",
          "num_txs": "0",
          "last_block_id": {
            "hash": "E776663F621787DC9062A47158534AA636C0B032",
            "parts": {
              "total": "1",
              "hash": "4AB432D9FFA13DFF05D53712E74CE9EA9A185558"
            }
          },
          "total_txs": "99824",
          "last_commit_hash": "688EDC176292CE7A6EFECAEE8B56401F04329E63",
          "data_hash": "",
          "validators_hash": "FB6D785A83CB40632A74CB05E08B6121D8276156",
          "consensus_hash": "D6B74BB35BDFFD8392340F2A379173548AE188FE",
          "app_hash": "83E32F29649DF4F411C82A50BA9ED6EE7CEC84CA",
          "last_results_hash": "",
          "evidence_hash": ""
        }
      },
      {
        "block_id": {
          "hash": "E776663F621787DC9062A47158534AA636C0B032",
          "parts": {
            "total": "1",
            "hash": "4AB432D9FFA13DFF05D53712E74CE9EA9A185558"
          }
        },
        "header": {
          "chain_id": "gaia-8001",
          "height": "870161",
          "time": "2018-10-26T03:51:54.950433093Z",
          "num_txs": "0",
          "last_block_id": {
            "hash": "5AB6566EDADFB6D59102A61FB9B94AE8F525D66B",
            "parts": {
              "total": "1",
              "hash": "74D94531570EA79FDD31CA132E0D37D419B33081"
            }
          },
          "total_txs": "99824",
          "last_commit_hash": "9E539DEA87E88461D3DFF2C8B9299D8D9568CD5B",
          "data_hash": "",
          "validators_hash": "FB6D785A83CB40632A74CB05E08B6121D8276156",
          "consensus_hash": "D6B74BB35BDFFD8392340F2A379173548AE188FE",
          "app_hash": "20D556AEE3FF300226573409CA807113B43B46F8",
          "last_results_hash": "",
          "evidence_hash": ""
        }
      },
      {
        "block_id": {
          "hash": "5AB6566EDADFB6D59102A61FB9B94AE8F525D66B",
          "parts": {
            "total": "1",
            "hash": "74D94531570EA79FDD31CA132E0D37D419B33081"
          }
        },
        "header": {
          "chain_id": "gaia-8001",
          "height": "870160",
          "time": "2018-10-26T03:51:52.432426946Z",
          "num_txs": "0",
          "last_block_id": {
            "hash": "97C046065A4ADAB5AB20C3A1D461240621D5360E",
            "parts": {
              "total": "1",
              "hash": "E4CE8D8563197190CFE68D06A757988261F34364"
            }
          },
          "total_txs": "99824",
          "last_commit_hash": "F09A753E292E476D25D7439BCFF44D0F7F8C4386",
          "data_hash": "",
          "validators_hash": "FB6D785A83CB40632A74CB05E08B6121D8276156",
          "consensus_hash": "D6B74BB35BDFFD8392340F2A379173548AE188FE",
          "app_hash": "236A8270D66DEE8C9C7D886190BDC4D12592137B",
          "last_results_hash": "",
          "evidence_hash": ""
        }
      },
      {
        "block_id": {
          "hash": "97C046065A4ADAB5AB20C3A1D461240621D5360E",
          "parts": {
            "total": "1",
            "hash": "E4CE8D8563197190CFE68D06A757988261F34364"
          }
        },
        "header": {
          "chain_id": "gaia-8001",
          "height": "870159",
          "time": "2018-10-26T03:51:43.395277352Z",
          "num_txs": "0",
          "last_block_id": {
            "hash": "EED959EB2A13EA804C50BD9AC7F0BCB24C542ECA",
            "parts": {
              "total": "1",
              "hash": "E9C103F338E0FF26CE4EFC0226E2B3CDE3632CEB"
            }
          },
          "total_txs": "99824",
          "last_commit_hash": "1EECFDCDC13A036FECB3B46FD82A57636DF6DBD4",
          "data_hash": "",
          "validators_hash": "FB6D785A83CB40632A74CB05E08B6121D8276156",
          "consensus_hash": "D6B74BB35BDFFD8392340F2A379173548AE188FE",
          "app_hash": "6D199B148D6CF1B46E222FAF1519165D1F01439F",
          "last_results_hash": "",
          "evidence_hash": ""
        }
      },
      {
        "block_id": {
          "hash": "EED959EB2A13EA804C50BD9AC7F0BCB24C542ECA",
          "parts": {
            "total": "1",
            "hash": "E9C103F338E0FF26CE4EFC0226E2B3CDE3632CEB"
          }
        },
        "header": {
          "chain_id": "gaia-8001",
          "height": "870158",
          "time": "2018-10-26T03:51:40.958871877Z",
          "num_txs": "0",
          "last_block_id": {
            "hash": "F989DB86E3108272EDD040CA5C71848DD43E8A1A",
            "parts": {
              "total": "1",
              "hash": "BFC669160F8DAC08658B1E2BCD9CD3B529D17D5E"
            }
          },
          "total_txs": "99824",
          "last_commit_hash": "3EB8CF860F02F33B5D5C27DE3D263B3B4715CCB4",
          "data_hash": "",
          "validators_hash": "FB6D785A83CB40632A74CB05E08B6121D8276156",
          "consensus_hash": "D6B74BB35BDFFD8392340F2A379173548AE188FE",
          "app_hash": "55841D14A41F01C03CC8004723A87B44B94CE836",
          "last_results_hash": "",
          "evidence_hash": ""
        }
      },
      {
        "block_id": {
          "hash": "F989DB86E3108272EDD040CA5C71848DD43E8A1A",
          "parts": {
            "total": "1",
            "hash": "BFC669160F8DAC08658B1E2BCD9CD3B529D17D5E"
          }
        },
        "header": {
          "chain_id": "gaia-8001",
          "height": "870157",
          "time": "2018-10-26T03:51:31.624607827Z",
          "num_txs": "0",
          "last_block_id": {
            "hash": "C3B852F74D5A9BFB0BF7CAF4FF6C5257AE0E8D4B",
            "parts": {
              "total": "1",
              "hash": "07C4436B9E029991A8549F0EE483BBA808339F14"
            }
          },
          "total_txs": "99824",
          "last_commit_hash": "B5CD4330F9F63083CF707882B668E9B1427EDC8E",
          "data_hash": "",
          "validators_hash": "FB6D785A83CB40632A74CB05E08B6121D8276156",
          "consensus_hash": "D6B74BB35BDFFD8392340F2A379173548AE188FE",
          "app_hash": "C728FD23A88E9BB92D75D5EDB681B63C35E12541",
          "last_results_hash": "",
          "evidence_hash": ""
        }
      },
      {
        "block_id": {
          "hash": "C3B852F74D5A9BFB0BF7CAF4FF6C5257AE0E8D4B",
          "parts": {
            "total": "1",
            "hash": "07C4436B9E029991A8549F0EE483BBA808339F14"
          }
        },
        "header": {
          "chain_id": "gaia-8001",
          "height": "870156",
          "time": "2018-10-26T03:51:22.749495564Z",
          "num_txs": "0",
          "last_block_id": {
            "hash": "7F89A08A3EB30912D4EBA4023E5F554E097B26D4",
            "parts": {
              "total": "1",
              "hash": "7737221981BE90D388584C453BB7C4D4E39073BE"
            }
          },
          "total_txs": "99824",
          "last_commit_hash": "36779AF1E374E8DD8649F4AB6ECDCA68B2B48D4D",
          "data_hash": "",
          "validators_hash": "FB6D785A83CB40632A74CB05E08B6121D8276156",
          "consensus_hash": "D6B74BB35BDFFD8392340F2A379173548AE188FE",
          "app_hash": "CAB3923B5CBE3668863AAD34E3948B378F9C76AE",
          "last_results_hash": "",
          "evidence_hash": ""
        }
      },
      {
        "block_id": {
          "hash": "7F89A08A3EB30912D4EBA4023E5F554E097B26D4",
          "parts": {
            "total": "1",
            "hash": "7737221981BE90D388584C453BB7C4D4E39073BE"
          }
        },
        "header": {
          "chain_id": "gaia-8001",
          "height": "870155",
          "time": "2018-10-26T03:51:13.391921763Z",
          "num_txs": "0",
          "last_block_id": {
            "hash": "58DED4B41BC6B6688DB2FB34CF0490CB99AFC4D8",
            "parts": {
              "total": "1",
              "hash": "6FD6EE35A1826F627B38E5F812873A5D250E6146"
            }
          },
          "total_txs": "99824",
          "last_commit_hash": "22431C0CDD5DA6D7DB67B430E2FDBB6725A3C490",
          "data_hash": "",
          "validators_hash": "FB6D785A83CB40632A74CB05E08B6121D8276156",
          "consensus_hash": "D6B74BB35BDFFD8392340F2A379173548AE188FE",
          "app_hash": "FEBE3F4567B9746CC62EF6309964A4E60E922030",
          "last_results_hash": "",
          "evidence_hash": ""
        }
      },
      {
        "block_id": {
          "hash": "58DED4B41BC6B6688DB2FB34CF0490CB99AFC4D8",
          "parts": {
            "total": "1",
            "hash": "6FD6EE35A1826F627B38E5F812873A5D250E6146"
          }
        },
        "header": {
          "chain_id": "gaia-8001",
          "height": "870154",
          "time": "2018-10-26T03:51:11.015571431Z",
          "num_txs": "0",
          "last_block_id": {
            "hash": "B2AA6C28750E9776B30C76007AAC98AA4601AA9F",
            "parts": {
              "total": "1",
              "hash": "AC6B0494589894DF5F1A2DE7F2BDBC5CCB896D79"
            }
          },
          "total_txs": "99824",
          "last_commit_hash": "391431E6B5C8E6975AE3A190FDA22B376F63B282",
          "data_hash": "",
          "validators_hash": "FB6D785A83CB40632A74CB05E08B6121D8276156",
          "consensus_hash": "D6B74BB35BDFFD8392340F2A379173548AE188FE",
          "app_hash": "FBED888B81960CBA2E0A78E23700CA2586AD9078",
          "last_results_hash": "",
          "evidence_hash": ""
        }
      },
      {
        "block_id": {
          "hash": "B2AA6C28750E9776B30C76007AAC98AA4601AA9F",
          "parts": {
            "total": "1",
            "hash": "AC6B0494589894DF5F1A2DE7F2BDBC5CCB896D79"
          }
        },
        "header": {
          "chain_id": "gaia-8001",
          "height": "870153",
          "time": "2018-10-26T03:51:01.747534121Z",
          "num_txs": "0",
          "last_block_id": {
            "hash": "EABA264313303AE96E296E22BD0FA9D9A0BD569E",
            "parts": {
              "total": "1",
              "hash": "9FD1FEE654F26F8C4FC6C50CBBEEDA3AE8517E2A"
            }
          },
          "total_txs": "99824",
          "last_commit_hash": "F6CEA8F50D3F651F07E20C9ECC28C39200697D87",
          "data_hash": "",
          "validators_hash": "FB6D785A83CB40632A74CB05E08B6121D8276156",
          "consensus_hash": "D6B74BB35BDFFD8392340F2A379173548AE188FE",
          "app_hash": "7AC23513C5EA6935F800EE36818B11DD1B857378",
          "last_results_hash": "",
          "evidence_hash": ""
        }
      },
      {
        "block_id": {
          "hash": "EABA264313303AE96E296E22BD0FA9D9A0BD569E",
          "parts": {
            "total": "1",
            "hash": "9FD1FEE654F26F8C4FC6C50CBBEEDA3AE8517E2A"
          }
        },
        "header": {
          "chain_id": "gaia-8001",
          "height": "870152",
          "time": "2018-10-26T03:50:52.682126891Z",
          "num_txs": "0",
          "last_block_id": {
            "hash": "612E311BA35E9F635E69932A796CFAE88983EB0B",
            "parts": {
              "total": "1",
              "hash": "F3C33A8C51DA9369065E495D939D2FD970E29055"
            }
          },
          "total_txs": "99824",
          "last_commit_hash": "7FB8316760EBD83DF7A423D1F082406D4F73BBDC",
          "data_hash": "",
          "validators_hash": "FB6D785A83CB40632A74CB05E08B6121D8276156",
          "consensus_hash": "D6B74BB35BDFFD8392340F2A379173548AE188FE",
          "app_hash": "744B4B157A8A110BDFF215101403B6835923A292",
          "last_results_hash": "",
          "evidence_hash": ""
        }
      },
      {
        "block_id": {
          "hash": "612E311BA35E9F635E69932A796CFAE88983EB0B",
          "parts": {
            "total": "1",
            "hash": "F3C33A8C51DA9369065E495D939D2FD970E29055"
          }
        },
        "header": {
          "chain_id": "gaia-8001",
          "height": "870151",
          "time": "2018-10-26T03:50:43.129254562Z",
          "num_txs": "0",
          "last_block_id": {
            "hash": "6C69465BABE0EF1FABCB1315E0DDB0C381278204",
            "parts": {
              "total": "1",
              "hash": "6EC9322D18DF32F1BD8239FBF84AD40F798102D9"
            }
          },
          "total_txs": "99824",
          "last_commit_hash": "36E79FB8D24680C76865D4E4DA1E965F8F23553E",
          "data_hash": "",
          "validators_hash": "FB6D785A83CB40632A74CB05E08B6121D8276156",
          "consensus_hash": "D6B74BB35BDFFD8392340F2A379173548AE188FE",
          "app_hash": "41F53533CF25141A82857BE7FC0A6DCD27A86E72",
          "last_results_hash": "",
          "evidence_hash": ""
        }
      },
      {
        "block_id": {
          "hash": "6C69465BABE0EF1FABCB1315E0DDB0C381278204",
          "parts": {
            "total": "1",
            "hash": "6EC9322D18DF32F1BD8239FBF84AD40F798102D9"
          }
        },
        "header": {
          "chain_id": "gaia-8001",
          "height": "870150",
          "time": "2018-10-26T03:50:40.621250113Z",
          "num_txs": "0",
          "last_block_id": {
            "hash": "773E9D415F587DAFC24C4705703944E36C865354",
            "parts": {
              "total": "1",
              "hash": "6483AB1961B834F01A1B11A35F00329CDDFD25A0"
            }
          },
          "total_txs": "99824",
          "last_commit_hash": "7D0B8407FEB6C8EF7F35A58F4DA9E4BCA93418DE",
          "data_hash": "",
          "validators_hash": "FB6D785A83CB40632A74CB05E08B6121D8276156",
          "consensus_hash": "D6B74BB35BDFFD8392340F2A379173548AE188FE",
          "app_hash": "8550BF4059DA3CD1A055C6A9388E18B16C892E74",
          "last_results_hash": "",
          "evidence_hash": ""
        }
      },
      {
        "block_id": {
          "hash": "773E9D415F587DAFC24C4705703944E36C865354",
          "parts": {
            "total": "1",
            "hash": "6483AB1961B834F01A1B11A35F00329CDDFD25A0"
          }
        },
        "header": {
          "chain_id": "gaia-8001",
          "height": "870149",
          "time": "2018-10-26T03:50:31.708795087Z",
          "num_txs": "0",
          "last_block_id": {
            "hash": "537EA7C20357CB62FF6018DD73E2AC58BF9BB5E2",
            "parts": {
              "total": "1",
              "hash": "05B34645707795D6BD90C347AC7C44BE151E14D8"
            }
          },
          "total_txs": "99824",
          "last_commit_hash": "4651957FBDE3A6E0886F112FBA4E2354AD5DBD9A",
          "data_hash": "",
          "validators_hash": "FB6D785A83CB40632A74CB05E08B6121D8276156",
          "consensus_hash": "D6B74BB35BDFFD8392340F2A379173548AE188FE",
          "app_hash": "47133F731A6B36454826B7D2AB5242A1BA32E35D",
          "last_results_hash": "",
          "evidence_hash": ""
        }
      }
    ]
  }
}`,
	"broadcast_tx_async.json": `
{
  "jsonrpc": "2.0",
  "id": "",
  "result": {
    "code": 0,
    "data": "",
    "log": "",
    "hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4"
  }
}`,
	"broadcast_tx_commit.json": `
{
  "jsonrpc": "2.0",
  "id": "",
  "result": {
    "check_tx": {
      "code": 65538,
      "log": "=== ABCI Log ===\nCodespace: 1\nCode:      2\nABCICode:  65538\nError:     --= Error =--\nData: common.FmtError{format:\"txBytes are empty\", args:[]interface {}(nil)}\nMsg Traces:\n--= /Error =--\n\n=== /ABCI Log ===\n"
    },
    "deliver_tx": {},
    "hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4",
    "height": "0"
  }
}`,
	"broadcast_tx_sync.json": `
{
  "jsonrpc": "2.0",
  "id": "",
  "result": {
    "code": 65538,
    "data": "",
    "log": "=== ABCI Log ===\nCodespace: 1\nCode:      2\nABCICode:  65538\nError:     --= Error =--\nData: common.FmtError{format:\"txBytes are empty\", args:[]interface {}(nil)}\nMsg Traces:\n--= /Error =--\n\n=== /ABCI Log ===\n",
    "hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4"
  }
}`,
	"commit.json": `
{
  "jsonrpc": "2.0",
  "id": "",
  "result": {
    "SignedHeader": {
      "header": {
        "chain_id": "gaia-8001",
        "height": "870169",
        "time": "2018-10-26T03:52:55.617565909Z",
        "num_txs": "0",
        "last_block_id": {
          "hash": "27F56F86AB02B26B77011FBFBFDBD185555D79CE",
          "parts": {
            "total": "1",
            "hash": "208F863FB187966AA38A154AFAEBB42B18886B2E"
          }
        },
        "total_txs": "99824",
        "last_commit_hash": "ECD2FD99B69C6576154FB00A75FA7884647275F5",
        "data_hash": "",
        "validators_hash": "FB6D785A83CB40632A74CB05E08B6121D8276156",
        "consensus_hash": "D6B74BB35BDFFD8392340F2A379173548AE188FE",
        "app_hash": "8D384E78F4A4A3BE341F62CF1F326342ED3DE369",
        "last_results_hash": "",
        "evidence_hash": ""
      },
      "commit": {
        "block_id": {
          "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
          "parts": {
            "total": "1",
            "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
          }
        },
        "precommits": [
          null,
          null,
          null,
          null,
          null,
          null,
          {
            "validator_address": "0DDF97111D73DB825138EC15EC27DE5FE8536901",
            "validator_index": "6",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.778059781Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "+NWo0ZKLO9yJF9wnpbH8kJj+g6na/oOxSVHi6pP+5zJsD7qjYiwWhxlTA+z/WVq8QDWePU6K1rp044zX00lLBw=="
          },
          null,
          null,
          {
            "validator_address": "296869AB66F7003A5AADA1A8DEEEA2913486668F",
            "validator_index": "9",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.787310675Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "EG/zCbCrUzQI1nnQGlH3txCjVwhvU4QvKqrsaNRZdt6E05sDDsAByqqDfr+pc8BXGbizi2ZHxIM1/y5qZiA7AA=="
          },
          {
            "validator_address": "29B93E4EECD8C591AB76C2D52FD7C43CBEEEC50B",
            "validator_index": "10",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.817374631Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "MQMHvL76sCUpPkAKGoAWGqbIk2nRzq+4F8RM6N4LdPUQnDFoD1Sj80/QM7O8BlgY/BYuNxWK9E/WhCzYels3Cw=="
          },
          null,
          {
            "validator_address": "33412C9A69C60CB0FAA24A5C5B3A906DE24B47C2",
            "validator_index": "12",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.837627809Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "ZQdKqA4P6y5i7Ub01gm1rj1a+emNxfgMvXfiPde4kOWT/8ggKcuoGzHpqqAHsj7utvfB2maHeklXQgiKyn+PCA=="
          },
          {
            "validator_address": "3363E8F97B02ECC00289E72173D827543047ACDA",
            "validator_index": "13",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.85044504Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "Ag5L9BFwZO/u4A0+X2SrVIdL0CB6Q3I4JHNmvPMAHPRNVTsSMZNmuMPJuTMvvsF22uw1+B8A13rVTe23p3nyCw=="
          },
          null,
          null,
          null,
          {
            "validator_address": "3E899BA36A93C7370347BD5FE85909F2ECD93F8C",
            "validator_index": "17",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.811725807Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "j7szltvkxAtcexheEZ9TztsS02PK1kaUA8seEDORE3qYp+jIU9k1JQZBLZIQuahcrqiTaTZ4NrsRM3JDY2poCw=="
          },
          null,
          {
            "validator_address": "4944F8BBBB5303E17BC0607DF332320EDBC750EA",
            "validator_index": "19",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.75786553Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "k+klrjdKguh6RCUy9ejA72b3+yfSeWDtxgTJmP2PuDRdJCDJC0MYBGJi8Qp/dNZ6wkkgYzmjwcDutV+TxcuwDA=="
          },
          null,
          null,
          {
            "validator_address": "5104F488EA6B1E2C33CBDA3C18389C88104DBC72",
            "validator_index": "22",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.780299315Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "AIWXFxqtAQeRLaXWqwBHyGe/3bTpQN9ZxaiNCpjc70VaPTcF4Ch8kw9LODCh/2hFKV6cpXZp6ezgEjKOFpr1CQ=="
          },
          null,
          null,
          {
            "validator_address": "5BDF8EA41B1835FBAD865B181A2A4B23F511B5E4",
            "validator_index": "25",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.739335555Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "pZEsv7ozMxw7t/pH/o8dXYmfMwzVxGhg0Y+rnO3qrMzDJSNNijLDSWHXNeSQhNwROy86RrqAsm+VoEee6ncMCA=="
          },
          {
            "validator_address": "5DE9674DA2D535D5B1547D90990E9087EAE15E39",
            "validator_index": "26",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.80385725Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "7WsoPWF8txIIYcVLMkFR3MuMpupgA0UHNv9OlUKhyDgJBqF5HynYgPsQTq5zD4jVP3TrhOXddWnUvoYAT3IGDA=="
          },
          {
            "validator_address": "5E06FB65FD0E0EBB11D0D74A1E480B4A8D7D0FE7",
            "validator_index": "27",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.767206066Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "ALmaRJJ8Ss7wUPSpGmHW27UOxE9LeqTjdDsH8Q5eRHYFtyuA7Nv+NfbxJaYSmTsAxgMcL91j3ryA/znH7/1OCw=="
          },
          null,
          {
            "validator_address": "651995D8471D07A22C149265F5237A89149795C1",
            "validator_index": "29",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.770414899Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "aCauDodBAqr+8fPtCDt2KHUurMTpLPyCGW8wnYVD0Q5Q8367jzwRVvPDMWsJxneOVu18acKhRI1lAlWIOQjdDQ=="
          },
          null,
          null,
          null,
          null,
          null,
          null,
          {
            "validator_address": "71BC71D28E4599BCA22214BBB2CE0DFEE90A9EE4",
            "validator_index": "36",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.791865019Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "tpWMHwR1Yo50utznwm1LVL0Oi/v9gCnDqst5HpMl265w43QmW56qhNDKdYx59onv2LkNl6FHscku7sZ9aBHeBA=="
          },
          null,
          null,
          {
            "validator_address": "7A32D912F7A790EEE5610653DC4DDBC0DB4F2497",
            "validator_index": "39",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.906796916Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "cGmhuRI8Z7KdtJL12OytZ3x7d88SGJb8jMbXcX48CawfThCSQXbZYlCkvS5HN6x0ZJ4c7i9+7AFB4r2ANR1vBg=="
          },
          null,
          null,
          null,
          null,
          null,
          null,
          {
            "validator_address": "86DDBEEF35D35786A68834B34869174AC314BB74",
            "validator_index": "46",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.793936015Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "XyISkDhYTarCnUOok/V08wvZ8ggsapau+nCp5BT/RObISnzFxWytyGuMa6Dnfl6LtxgSyN0CatieJQB+VipbBQ=="
          },
          {
            "validator_address": "89B357643A945BD433CB5B2B9CBD9B90CB9CCD5E",
            "validator_index": "47",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.833512472Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "dtVn/19XODNayoIeBXLYqRqhRF2xaVPH00pbJLc05j6C54Ou4ItV9JVpfT76j+HcycysRKfCpcgqAmnV+HPnDQ=="
          },
          {
            "validator_address": "8DF81F7913F1EBF856618B1475C12EEEC1214BB0",
            "validator_index": "48",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.739131862Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "73BDv2ZtTa7+8bRNlO8r7v7f/RtZYOvgI84HAw5g/oZAdl5D2BOdtYiThECR3pI1nsuj2/c8YVDScdppczgfAA=="
          },
          null,
          {
            "validator_address": "8F06FE1FD042683C426137FBC4BBAA4288924217",
            "validator_index": "50",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.79696179Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "nJlKTykTpEu4EvkWEGLpJm6DmVbEWLOu/e8DB4Px7oEq3kEipI6ySbtsq8tfTJS3lUPwT9C8GZtCHxigZF5yDQ=="
          },
          {
            "validator_address": "8F302C08318444DDC43EABDE96E869D0B1F5A421",
            "validator_index": "51",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.811881801Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "LAmbJJhfrBAMpS4Fq8eVGDwmnGCUSMJ+jnZH/hB3Y4hrIES1eHQ+2F8OY6GzntyvC2+kbytb/mdD8AcbeG5jAg=="
          },
          {
            "validator_address": "927CC759310F4DF5C8936EF6D666ED3135D609B7",
            "validator_index": "52",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:44:56.612974293Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "6jreklK/UDDzPvB+IV/CAYlMPjBP4280hM4y6NWTYDqD/YuGgdoBZbtR6/F/hCm3yNH3RBmWq0Jk/BOoP6U3Ag=="
          },
          {
            "validator_address": "9524C2BBF5C9BCC7AFEE121AF03E31D8480DD5E5",
            "validator_index": "53",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.763154858Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "2I06TqKvII8YWB3OKMQOKQ4wZb1TnV1h09ctameKp/6sYOe7IlBFJsdNasqDsGUuB64aTFsVl263SiCzGF1bDw=="
          },
          {
            "validator_address": "96BDE45EC52606CF5CB59853097871AD60680549",
            "validator_index": "54",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.806633715Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "oGQA9NfD/OrZDhSVtcUMa2qCPGomr0bRuARQ0VmCTfZA5wb8rp5zghYXj6YQuLDhOPclvlGYojWvrN7G9pUHCg=="
          },
          null,
          {
            "validator_address": "9D93AD02FE1BD7E76B9E58D7EB461088EAD011E5",
            "validator_index": "56",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.838255493Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "tsSLyAKMDHTr/+WNbJsdMXLbLioDP0FL27nJtFnNvshgATUqePT50shs3g6vAlI7Gd05dWItfDG0i9T3hPY0Cw=="
          },
          {
            "validator_address": "A0ED8A257D7862094FE8015BCE2135D111722DF6",
            "validator_index": "57",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.902738007Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "gffnREFpA50V/LcQsVBJpqv8cUMKRgMZpVgIY/JlYSLtvKVi29ptXbdi5gJ/l9AiHYzXxQSoRkOSdkZ7dtOgAw=="
          },
          {
            "validator_address": "A5868971300F155020634D7BB49C7634F7AACCD3",
            "validator_index": "58",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.801667953Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "HaO+U6WiAyD9ahOy3AdnwIXSax3NDJARou3jzCdyP1Nrb71WrJe3uoqzpgnh9LcdYNsC97TYS5s6NuHb9qs4Dg=="
          },
          null,
          {
            "validator_address": "AFCFB9B8DF23559CB5EB97C7CF1D77B411B2F47F",
            "validator_index": "60",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.865093074Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "ZbIGOwC8fu2xK55gFARbvtDsxaHAc097xCTC+cBcTx8HHpgYpEkSwi9pHfZvZkifP9pdwjp/aLgCYV6xznCwDA=="
          },
          {
            "validator_address": "B0155252D73B7EEB74D2A8CC814397E66970A839",
            "validator_index": "61",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.780453186Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "RR2MQH4Ec3/6JJCKmoI+cOUL7a63LoEUWbfTcNlBz4oW9BRL3Zhxt4sttayh2nQsJrQcYy++XW8bHJ3RuZPIAA=="
          },
          {
            "validator_address": "C39A8A2B85198DCB8F23AD1F6C198E2BA7AB5D60",
            "validator_index": "62",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.807344039Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "QDff6UsHM8cdXXqdZoi0m+q84gJCmudM3JoicjlMtfw7E2lEXWLBnGUfGSCn/u2W4c++HOZ/v+TIiIgQ5275CQ=="
          },
          null,
          {
            "validator_address": "D31B1198AE055BC52C600C19751011D95A64298A",
            "validator_index": "64",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.67663613Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "/jdcvqeGUWqe/fiXuHdsqALcudbhUf7eqUcr6rZ4MVRQSHkvDC0y8wA/OwUf5ydls5B2gp3b7mi708Dq4yCKAw=="
          },
          null,
          {
            "validator_address": "D77FB84A729972CD3E757C9BEB61561855832B92",
            "validator_index": "66",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.819113084Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "7CeFJGDI4LqHttWgYJGOzm/FsJM11czizgWNeWJG/YLbU70oxU/eujsy0gaUSggLUI5gxjmJBDy6mfGU7zh4Cg=="
          },
          {
            "validator_address": "D981CB6DDF336E311A33C585C23C10F9F033A305",
            "validator_index": "67",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.759749215Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "lcbHKLGnwn1y2O5iy6XgxR1+P+KiogdSGm/wvuNNt38xx/kqEA2bmtbmCn5HQmyksjlCBGk+HW+WSv27x+sJAg=="
          },
          {
            "validator_address": "DA38013D7BCC3EC46109F8CDDF14F8D080CC4A9D",
            "validator_index": "68",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.811658981Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "IfxKljlCwGC4lZjLs4VRUBcaVhdFlumQI/9Ofwq5dBNecnQsAt+dm7IXnBwigkt5Hhi1ZrSA23fZ5k/A7DdMCQ=="
          },
          null,
          null,
          {
            "validator_address": "DDDC63E5B9C7D1509805601E2EF6934BAEAC9115",
            "validator_index": "71",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.787206302Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "uwT/1fXES5Asz6ZyANVsUuU21Jc3RyrpOFEiitoecHFYTa469i9GqquQsnrBSKVyNy4vgVqpdJ8qmadxqDiBAw=="
          },
          {
            "validator_address": "E39B5778E0D4297A46F45CF6940B0A163A430FF8",
            "validator_index": "72",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.798095801Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "2FcSpgD8kdluw4FWTNczKyRu1MJR7MBAtiiNGcMJYxzy0YeSR9NVBvE3kqeYwraOjBPk0rteou2qvNU0g/FzDg=="
          },
          {
            "validator_address": "E3B19738F7132F950668504D326C7660F159D170",
            "validator_index": "73",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.751814783Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "cEztXk1uwSjhCfpXIddD3Y1BJkp8/tc6TL/Iiyfd7XZhRfwupSp8NJIUP8H1hIkFWWWQasPGNrTzymy0jMC+Dw=="
          },
          {
            "validator_address": "E42F8AEABB685219F7027F54D2B0DBC7846F633A",
            "validator_index": "74",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.810388577Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "HGvwa8rU8U4/NdiXUt6F3ay3+vWZHoJT0kxUVv0PnOVE3ZIVHbD40B7hyCJIlxDB3DMD/ZbxB8XmSO/Ue+4XCQ=="
          },
          {
            "validator_address": "E4E6C2E81269DFFA196C17F98240E491479A6008",
            "validator_index": "75",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.791549052Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "RohUveYqBu+ON243hKa8F6U0cDADPC0KQEkESSk9gVZ+Lliw5XNNRuWjDfONCZoTIWxCRyIhOCio4jCTP84cDQ=="
          },
          {
            "validator_address": "E50ADECD5FD27FA2CD610C07F8AED36A2FC4A6A6",
            "validator_index": "76",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.76106743Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "spd6q4ZjIizXqmYE2vQExf9px4ps1y4K94J50sndQsO+OV/qamMlyVVzDuvDc/aAS1U+PcrXkbNKDe1U+so9Ag=="
          },
          {
            "validator_address": "E56E00A6B6AC1BE546CE9A9F34C7C1BB7585A8C1",
            "validator_index": "77",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.800035638Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "SBokRpEJzYw6ZXaEo3Rb4P7N8yI1haF9Mq7p7+pQm0KLpsZ/RTlxp51K19/Idp5b8oipgc9SzbRnf7vcPaY+DA=="
          },
          {
            "validator_address": "EE67A5CA3269434AFAA41CE43E17B467488B5210",
            "validator_index": "78",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.778359385Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "AxfXhEWTsFphV9gOd0fTmWbicAHGbOMyYxVuDuD4WAO7tr4RkIUqo2Q6HG77+HpLJ1oYOEaY1fHLC4JcyCUTCw=="
          },
          {
            "validator_address": "F23FF36BD5B90C33CE3A03ED72DBDCF5EC07D6AF",
            "validator_index": "79",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.837983566Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "WT+k86V323hffNKzRUzDjdJDDPxtNktzjO5zscFLFSQnwG2X67nJbiWn5we+8J68GJu7wIstlJCnYm/Srt2hBQ=="
          },
          null,
          {
            "validator_address": "F6738260186D33D9C14FC6E7017AFE6BB952A63D",
            "validator_index": "81",
            "height": "870169",
            "round": "0",
            "timestamp": "2018-10-26T03:52:56.709460782Z",
            "type": 2,
            "block_id": {
              "hash": "A5FCD9A86F5B23DBA0C32D11B6B5EB12A2182A75",
              "parts": {
                "total": "1",
                "hash": "9400F18B7BA6BD940B84F97CF21DFAA3CF7B6B99"
              }
            },
            "signature": "c/GwNoF+uGFV1/7mWOKLKPk9iQwcssQ7Ld56HC01I0TEX/34Pq9xM5Iunmf1uaGZ6gIPvMzW0jWlo5+AfrjXCQ=="
          }
        ]
      }
    },
    "canonical": false
  }
}`,
	"consensus_state.json": `
{
  "jsonrpc": "2.0",
  "id": "",
  "result": {
    "round_state": {
      "height/round/step": "870167/0/1",
      "start_time": "2018-10-26T03:52:47.670119461Z",
      "proposal_block_hash": "",
      "locked_block_hash": "",
      "valid_block_hash": "",
      "height_vote_set": [
        {
          "round": "0",
          "prevotes": [
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote"
          ],
          "prevotes_bit_array": "BA{82:__________________________________________________________________________________} 0/46753 = 0.00",
          "precommits": [
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote"
          ],
          "precommits_bit_array": "BA{82:__________________________________________________________________________________} 0/46753 = 0.00"
        }
      ]
    }
  }
}`,
	"dump_consensus_state.json": `
{
  "jsonrpc": "2.0",
  "id": "",
  "result": {
    "round_state": {
      "height": "870167",
      "round": "0",
      "step": 1,
      "start_time": "2018-10-26T03:52:47.670119461Z",
      "commit_time": "2018-10-26T03:52:42.670119461Z",
      "validators": {
        "validators": [
          {
            "address": "0077F68359FE12C678980AC3442460CB1CC1ABC4",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "zf+aS3anGNm9CrpliDTqpUlHCBpi0qZknUo+aYOXC+s="
            },
            "voting_power": "169",
            "accum": "4702"
          },
          {
            "address": "0081A939ABBB733E8DC36E406EC51E4309B9AE3A",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "Sk7VmvaRgatGyVrbkbQ4QBCeda/DIUs6E3NzstWQV6g="
            },
            "voting_power": "154",
            "accum": "8494"
          },
          {
            "address": "01F78669F9515FD83DF9250F5C0EE143D3DAD65C",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "1K9/Z+oiK380Obf+LYVN6P7ydFIWCFvNQpT+ToRUURg="
            },
            "voting_power": "627",
            "accum": "-11067"
          },
          {
            "address": "0548261C50222FD710EB5EBDF03A0E6567B21D20",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "8K3clCjVU33BTIpUhdahGmu++WxHj4NUE9krCRkk++s="
            },
            "voting_power": "1181",
            "accum": "-6426"
          },
          {
            "address": "066C70AF46E5E1B473BD125FD8937EF90E555641",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "To+MjLVGM4gNuE7JahfDwIc5Q+31H+Q60Cy4AoU2fjE="
            },
            "voting_power": "5",
            "accum": "14623"
          },
          {
            "address": "0A692DBFFE9E5DD28E7DECC33CED1AEB4C0D014E",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "N3K5kDdfcKJurfaa6s2zfKgtYvz1Pagz7VWi9ZfX8yM="
            },
            "voting_power": "439",
            "accum": "-4020"
          },
          {
            "address": "0DDF97111D73DB825138EC15EC27DE5FE8536901",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "H0SIA/BU6Xj8oT5bQkvLpEITN3CqFLbMeBcQ72NZrAE="
            },
            "voting_power": "18",
            "accum": "15036"
          },
          {
            "address": "1ECD8C5FE1341D74CDC633A245DD56858DE2F9A2",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "lUMCRAeu47BsOhNvCQTQJQeB68z0/VaElC9j5gDt9y8="
            },
            "voting_power": "578",
            "accum": "6350"
          },
          {
            "address": "20FBB777AB2676E19A31E2EB35B2C7B426C99C13",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "/DtRehBSochyxchfAn/Tb4tbwC9u6Ryq++VeBE72FuA="
            },
            "voting_power": "39",
            "accum": "14178"
          },
          {
            "address": "296869AB66F7003A5AADA1A8DEEEA2913486668F",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "4cI0avx/jOL9eWe7PyRm4Mhv1QSSrhxpyr30f7tJ7i0="
            },
            "voting_power": "10",
            "accum": "-24623"
          },
          {
            "address": "29B93E4EECD8C591AB76C2D52FD7C43CBEEEC50B",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "GTwvM3OOWuXdwH7suTJIWdYdtwUU1hLAHIXmpdzGfl0="
            },
            "voting_power": "7",
            "accum": "10603"
          },
          {
            "address": "2E6AF0D5B1A85E173F5EBEDA93D7E7D97A88D06C",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "G3YJxcM2oMCW8R5s/HSqmnudsQjcj+e/vpZBtMyob1w="
            },
            "voting_power": "121",
            "accum": "-5768"
          },
          {
            "address": "33412C9A69C60CB0FAA24A5C5B3A906DE24B47C2",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "4DEDoU/RsHMPS54GzgwkWnW5zPQfMt9aInFFc3GyfA8="
            },
            "voting_power": "40",
            "accum": "-24233"
          },
          {
            "address": "3363E8F97B02ECC00289E72173D827543047ACDA",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "mPnu910hOOa1tAQ7pbOLFDxvllbQUmrbtGjqQrYg1nM="
            },
            "voting_power": "89",
            "accum": "5451"
          },
          {
            "address": "36E9B7FAEC964F97D84C028ED62493E373AD0B38",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "ViJzrw9Unk1XyDg4f0UQLFE4zfUaFh23wZZY0f31Qg0="
            },
            "voting_power": "139",
            "accum": "10605"
          },
          {
            "address": "398610C4CF11C84C89AC6975470972EE75DA17E4",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "5GAxIHzu06l0n3MU8wNQrCcrrnpZR9EKe5qZMsM2h/E="
            },
            "voting_power": "490",
            "accum": "-6646"
          },
          {
            "address": "3AC99B708404452366E0953A4BE849226DC86CBC",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "ISIM341M+EUYOMpwhn9gyCvRUZ06xomp/+lOwa50meQ="
            },
            "voting_power": "480",
            "accum": "-8795"
          },
          {
            "address": "3E899BA36A93C7370347BD5FE85909F2ECD93F8C",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "M2Iu/A8RQcmoiIEar/iV4wBjVj0I0gcd/nNnvwmnT2g="
            },
            "voting_power": "39",
            "accum": "1466"
          },
          {
            "address": "411C430038318FE4DB32AFC508A6C12782161294",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "LmJN98a+ce+shthBz95qKsdwUwWRs1QdBLO4HGNQDpM="
            },
            "voting_power": "107",
            "accum": "-24961"
          },
          {
            "address": "4944F8BBBB5303E17BC0607DF332320EDBC750EA",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "A6GzeXUM3vsXaDAEYMSDgSKkqn9AoUYjs8empH46MGY="
            },
            "voting_power": "1288",
            "accum": "9807"
          },
          {
            "address": "4B9C6FF8B36EE6134AA8D822772C0162C471186D",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "g0f/B/FxRDU0VbS7vMXs8cJwSUITpbtn3NjEQ0NJ5Fg="
            },
            "voting_power": "20",
            "accum": "-8548"
          },
          {
            "address": "50C18305F49BE332E4C6BCDE0DEECD4A22B1D6A4",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "PX8v6nM+dk8fnqQWuXaRpRrdlhL+sMUHYo75CJbnY9U="
            },
            "voting_power": "109",
            "accum": "17967"
          },
          {
            "address": "5104F488EA6B1E2C33CBDA3C18389C88104DBC72",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "JREG9eji9gfJR2b61qMu8jNZ835ezPY2A4o+zjYMfDg="
            },
            "voting_power": "299",
            "accum": "7304"
          },
          {
            "address": "570B6A755ED339EAC56BB3390A48069F6AC9449B",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "5XCMLGNTGkMh9/JsUzi/8Xl7Y0cdDe38stX+zvBLXuE="
            },
            "voting_power": "44",
            "accum": "16078"
          },
          {
            "address": "5B9628695CDFC6025857578114E3D3126687EAE1",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "NCwjL8K9R2CLOMYub3MhAkkB0fktT0mZC75TEtcXyQ4="
            },
            "voting_power": "130",
            "accum": "16074"
          },
          {
            "address": "5BDF8EA41B1835FBAD865B181A2A4B23F511B5E4",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "RBuRRNEzA9RA1Wrdi9PPFQJ29/n/bqN9O2tQv9Gq248="
            },
            "voting_power": "6040",
            "accum": "-21204"
          },
          {
            "address": "5DE9674DA2D535D5B1547D90990E9087EAE15E39",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "yRk1EO5Ta0iCI68BjluGEyZIlkBTGvUzbl+s/Rujhco="
            },
            "voting_power": "5",
            "accum": "14623"
          },
          {
            "address": "5E06FB65FD0E0EBB11D0D74A1E480B4A8D7D0FE7",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "iPPJziJLaL5QA6Mr6uaUGwoNrGts45W6LVIKbo11GfA="
            },
            "voting_power": "41",
            "accum": "-4298"
          },
          {
            "address": "5E8673673E37450F01B64138FBF4B172BDB52D68",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "CuPin/hLM2tK6KvcUlOi/Brm5xNi3zq6UPgpLFPSzt8="
            },
            "voting_power": "776",
            "accum": "-18801"
          },
          {
            "address": "651995D8471D07A22C149265F5237A89149795C1",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "WX9FldKAX82ALt3u3v80v/1zLUNJQO689C9vwXooDuQ="
            },
            "voting_power": "142",
            "accum": "18211"
          },
          {
            "address": "668309BE1A6F5317197B0BC6E47153D249DD5DE9",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "DKXu53utgN9Q/4dd7uRC0c8j1bNHTGSVDzhlEcrRFl8="
            },
            "voting_power": "103",
            "accum": "-18770"
          },
          {
            "address": "69D99B2C66043ACBEAA8447525C356AFC6408E0C",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "8pfpbIxBBiu88hpxS3CeRpv7kClEjl8SwVgckDNBGlE="
            },
            "voting_power": "541",
            "accum": "-410"
          },
          {
            "address": "69EDADFA4A803E9DBF548C900D4A465E7D154262",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "LRK6flgMsHrBYkIIahegEkzlwC7gQ3dAyw0jcAd1u/k="
            },
            "voting_power": "107",
            "accum": "-25012"
          },
          {
            "address": "6B275FFA0571EEFC4482CD7A6CAE836F5785CB4E",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "pggq0WUTlsiA8cBiptRgxlw4WUG2GmXfLQxkbMbCBPk="
            },
            "voting_power": "202",
            "accum": "16157"
          },
          {
            "address": "6D3AD380E398D8592B03AA124BAFED5F125DF863",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "ENAVynNXVpj/IdYx9kCPKaPs4bWSxRIHNlmS9QiDuZQ="
            },
            "voting_power": "54",
            "accum": "-2889"
          },
          {
            "address": "6F6C93B8BF3495063FEC9B5C771158607F4AE9D6",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "S8s6fdAQNQ3bN9SNVAsHB/j8uv1CM1roxeLesL+fh4g="
            },
            "voting_power": "141",
            "accum": "7071"
          },
          {
            "address": "71BC71D28E4599BCA22214BBB2CE0DFEE90A9EE4",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "MWID5oHOb367w/js26+cbQdyGqw2KEpe967D44z/eSQ="
            },
            "voting_power": "1",
            "accum": "-24087"
          },
          {
            "address": "732CEEF54C374DDC6ADECBFD707AEFD07FEDC143",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "HjSC7VkhKih6xMhudlqfaFE8ZZnP8RKJPv4iqR7RhcE="
            },
            "voting_power": "357",
            "accum": "-7150"
          },
          {
            "address": "7A20421F99FC4362B13C2FD43548F3E75B057017",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "HWtL6kXsdXL6Lxoq4XXGG3FkpPAmDBFwrP1ndYCvsnA="
            },
            "voting_power": "20",
            "accum": "11548"
          },
          {
            "address": "7A32D912F7A790EEE5610653DC4DDBC0DB4F2497",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "d5o/ErYmwkS/Cw06SH1l3vrBDXRUJLtjj8kcz9CEiFQ="
            },
            "voting_power": "79",
            "accum": "1483"
          },
          {
            "address": "7EB7593F519F1180C7E872D0494611464EDF03AE",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "VakMQSPBEuSC9Nwuv8WWhrZVUmH31bUR4+G6pJhkgE8="
            },
            "voting_power": "1807",
            "accum": "3131"
          },
          {
            "address": "807CF675BA8479FC4FEC1047215DF07970040F1F",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "2p8s/pRBZPjYWKKMlR7AOXypDzDmPo762iXlKpCwtco="
            },
            "voting_power": "14",
            "accum": "-24578"
          },
          {
            "address": "8340F82D32C4A221C57FC33FCBA44309C4DF1BBD",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "FSj9jMpCOdWfX0rCQ0TpnEa8/uOroQl4V4Bk3Msck3w="
            },
            "voting_power": "65",
            "accum": "3328"
          },
          {
            "address": "83AB321E012C57746689FED3A6064DF331F81007",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "aUViYC2znC55sleHfmsIN9hZ45SbYPbDcYA0gVzglsc="
            },
            "voting_power": "946",
            "accum": "574"
          },
          {
            "address": "850D387A41ADB21F7D67EAC5076D594DF3D861C7",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "Sowu0HYQbGFJ4dA3urnWwXqMc1+6kzQHre+9/iVlpjg="
            },
            "voting_power": "10",
            "accum": "386"
          },
          {
            "address": "85335B5077FAE8E2926AD338634079124EE1C4A1",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "lGnYgu+rv9HZj3Y+fS6Ad/GszZgtwwyZcM9mAc/tuu4="
            },
            "voting_power": "475",
            "accum": "-23502"
          },
          {
            "address": "86DDBEEF35D35786A68834B34869174AC314BB74",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "ilDFePqWkNRNWUEoWKq42jZC5CD4rglfs7+gCuq6Qc8="
            },
            "voting_power": "1867",
            "accum": "-10419"
          },
          {
            "address": "89B357643A945BD433CB5B2B9CBD9B90CB9CCD5E",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "UVRp+iVQUv9erXJwqfI6ClWIXk/8fSs0CRudMMU49Zo="
            },
            "voting_power": "36",
            "accum": "-17396"
          },
          {
            "address": "8DF81F7913F1EBF856618B1475C12EEEC1214BB0",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "0WFwqm+9/VGNlZDcAlkqCdQm/8M+W19LWnzXIDg/py0="
            },
            "voting_power": "221",
            "accum": "6279"
          },
          {
            "address": "8EA62D00960B7557560E0D78278B6B36165AD0A8",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "43IDu79dbTK5jDg1x2xR09pzWgWdAGHoIrz4Hk9H//Q="
            },
            "voting_power": "10",
            "accum": "-1684"
          },
          {
            "address": "8F06FE1FD042683C426137FBC4BBAA4288924217",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "lB+CgSbTzpHXJHUU15fJIPaPl6XF0nSfdNxBnlKJqTk="
            },
            "voting_power": "2466",
            "accum": "2855"
          },
          {
            "address": "8F302C08318444DDC43EABDE96E869D0B1F5A421",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "+yk80I86yZ+wqwf/9iDEfewwZGDaXCKosDuWZnlSf4M="
            },
            "voting_power": "110",
            "accum": "-1766"
          },
          {
            "address": "927CC759310F4DF5C8936EF6D666ED3135D609B7",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "tkdOxR2CIOctyc5e9IMGwnMwB420OTN3Rfb6fUMVhr8="
            },
            "voting_power": "11",
            "accum": "3334"
          },
          {
            "address": "9524C2BBF5C9BCC7AFEE121AF03E31D8480DD5E5",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "nWft/3AE/5/Iv5ePNUbkFHeG8HpNnnIFA3hQKR6qRmQ="
            },
            "voting_power": "5",
            "accum": "14623"
          },
          {
            "address": "96BDE45EC52606CF5CB59853097871AD60680549",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "4JoJuRfaANhdM1x3AWRo1/Cj9DH3VA+fi1SynzknV+w="
            },
            "voting_power": "19",
            "accum": "-10061"
          },
          {
            "address": "989E892A603E6BA7F325BC52555DFB47A8751E66",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "zkQOeG4/QylfQCqxV8WPzC0bUvIkRb467gbL80m4/KA="
            },
            "voting_power": "10",
            "accum": "-17159"
          },
          {
            "address": "9D93AD02FE1BD7E76B9E58D7EB461088EAD011E5",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "jj0Y/Fy8JSJR3g+PHU6Ce0ecYwHGUVJ4bVyR7WwcyLI="
            },
            "voting_power": "614",
            "accum": "-12946"
          },
          {
            "address": "A0ED8A257D7862094FE8015BCE2135D111722DF6",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "dvLH3HPQ3YH0oDefmGTPvpwn/gTFhvBAQiYNV4DVs5k="
            },
            "voting_power": "195",
            "accum": "5698"
          },
          {
            "address": "A5868971300F155020634D7BB49C7634F7AACCD3",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "cCoFsZzKZ9SQZbHe4NueVObIezP6ts0tRTZ/aN96dig="
            },
            "voting_power": "18",
            "accum": "15034"
          },
          {
            "address": "AC96444A9083890B7C7B430DB17BF1D3EA5EED04",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "QzxYamaSVNx+Idlb9bCcYsPI4Oa/zQ+rx4MN+XA2xLI="
            },
            "voting_power": "649",
            "accum": "-3715"
          },
          {
            "address": "AFCFB9B8DF23559CB5EB97C7CF1D77B411B2F47F",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "hWd0eUfhxpAqDZrKTgXcZJZXv5R92mQ4JJvZcAab0Ic="
            },
            "voting_power": "34",
            "accum": "-13650"
          },
          {
            "address": "B0155252D73B7EEB74D2A8CC814397E66970A839",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "w3rKE+tQoLK8G+XPmjn+NszCk07iQ0sWaBbN5hQZcBY="
            },
            "voting_power": "124",
            "accum": "14299"
          },
          {
            "address": "C39A8A2B85198DCB8F23AD1F6C198E2BA7AB5D60",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "R/3f7VruxWpu+2hiHlVpplTwoOou5kfQI1k/6/9H/y8="
            },
            "voting_power": "181",
            "accum": "-17780"
          },
          {
            "address": "C65E1E696C038A9A4A3565014ED168E97E7E63DF",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "pdxc69pGnaiK9oF+a49PjCnNww2PLYRBWlPFW/X6HNY="
            },
            "voting_power": "10",
            "accum": "-17159"
          },
          {
            "address": "D31B1198AE055BC52C600C19751011D95A64298A",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "NErIb5Hsp4ai02anNfdYui0Lm38uno+YzVJXO3xWS2g="
            },
            "voting_power": "7571",
            "accum": "20483"
          },
          {
            "address": "D6285C3422C8886445D53757CBE12D2B91B5A9D0",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "NQX4yKpOztKrmgBhGIC5WOALOLOq3LTpbzsN4ZLXGec="
            },
            "voting_power": "50",
            "accum": "4194"
          },
          {
            "address": "D77FB84A729972CD3E757C9BEB61561855832B92",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "UjTvuOew2EaooduJBiYmBWeF5ai0yFJG8uio5YXpJgg="
            },
            "voting_power": "68",
            "accum": "-26029"
          },
          {
            "address": "D981CB6DDF336E311A33C585C23C10F9F033A305",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "0ytofYCtFJT6zOYl0PMNqHan4cUwZg6ZnXHksICYfJ0="
            },
            "voting_power": "161",
            "accum": "19354"
          },
          {
            "address": "DA38013D7BCC3EC46109F8CDDF14F8D080CC4A9D",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "q+y8mjbjNKBC3HSWgVC8HVU34QC/DQDxtuSmx9eL4qI="
            },
            "voting_power": "14",
            "accum": "-24578"
          },
          {
            "address": "DB71F2DC8C967F0BABE9EF797760F17EF3D99ACE",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "u7qa9ST9HyzuL+xj3GyEoC+uBD52lVogDZRka26XxRI="
            },
            "voting_power": "52",
            "accum": "1247"
          },
          {
            "address": "DBA70FA7E9D55E035AD87B41C4DC0C38511FD09A",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "3O/S1+hdkUi5BxAohHHkxytw7S+tukjby46dRr6VhPE="
            },
            "voting_power": "3648",
            "accum": "-3558"
          },
          {
            "address": "DDDC63E5B9C7D1509805601E2EF6934BAEAC9115",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "pOYVpP34JYpnOMUbDzfbFHvLA2vfgJ8n/eknvqPmBS4="
            },
            "voting_power": "28",
            "accum": "-3310"
          },
          {
            "address": "E39B5778E0D4297A46F45CF6940B0A163A430FF8",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "0HqB2x6x5HzeozpHatePECw07x1UcDdSz8kQGNznnA8="
            },
            "voting_power": "1768",
            "accum": "-21526"
          },
          {
            "address": "E3B19738F7132F950668504D326C7660F159D170",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "tOEqjO2t51PEgO9Tv0B7qM0yPmy1n5tMa3Beg0tp3ns="
            },
            "voting_power": "33",
            "accum": "11453"
          },
          {
            "address": "E42F8AEABB685219F7027F54D2B0DBC7846F633A",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "eeImG09hOPo1W7j7lKepN/Lx6I9GGHqVBVEKmznxACc="
            },
            "voting_power": "240",
            "accum": "-4852"
          },
          {
            "address": "E4E6C2E81269DFFA196C17F98240E491479A6008",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "ydjx2ea+PVuChrny6X2dluJwyXta+BsNQRsgHXp8fXw="
            },
            "voting_power": "19",
            "accum": "-10061"
          },
          {
            "address": "E50ADECD5FD27FA2CD610C07F8AED36A2FC4A6A6",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "BaaCxmYHKJ6obIzTCdRtjw1cc8d2mUJcMbLWCjf1aLo="
            },
            "voting_power": "7171",
            "accum": "19713"
          },
          {
            "address": "E56E00A6B6AC1BE546CE9A9F34C7C1BB7585A8C1",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "TBsp7W9VwXn2YhLL5Dn9Ihpgy4dr+5bgT3uy5g+fjtQ="
            },
            "voting_power": "74",
            "accum": "-11220"
          },
          {
            "address": "EE67A5CA3269434AFAA41CE43E17B467488B5210",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "CEQGeQBsQNFXg9XpExD3cfD8ucyiAB5vGw6o/Ux76q8="
            },
            "voting_power": "179",
            "accum": "-13497"
          },
          {
            "address": "F23FF36BD5B90C33CE3A03ED72DBDCF5EC07D6AF",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "2JoNf1gavJ1d6XFIumO1Mki5GVMOcg58AioHksU3maE="
            },
            "voting_power": "199",
            "accum": "-1318"
          },
          {
            "address": "F4268597DE93E3E2CED38845010F8AA34E2A879E",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "Epd2FDKZwDybzT38Z7WB0y2jiLn9/2OLzmY3Zu18l6I="
            },
            "voting_power": "325",
            "accum": "8999"
          },
          {
            "address": "F6738260186D33D9C14FC6E7017AFE6BB952A63D",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "2qtEBT+Tc+SD2wJsdrVMHXrBKfvesxtmtSKDK5fXwA0="
            },
            "voting_power": "25",
            "accum": "-21887"
          }
        ],
        "proposer": {
          "address": "5BDF8EA41B1835FBAD865B181A2A4B23F511B5E4",
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "RBuRRNEzA9RA1Wrdi9PPFQJ29/n/bqN9O2tQv9Gq248="
          },
          "voting_power": "6040",
          "accum": "-21204"
        }
      },
      "proposal": null,
      "proposal_block": null,
      "proposal_block_parts": null,
      "locked_round": "0",
      "locked_block": null,
      "locked_block_parts": null,
      "valid_round": "0",
      "valid_block": null,
      "valid_block_parts": null,
      "votes": [
        {
          "round": "0",
          "prevotes": [
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote"
          ],
          "prevotes_bit_array": "BA{82:__________________________________________________________________________________} 0/46753 = 0.00",
          "precommits": [
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote",
            "nil-Vote"
          ],
          "precommits_bit_array": "BA{82:__________________________________________________________________________________} 0/46753 = 0.00"
        }
      ],
      "commit_round": "-1",
      "last_commit": {
        "votes": [
          "Vote{0:0077F68359FE 870166/01/2(Precommit) 18463C2EAE37 C518DCE7D574 @ 2018-10-26T03:52:42.189950373Z}",
          "Vote{1:0081A939ABBB 870166/01/2(Precommit) 18463C2EAE37 9DA68C9E6C26 @ 2018-10-26T03:52:42.067800178Z}",
          "Vote{2:01F78669F951 870166/01/2(Precommit) 18463C2EAE37 5016BD69B947 @ 2018-10-26T03:52:42.101701501Z}",
          "Vote{3:0548261C5022 870166/01/2(Precommit) 18463C2EAE37 B8C414384A15 @ 2018-10-26T03:52:42.130751353Z}",
          "Vote{4:066C70AF46E5 870166/01/2(Precommit) 18463C2EAE37 10C8CF3D0AD3 @ 2018-10-26T03:52:42.200738426Z}",
          "Vote{5:0A692DBFFE9E 870166/01/2(Precommit) 18463C2EAE37 F9679AE0BA45 @ 2018-10-26T03:52:42.282420641Z}",
          "Vote{6:0DDF97111D73 870166/01/2(Precommit) 18463C2EAE37 D87692048BC3 @ 2018-10-26T03:52:41.976166483Z}",
          "Vote{7:1ECD8C5FE134 870166/01/2(Precommit) 18463C2EAE37 8F03E3E3B838 @ 2018-10-26T03:52:42.235497128Z}",
          "nil-Vote",
          "Vote{9:296869AB66F7 870166/01/2(Precommit) 18463C2EAE37 E064634FF761 @ 2018-10-26T03:52:42.005874479Z}",
          "Vote{10:29B93E4EECD8 870166/01/2(Precommit) 18463C2EAE37 AAB84E57A7D5 @ 2018-10-26T03:52:42.057300716Z}",
          "Vote{11:2E6AF0D5B1A8 870166/01/2(Precommit) 18463C2EAE37 CFD8C8DC8A76 @ 2018-10-26T03:52:42.057249847Z}",
          "Vote{12:33412C9A69C6 870166/01/2(Precommit) 18463C2EAE37 A3C9E81E52CE @ 2018-10-26T03:52:42.128096872Z}",
          "Vote{13:3363E8F97B02 870166/01/2(Precommit) 18463C2EAE37 DF8DE432A813 @ 2018-10-26T03:52:42.100909943Z}",
          "Vote{14:36E9B7FAEC96 870166/01/2(Precommit) 18463C2EAE37 9C369F2F45C0 @ 2018-10-26T03:52:42.128502084Z}",
          "Vote{15:398610C4CF11 870166/01/2(Precommit) 18463C2EAE37 84462449A4EC @ 2018-10-26T03:52:37.293349454Z}",
          "Vote{16:3AC99B708404 870166/01/2(Precommit) 18463C2EAE37 0ECC5C462BD4 @ 2018-10-26T03:52:42.202258147Z}",
          "Vote{17:3E899BA36A93 870166/01/2(Precommit) 18463C2EAE37 231B06F4B974 @ 2018-10-26T03:52:42.115752723Z}",
          "Vote{18:411C43003831 870166/01/2(Precommit) 18463C2EAE37 118D7A8F4B38 @ 2018-10-26T03:52:42.100518689Z}",
          "Vote{19:4944F8BBBB53 870166/01/2(Precommit) 18463C2EAE37 1520E269B87A @ 2018-10-26T03:52:41.999978795Z}",
          "Vote{20:4B9C6FF8B36E 870166/01/2(Precommit) 18463C2EAE37 237BEF024FD9 @ 2018-10-26T03:52:42.100321876Z}",
          "Vote{21:50C18305F49B 870166/01/2(Precommit) 18463C2EAE37 587575868A39 @ 2018-10-26T03:52:42.139732103Z}",
          "Vote{22:5104F488EA6B 870166/01/2(Precommit) 18463C2EAE37 8C509310B4A1 @ 2018-10-26T03:52:42.029139417Z}",
          "Vote{23:570B6A755ED3 870166/01/2(Precommit) 18463C2EAE37 A7DF6F5D46BF @ 2018-10-26T03:52:44.06663216Z}",
          "Vote{24:5B9628695CDF 870166/01/2(Precommit) 18463C2EAE37 ED0435787A50 @ 2018-10-26T03:52:12.614433636Z}",
          "Vote{25:5BDF8EA41B18 870166/01/2(Precommit) 18463C2EAE37 2C383F23D5D1 @ 2018-10-26T03:52:41.990302733Z}",
          "Vote{26:5DE9674DA2D5 870166/01/2(Precommit) 18463C2EAE37 EE14071EEF5C @ 2018-10-26T03:52:42.07235729Z}",
          "Vote{27:5E06FB65FD0E 870166/01/2(Precommit) 18463C2EAE37 C527164286E9 @ 2018-10-26T03:52:41.965453549Z}",
          "Vote{28:5E8673673E37 870166/01/2(Precommit) 18463C2EAE37 2682E35873C1 @ 2018-10-26T03:52:42.081514045Z}",
          "Vote{29:651995D8471D 870166/01/2(Precommit) 18463C2EAE37 522602AEF363 @ 2018-10-26T03:52:42.018934378Z}",
          "Vote{30:668309BE1A6F 870166/01/2(Precommit) 18463C2EAE37 5D853B6CF2AA @ 2018-10-26T03:52:42.098763494Z}",
          "Vote{31:69D99B2C6604 870166/01/2(Precommit) 18463C2EAE37 381DEF079EF8 @ 2018-10-26T03:52:41.996341936Z}",
          "Vote{32:69EDADFA4A80 870166/01/2(Precommit) 18463C2EAE37 33BB8D20C1E1 @ 2018-10-26T03:52:42.223120502Z}",
          "Vote{33:6B275FFA0571 870166/01/2(Precommit) 18463C2EAE37 BEBD4A1CDB95 @ 2018-10-26T03:52:42.03667375Z}",
          "Vote{34:6D3AD380E398 870166/01/2(Precommit) 18463C2EAE37 D19344256C81 @ 2018-10-26T03:50:54.681101074Z}",
          "Vote{35:6F6C93B8BF34 870166/01/2(Precommit) 18463C2EAE37 84EAEC34ABDC @ 2018-10-26T03:52:42.190635843Z}",
          "Vote{36:71BC71D28E45 870166/01/2(Precommit) 18463C2EAE37 53FFF13C3D47 @ 2018-10-26T03:52:42.041132395Z}",
          "Vote{37:732CEEF54C37 870166/01/2(Precommit) 18463C2EAE37 A90086DF4BD3 @ 2018-10-26T03:52:42.016726928Z}",
          "Vote{38:7A20421F99FC 870166/01/2(Precommit) 18463C2EAE37 6E44BD336E62 @ 2018-10-26T03:52:44.706679123Z}",
          "Vote{39:7A32D912F7A7 870166/01/2(Precommit) 18463C2EAE37 4B3D3036B712 @ 2018-10-26T03:52:42.080837884Z}",
          "Vote{40:7EB7593F519F 870166/01/2(Precommit) 18463C2EAE37 31D5F931D617 @ 2018-10-26T03:52:42.197659052Z}",
          "Vote{41:807CF675BA84 870166/01/2(Precommit) 18463C2EAE37 4CD5170EAAFB @ 2018-10-26T03:52:42.269727839Z}",
          "Vote{42:8340F82D32C4 870166/01/2(Precommit) 18463C2EAE37 616D667BF8DF @ 2018-10-26T03:52:42.131325444Z}",
          "nil-Vote",
          "Vote{44:850D387A41AD 870166/01/2(Precommit) 18463C2EAE37 46D53C91402B @ 2018-10-26T03:52:42.165995379Z}",
          "Vote{45:85335B5077FA 870166/01/2(Precommit) 18463C2EAE37 C4CFC21BD6B3 @ 2018-10-26T03:52:42.116412347Z}",
          "Vote{46:86DDBEEF35D3 870166/01/2(Precommit) 18463C2EAE37 7E6C3DB7DED5 @ 2018-10-26T03:52:41.993818104Z}",
          "Vote{47:89B357643A94 870166/01/2(Precommit) 18463C2EAE37 8A53725F69BF @ 2018-10-26T03:52:42.038224523Z}",
          "Vote{48:8DF81F7913F1 870166/01/2(Precommit) 18463C2EAE37 F7AC1BD15416 @ 2018-10-26T03:52:42.056282558Z}",
          "Vote{49:8EA62D00960B 870166/01/2(Precommit) 18463C2EAE37 C828A59091A5 @ 2018-10-26T03:52:42.199660098Z}",
          "Vote{50:8F06FE1FD042 870166/01/2(Precommit) 18463C2EAE37 3805491F3F26 @ 2018-10-26T03:52:42.01093748Z}",
          "Vote{51:8F302C083184 870166/01/2(Precommit) 18463C2EAE37 5C97D0D535C5 @ 2018-10-26T03:52:42.026036164Z}",
          "Vote{52:927CC759310F 870166/01/2(Precommit) 18463C2EAE37 B113AD1F8519 @ 2018-10-26T03:44:41.919412991Z}",
          "Vote{53:9524C2BBF5C9 870166/01/2(Precommit) 18463C2EAE37 403DF323B43C @ 2018-10-26T03:52:42.012608011Z}",
          "Vote{54:96BDE45EC526 870166/01/2(Precommit) 18463C2EAE37 D00F9E67C576 @ 2018-10-26T03:52:42.007833207Z}",
          "Vote{55:989E892A603E 870166/01/2(Precommit) 18463C2EAE37 F886458E36F3 @ 2018-10-26T03:52:42.164627682Z}",
          "Vote{56:9D93AD02FE1B 870166/01/2(Precommit) 18463C2EAE37 D84D5035A365 @ 2018-10-26T03:52:41.993763921Z}",
          "Vote{57:A0ED8A257D78 870166/01/2(Precommit) 18463C2EAE37 6F9290EC8F7C @ 2018-10-26T03:52:42.100837729Z}",
          "Vote{58:A5868971300F 870166/01/2(Precommit) 18463C2EAE37 886384778C93 @ 2018-10-26T03:52:41.977351828Z}",
          "Vote{59:AC96444A9083 870166/01/2(Precommit) 18463C2EAE37 814B218E1005 @ 2018-10-26T03:52:42.25660217Z}",
          "Vote{60:AFCFB9B8DF23 870166/01/2(Precommit) 18463C2EAE37 3FEA1DA172B0 @ 2018-10-26T03:52:42.080357596Z}",
          "Vote{61:B0155252D73B 870166/01/2(Precommit) 18463C2EAE37 0518FB3C9B3D @ 2018-10-26T03:52:42.006185788Z}",
          "Vote{62:C39A8A2B8519 870166/01/2(Precommit) 18463C2EAE37 C5BBEAE357D9 @ 2018-10-26T03:52:41.985158559Z}",
          "Vote{63:C65E1E696C03 870166/01/2(Precommit) 18463C2EAE37 D2D9C7BCD790 @ 2018-10-26T03:52:42.055016597Z}",
          "Vote{64:D31B1198AE05 870166/01/2(Precommit) 18463C2EAE37 2233E79EAECE @ 2018-10-26T03:52:42.168394334Z}",
          "Vote{65:D6285C3422C8 870166/01/2(Precommit) 18463C2EAE37 A7848FF7EA9E @ 2018-10-26T03:52:42.094849005Z}",
          "Vote{66:D77FB84A7299 870166/01/2(Precommit) 18463C2EAE37 F9DC9B654782 @ 2018-10-26T03:52:41.97533167Z}",
          "Vote{67:D981CB6DDF33 870166/01/2(Precommit) 18463C2EAE37 4EE551B75C2A @ 2018-10-26T03:52:42.05154154Z}",
          "Vote{68:DA38013D7BCC 870166/01/2(Precommit) 18463C2EAE37 30030B9564C6 @ 2018-10-26T03:52:42.005390795Z}",
          "Vote{69:DB71F2DC8C96 870166/01/2(Precommit) 18463C2EAE37 465F208D4942 @ 2018-10-26T03:52:42.077737592Z}",
          "Vote{70:DBA70FA7E9D5 870166/01/2(Precommit) 18463C2EAE37 7DB148F29BD2 @ 2018-10-26T03:52:42.149066786Z}",
          "Vote{71:DDDC63E5B9C7 870166/01/2(Precommit) 18463C2EAE37 700A03AC5E98 @ 2018-10-26T03:52:41.965200538Z}",
          "Vote{72:E39B5778E0D4 870166/01/2(Precommit) 18463C2EAE37 5AA9A110A849 @ 2018-10-26T03:52:42.006840979Z}",
          "Vote{73:E3B19738F713 870166/01/2(Precommit) 18463C2EAE37 C99E1D71CEB4 @ 2018-10-26T03:52:41.95015111Z}",
          "Vote{74:E42F8AEABB68 870166/01/2(Precommit) 18463C2EAE37 9D8D6797AB04 @ 2018-10-26T03:52:42.021793181Z}",
          "Vote{75:E4E6C2E81269 870166/01/2(Precommit) 18463C2EAE37 A12E2F93C1E7 @ 2018-10-26T03:52:42.06499155Z}",
          "Vote{76:E50ADECD5FD2 870166/01/2(Precommit) 18463C2EAE37 09E21A1120FD @ 2018-10-26T03:52:42.033090607Z}",
          "Vote{77:E56E00A6B6AC 870166/01/2(Precommit) 18463C2EAE37 0259ABA8DE3C @ 2018-10-26T03:52:42.038447217Z}",
          "Vote{78:EE67A5CA3269 870166/01/2(Precommit) 18463C2EAE37 BCDFE0B5E28C @ 2018-10-26T03:52:41.988588971Z}",
          "Vote{79:F23FF36BD5B9 870166/01/2(Precommit) 18463C2EAE37 E532CD6460BE @ 2018-10-26T03:52:42.059432589Z}",
          "Vote{80:F4268597DE93 870166/01/2(Precommit) 18463C2EAE37 14E654A9736C @ 2018-10-26T03:52:42.021015532Z}",
          "Vote{81:F6738260186D 870166/01/2(Precommit) 18463C2EAE37 1122D67DC6AF @ 2018-10-26T03:52:42.002597132Z}"
        ],
        "votes_bit_array": "BA{82:xxxxxxxx_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx} 45768/46753 = 0.98",
        "peer_maj_23s": {}
      },
      "last_validators": {
        "validators": [
          {
            "address": "0077F68359FE12C678980AC3442460CB1CC1ABC4",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "zf+aS3anGNm9CrpliDTqpUlHCBpi0qZknUo+aYOXC+s="
            },
            "voting_power": "169",
            "accum": "4533"
          },
          {
            "address": "0081A939ABBB733E8DC36E406EC51E4309B9AE3A",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "Sk7VmvaRgatGyVrbkbQ4QBCeda/DIUs6E3NzstWQV6g="
            },
            "voting_power": "154",
            "accum": "8340"
          },
          {
            "address": "01F78669F9515FD83DF9250F5C0EE143D3DAD65C",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "1K9/Z+oiK380Obf+LYVN6P7ydFIWCFvNQpT+ToRUURg="
            },
            "voting_power": "627",
            "accum": "-11694"
          },
          {
            "address": "0548261C50222FD710EB5EBDF03A0E6567B21D20",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "8K3clCjVU33BTIpUhdahGmu++WxHj4NUE9krCRkk++s="
            },
            "voting_power": "1181",
            "accum": "-7607"
          },
          {
            "address": "066C70AF46E5E1B473BD125FD8937EF90E555641",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "To+MjLVGM4gNuE7JahfDwIc5Q+31H+Q60Cy4AoU2fjE="
            },
            "voting_power": "5",
            "accum": "14618"
          },
          {
            "address": "0A692DBFFE9E5DD28E7DECC33CED1AEB4C0D014E",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "N3K5kDdfcKJurfaa6s2zfKgtYvz1Pagz7VWi9ZfX8yM="
            },
            "voting_power": "439",
            "accum": "-4459"
          },
          {
            "address": "0DDF97111D73DB825138EC15EC27DE5FE8536901",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "H0SIA/BU6Xj8oT5bQkvLpEITN3CqFLbMeBcQ72NZrAE="
            },
            "voting_power": "18",
            "accum": "15018"
          },
          {
            "address": "1ECD8C5FE1341D74CDC633A245DD56858DE2F9A2",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "lUMCRAeu47BsOhNvCQTQJQeB68z0/VaElC9j5gDt9y8="
            },
            "voting_power": "578",
            "accum": "5772"
          },
          {
            "address": "20FBB777AB2676E19A31E2EB35B2C7B426C99C13",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "/DtRehBSochyxchfAn/Tb4tbwC9u6Ryq++VeBE72FuA="
            },
            "voting_power": "39",
            "accum": "14139"
          },
          {
            "address": "296869AB66F7003A5AADA1A8DEEEA2913486668F",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "4cI0avx/jOL9eWe7PyRm4Mhv1QSSrhxpyr30f7tJ7i0="
            },
            "voting_power": "10",
            "accum": "-24633"
          },
          {
            "address": "29B93E4EECD8C591AB76C2D52FD7C43CBEEEC50B",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "GTwvM3OOWuXdwH7suTJIWdYdtwUU1hLAHIXmpdzGfl0="
            },
            "voting_power": "7",
            "accum": "10596"
          },
          {
            "address": "2E6AF0D5B1A85E173F5EBEDA93D7E7D97A88D06C",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "G3YJxcM2oMCW8R5s/HSqmnudsQjcj+e/vpZBtMyob1w="
            },
            "voting_power": "121",
            "accum": "-5889"
          },
          {
            "address": "33412C9A69C60CB0FAA24A5C5B3A906DE24B47C2",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "4DEDoU/RsHMPS54GzgwkWnW5zPQfMt9aInFFc3GyfA8="
            },
            "voting_power": "40",
            "accum": "-24273"
          },
          {
            "address": "3363E8F97B02ECC00289E72173D827543047ACDA",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "mPnu910hOOa1tAQ7pbOLFDxvllbQUmrbtGjqQrYg1nM="
            },
            "voting_power": "89",
            "accum": "5362"
          },
          {
            "address": "36E9B7FAEC964F97D84C028ED62493E373AD0B38",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "ViJzrw9Unk1XyDg4f0UQLFE4zfUaFh23wZZY0f31Qg0="
            },
            "voting_power": "139",
            "accum": "10466"
          },
          {
            "address": "398610C4CF11C84C89AC6975470972EE75DA17E4",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "5GAxIHzu06l0n3MU8wNQrCcrrnpZR9EKe5qZMsM2h/E="
            },
            "voting_power": "490",
            "accum": "-7136"
          },
          {
            "address": "3AC99B708404452366E0953A4BE849226DC86CBC",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "ISIM341M+EUYOMpwhn9gyCvRUZ06xomp/+lOwa50meQ="
            },
            "voting_power": "480",
            "accum": "-9275"
          },
          {
            "address": "3E899BA36A93C7370347BD5FE85909F2ECD93F8C",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "M2Iu/A8RQcmoiIEar/iV4wBjVj0I0gcd/nNnvwmnT2g="
            },
            "voting_power": "39",
            "accum": "1427"
          },
          {
            "address": "411C430038318FE4DB32AFC508A6C12782161294",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "LmJN98a+ce+shthBz95qKsdwUwWRs1QdBLO4HGNQDpM="
            },
            "voting_power": "107",
            "accum": "-25068"
          },
          {
            "address": "4944F8BBBB5303E17BC0607DF332320EDBC750EA",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "A6GzeXUM3vsXaDAEYMSDgSKkqn9AoUYjs8empH46MGY="
            },
            "voting_power": "1288",
            "accum": "8519"
          },
          {
            "address": "4B9C6FF8B36EE6134AA8D822772C0162C471186D",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "g0f/B/FxRDU0VbS7vMXs8cJwSUITpbtn3NjEQ0NJ5Fg="
            },
            "voting_power": "20",
            "accum": "-8568"
          },
          {
            "address": "50C18305F49BE332E4C6BCDE0DEECD4A22B1D6A4",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "PX8v6nM+dk8fnqQWuXaRpRrdlhL+sMUHYo75CJbnY9U="
            },
            "voting_power": "109",
            "accum": "17858"
          },
          {
            "address": "5104F488EA6B1E2C33CBDA3C18389C88104DBC72",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "JREG9eji9gfJR2b61qMu8jNZ835ezPY2A4o+zjYMfDg="
            },
            "voting_power": "299",
            "accum": "7005"
          },
          {
            "address": "570B6A755ED339EAC56BB3390A48069F6AC9449B",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "5XCMLGNTGkMh9/JsUzi/8Xl7Y0cdDe38stX+zvBLXuE="
            },
            "voting_power": "44",
            "accum": "16034"
          },
          {
            "address": "5B9628695CDFC6025857578114E3D3126687EAE1",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "NCwjL8K9R2CLOMYub3MhAkkB0fktT0mZC75TEtcXyQ4="
            },
            "voting_power": "130",
            "accum": "15944"
          },
          {
            "address": "5BDF8EA41B1835FBAD865B181A2A4B23F511B5E4",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "RBuRRNEzA9RA1Wrdi9PPFQJ29/n/bqN9O2tQv9Gq248="
            },
            "voting_power": "6040",
            "accum": "19509"
          },
          {
            "address": "5DE9674DA2D535D5B1547D90990E9087EAE15E39",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "yRk1EO5Ta0iCI68BjluGEyZIlkBTGvUzbl+s/Rujhco="
            },
            "voting_power": "5",
            "accum": "14618"
          },
          {
            "address": "5E06FB65FD0E0EBB11D0D74A1E480B4A8D7D0FE7",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "iPPJziJLaL5QA6Mr6uaUGwoNrGts45W6LVIKbo11GfA="
            },
            "voting_power": "41",
            "accum": "-4339"
          },
          {
            "address": "5E8673673E37450F01B64138FBF4B172BDB52D68",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "CuPin/hLM2tK6KvcUlOi/Brm5xNi3zq6UPgpLFPSzt8="
            },
            "voting_power": "776",
            "accum": "-19577"
          },
          {
            "address": "651995D8471D07A22C149265F5237A89149795C1",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "WX9FldKAX82ALt3u3v80v/1zLUNJQO689C9vwXooDuQ="
            },
            "voting_power": "142",
            "accum": "18069"
          },
          {
            "address": "668309BE1A6F5317197B0BC6E47153D249DD5DE9",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "DKXu53utgN9Q/4dd7uRC0c8j1bNHTGSVDzhlEcrRFl8="
            },
            "voting_power": "103",
            "accum": "-18873"
          },
          {
            "address": "69D99B2C66043ACBEAA8447525C356AFC6408E0C",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "8pfpbIxBBiu88hpxS3CeRpv7kClEjl8SwVgckDNBGlE="
            },
            "voting_power": "541",
            "accum": "-951"
          },
          {
            "address": "69EDADFA4A803E9DBF548C900D4A465E7D154262",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "LRK6flgMsHrBYkIIahegEkzlwC7gQ3dAyw0jcAd1u/k="
            },
            "voting_power": "107",
            "accum": "-25119"
          },
          {
            "address": "6B275FFA0571EEFC4482CD7A6CAE836F5785CB4E",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "pggq0WUTlsiA8cBiptRgxlw4WUG2GmXfLQxkbMbCBPk="
            },
            "voting_power": "202",
            "accum": "15955"
          },
          {
            "address": "6D3AD380E398D8592B03AA124BAFED5F125DF863",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "ENAVynNXVpj/IdYx9kCPKaPs4bWSxRIHNlmS9QiDuZQ="
            },
            "voting_power": "54",
            "accum": "-2943"
          },
          {
            "address": "6F6C93B8BF3495063FEC9B5C771158607F4AE9D6",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "S8s6fdAQNQ3bN9SNVAsHB/j8uv1CM1roxeLesL+fh4g="
            },
            "voting_power": "141",
            "accum": "6930"
          },
          {
            "address": "71BC71D28E4599BCA22214BBB2CE0DFEE90A9EE4",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "MWID5oHOb367w/js26+cbQdyGqw2KEpe967D44z/eSQ="
            },
            "voting_power": "1",
            "accum": "-24088"
          },
          {
            "address": "732CEEF54C374DDC6ADECBFD707AEFD07FEDC143",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "HjSC7VkhKih6xMhudlqfaFE8ZZnP8RKJPv4iqR7RhcE="
            },
            "voting_power": "357",
            "accum": "-7507"
          },
          {
            "address": "7A20421F99FC4362B13C2FD43548F3E75B057017",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "HWtL6kXsdXL6Lxoq4XXGG3FkpPAmDBFwrP1ndYCvsnA="
            },
            "voting_power": "20",
            "accum": "11528"
          },
          {
            "address": "7A32D912F7A790EEE5610653DC4DDBC0DB4F2497",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "d5o/ErYmwkS/Cw06SH1l3vrBDXRUJLtjj8kcz9CEiFQ="
            },
            "voting_power": "79",
            "accum": "1404"
          },
          {
            "address": "7EB7593F519F1180C7E872D0494611464EDF03AE",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "VakMQSPBEuSC9Nwuv8WWhrZVUmH31bUR4+G6pJhkgE8="
            },
            "voting_power": "1807",
            "accum": "1324"
          },
          {
            "address": "807CF675BA8479FC4FEC1047215DF07970040F1F",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "2p8s/pRBZPjYWKKMlR7AOXypDzDmPo762iXlKpCwtco="
            },
            "voting_power": "14",
            "accum": "-24592"
          },
          {
            "address": "8340F82D32C4A221C57FC33FCBA44309C4DF1BBD",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "FSj9jMpCOdWfX0rCQ0TpnEa8/uOroQl4V4Bk3Msck3w="
            },
            "voting_power": "65",
            "accum": "3263"
          },
          {
            "address": "83AB321E012C57746689FED3A6064DF331F81007",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "aUViYC2znC55sleHfmsIN9hZ45SbYPbDcYA0gVzglsc="
            },
            "voting_power": "946",
            "accum": "-372"
          },
          {
            "address": "850D387A41ADB21F7D67EAC5076D594DF3D861C7",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "Sowu0HYQbGFJ4dA3urnWwXqMc1+6kzQHre+9/iVlpjg="
            },
            "voting_power": "10",
            "accum": "376"
          },
          {
            "address": "85335B5077FAE8E2926AD338634079124EE1C4A1",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "lGnYgu+rv9HZj3Y+fS6Ad/GszZgtwwyZcM9mAc/tuu4="
            },
            "voting_power": "475",
            "accum": "-23977"
          },
          {
            "address": "86DDBEEF35D35786A68834B34869174AC314BB74",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "ilDFePqWkNRNWUEoWKq42jZC5CD4rglfs7+gCuq6Qc8="
            },
            "voting_power": "1867",
            "accum": "-12286"
          },
          {
            "address": "89B357643A945BD433CB5B2B9CBD9B90CB9CCD5E",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "UVRp+iVQUv9erXJwqfI6ClWIXk/8fSs0CRudMMU49Zo="
            },
            "voting_power": "36",
            "accum": "-17432"
          },
          {
            "address": "8DF81F7913F1EBF856618B1475C12EEEC1214BB0",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "0WFwqm+9/VGNlZDcAlkqCdQm/8M+W19LWnzXIDg/py0="
            },
            "voting_power": "221",
            "accum": "6058"
          },
          {
            "address": "8EA62D00960B7557560E0D78278B6B36165AD0A8",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "43IDu79dbTK5jDg1x2xR09pzWgWdAGHoIrz4Hk9H//Q="
            },
            "voting_power": "10",
            "accum": "-1694"
          },
          {
            "address": "8F06FE1FD042683C426137FBC4BBAA4288924217",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "lB+CgSbTzpHXJHUU15fJIPaPl6XF0nSfdNxBnlKJqTk="
            },
            "voting_power": "2466",
            "accum": "389"
          },
          {
            "address": "8F302C08318444DDC43EABDE96E869D0B1F5A421",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "+yk80I86yZ+wqwf/9iDEfewwZGDaXCKosDuWZnlSf4M="
            },
            "voting_power": "110",
            "accum": "-1876"
          },
          {
            "address": "927CC759310F4DF5C8936EF6D666ED3135D609B7",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "tkdOxR2CIOctyc5e9IMGwnMwB420OTN3Rfb6fUMVhr8="
            },
            "voting_power": "11",
            "accum": "3323"
          },
          {
            "address": "9524C2BBF5C9BCC7AFEE121AF03E31D8480DD5E5",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "nWft/3AE/5/Iv5ePNUbkFHeG8HpNnnIFA3hQKR6qRmQ="
            },
            "voting_power": "5",
            "accum": "14618"
          },
          {
            "address": "96BDE45EC52606CF5CB59853097871AD60680549",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "4JoJuRfaANhdM1x3AWRo1/Cj9DH3VA+fi1SynzknV+w="
            },
            "voting_power": "19",
            "accum": "-10080"
          },
          {
            "address": "989E892A603E6BA7F325BC52555DFB47A8751E66",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "zkQOeG4/QylfQCqxV8WPzC0bUvIkRb467gbL80m4/KA="
            },
            "voting_power": "10",
            "accum": "-17169"
          },
          {
            "address": "9D93AD02FE1BD7E76B9E58D7EB461088EAD011E5",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "jj0Y/Fy8JSJR3g+PHU6Ce0ecYwHGUVJ4bVyR7WwcyLI="
            },
            "voting_power": "614",
            "accum": "-13560"
          },
          {
            "address": "A0ED8A257D7862094FE8015BCE2135D111722DF6",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "dvLH3HPQ3YH0oDefmGTPvpwn/gTFhvBAQiYNV4DVs5k="
            },
            "voting_power": "195",
            "accum": "5503"
          },
          {
            "address": "A5868971300F155020634D7BB49C7634F7AACCD3",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "cCoFsZzKZ9SQZbHe4NueVObIezP6ts0tRTZ/aN96dig="
            },
            "voting_power": "18",
            "accum": "15016"
          },
          {
            "address": "AC96444A9083890B7C7B430DB17BF1D3EA5EED04",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "QzxYamaSVNx+Idlb9bCcYsPI4Oa/zQ+rx4MN+XA2xLI="
            },
            "voting_power": "649",
            "accum": "-4364"
          },
          {
            "address": "AFCFB9B8DF23559CB5EB97C7CF1D77B411B2F47F",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "hWd0eUfhxpAqDZrKTgXcZJZXv5R92mQ4JJvZcAab0Ic="
            },
            "voting_power": "34",
            "accum": "-13684"
          },
          {
            "address": "B0155252D73B7EEB74D2A8CC814397E66970A839",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "w3rKE+tQoLK8G+XPmjn+NszCk07iQ0sWaBbN5hQZcBY="
            },
            "voting_power": "124",
            "accum": "14175"
          },
          {
            "address": "C39A8A2B85198DCB8F23AD1F6C198E2BA7AB5D60",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "R/3f7VruxWpu+2hiHlVpplTwoOou5kfQI1k/6/9H/y8="
            },
            "voting_power": "181",
            "accum": "-17961"
          },
          {
            "address": "C65E1E696C038A9A4A3565014ED168E97E7E63DF",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "pdxc69pGnaiK9oF+a49PjCnNww2PLYRBWlPFW/X6HNY="
            },
            "voting_power": "10",
            "accum": "-17169"
          },
          {
            "address": "D31B1198AE055BC52C600C19751011D95A64298A",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "NErIb5Hsp4ai02anNfdYui0Lm38uno+YzVJXO3xWS2g="
            },
            "voting_power": "7571",
            "accum": "12912"
          },
          {
            "address": "D6285C3422C8886445D53757CBE12D2B91B5A9D0",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "NQX4yKpOztKrmgBhGIC5WOALOLOq3LTpbzsN4ZLXGec="
            },
            "voting_power": "50",
            "accum": "4144"
          },
          {
            "address": "D77FB84A729972CD3E757C9BEB61561855832B92",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "UjTvuOew2EaooduJBiYmBWeF5ai0yFJG8uio5YXpJgg="
            },
            "voting_power": "68",
            "accum": "-26097"
          },
          {
            "address": "D981CB6DDF336E311A33C585C23C10F9F033A305",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "0ytofYCtFJT6zOYl0PMNqHan4cUwZg6ZnXHksICYfJ0="
            },
            "voting_power": "161",
            "accum": "19193"
          },
          {
            "address": "DA38013D7BCC3EC46109F8CDDF14F8D080CC4A9D",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "q+y8mjbjNKBC3HSWgVC8HVU34QC/DQDxtuSmx9eL4qI="
            },
            "voting_power": "14",
            "accum": "-24592"
          },
          {
            "address": "DB71F2DC8C967F0BABE9EF797760F17EF3D99ACE",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "u7qa9ST9HyzuL+xj3GyEoC+uBD52lVogDZRka26XxRI="
            },
            "voting_power": "52",
            "accum": "1195"
          },
          {
            "address": "DBA70FA7E9D55E035AD87B41C4DC0C38511FD09A",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "3O/S1+hdkUi5BxAohHHkxytw7S+tukjby46dRr6VhPE="
            },
            "voting_power": "3648",
            "accum": "-7206"
          },
          {
            "address": "DDDC63E5B9C7D1509805601E2EF6934BAEAC9115",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "pOYVpP34JYpnOMUbDzfbFHvLA2vfgJ8n/eknvqPmBS4="
            },
            "voting_power": "28",
            "accum": "-3338"
          },
          {
            "address": "E39B5778E0D4297A46F45CF6940B0A163A430FF8",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "0HqB2x6x5HzeozpHatePECw07x1UcDdSz8kQGNznnA8="
            },
            "voting_power": "1768",
            "accum": "-23294"
          },
          {
            "address": "E3B19738F7132F950668504D326C7660F159D170",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "tOEqjO2t51PEgO9Tv0B7qM0yPmy1n5tMa3Beg0tp3ns="
            },
            "voting_power": "33",
            "accum": "11420"
          },
          {
            "address": "E42F8AEABB685219F7027F54D2B0DBC7846F633A",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "eeImG09hOPo1W7j7lKepN/Lx6I9GGHqVBVEKmznxACc="
            },
            "voting_power": "240",
            "accum": "-5092"
          },
          {
            "address": "E4E6C2E81269DFFA196C17F98240E491479A6008",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "ydjx2ea+PVuChrny6X2dluJwyXta+BsNQRsgHXp8fXw="
            },
            "voting_power": "19",
            "accum": "-10080"
          },
          {
            "address": "E50ADECD5FD27FA2CD610C07F8AED36A2FC4A6A6",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "BaaCxmYHKJ6obIzTCdRtjw1cc8d2mUJcMbLWCjf1aLo="
            },
            "voting_power": "7171",
            "accum": "12542"
          },
          {
            "address": "E56E00A6B6AC1BE546CE9A9F34C7C1BB7585A8C1",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "TBsp7W9VwXn2YhLL5Dn9Ihpgy4dr+5bgT3uy5g+fjtQ="
            },
            "voting_power": "74",
            "accum": "-11294"
          },
          {
            "address": "EE67A5CA3269434AFAA41CE43E17B467488B5210",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "CEQGeQBsQNFXg9XpExD3cfD8ucyiAB5vGw6o/Ux76q8="
            },
            "voting_power": "179",
            "accum": "-13676"
          },
          {
            "address": "F23FF36BD5B90C33CE3A03ED72DBDCF5EC07D6AF",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "2JoNf1gavJ1d6XFIumO1Mki5GVMOcg58AioHksU3maE="
            },
            "voting_power": "199",
            "accum": "-1517"
          },
          {
            "address": "F4268597DE93E3E2CED38845010F8AA34E2A879E",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "Epd2FDKZwDybzT38Z7WB0y2jiLn9/2OLzmY3Zu18l6I="
            },
            "voting_power": "325",
            "accum": "8674"
          },
          {
            "address": "F6738260186D33D9C14FC6E7017AFE6BB952A63D",
            "pub_key": {
              "type": "tendermint/PubKeyEd25519",
              "value": "2qtEBT+Tc+SD2wJsdrVMHXrBKfvesxtmtSKDK5fXwA0="
            },
            "voting_power": "25",
            "accum": "-21912"
          }
        ],
        "proposer": {
          "address": "69EDADFA4A803E9DBF548C900D4A465E7D154262",
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "LRK6flgMsHrBYkIIahegEkzlwC7gQ3dAyw0jcAd1u/k="
          },
          "voting_power": "107",
          "accum": "-25119"
        }
      }
    },
    "peers": [
      {
        "node_address": "8e44e63e3b536dd780feff44fc83596c30affb7d@133.130.90.146:26656",
        "peer_state": {
          "round_state": {
            "height": "870167",
            "round": "0",
            "step": 1,
            "start_time": "2018-10-26T03:52:43.258003423Z",
            "proposal": false,
            "proposal_block_parts_header": {
              "total": "0",
              "hash": ""
            },
            "proposal_block_parts": null,
            "proposal_pol_round": "-1",
            "proposal_pol": "__________________________________________________________________________________",
            "prevotes": "__________________________________________________________________________________",
            "precommits": "__________________________________________________________________________________",
            "last_commit_round": "1",
            "last_commit": "xxxxxxxx_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
            "catchup_commit_round": "-1",
            "catchup_commit": "__________________________________________________________________________________"
          },
          "stats": {
            "last_vote_height": "870166",
            "votes": "34019",
            "last_block_part_height": "870166",
            "block_parts": "32999"
          }
        }
      },
      {
        "node_address": "7043e77f36609425116882066e6b0d1cf6a78185@209.250.244.51:26656",
        "peer_state": {
          "round_state": {
            "height": "870167",
            "round": "0",
            "step": 1,
            "start_time": "2018-10-26T03:52:47.041201633Z",
            "proposal": false,
            "proposal_block_parts_header": {
              "total": "0",
              "hash": ""
            },
            "proposal_block_parts": null,
            "proposal_pol_round": "-1",
            "proposal_pol": "__________________________________________________________________________________",
            "prevotes": "__________________________________________________________________________________",
            "precommits": "__________________________________________________________________________________",
            "last_commit_round": "1",
            "last_commit": "xxxxxxxx_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
            "catchup_commit_round": "-1",
            "catchup_commit": "__________________________________________________________________________________"
          },
          "stats": {
            "last_vote_height": "870166",
            "votes": "19336",
            "last_block_part_height": "870166",
            "block_parts": "19279"
          }
        }
      },
      {
        "node_address": "9c5c946df2cab55d111d1fa7e6b03c50666160c4@142.93.31.31:26656",
        "peer_state": {
          "round_state": {
            "height": "870167",
            "round": "0",
            "step": 1,
            "start_time": "2018-10-26T03:52:47.276022789Z",
            "proposal": false,
            "proposal_block_parts_header": {
              "total": "0",
              "hash": ""
            },
            "proposal_block_parts": null,
            "proposal_pol_round": "-1",
            "proposal_pol": "__________________________________________________________________________________",
            "prevotes": "__________________________________________________________________________________",
            "precommits": "__________________________________________________________________________________",
            "last_commit_round": "1",
            "last_commit": "xxxxxxxx_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
            "catchup_commit_round": "-1",
            "catchup_commit": "__________________________________________________________________________________"
          },
          "stats": {
            "last_vote_height": "870166",
            "votes": "17282",
            "last_block_part_height": "870166",
            "block_parts": "17257"
          }
        }
      },
      {
        "node_address": "d346857fed996ffa1a034f4c983e06eaf9c00664@35.180.139.254:26656",
        "peer_state": {
          "round_state": {
            "height": "870167",
            "round": "0",
            "step": 1,
            "start_time": "2018-10-26T03:52:46.861542535Z",
            "proposal": false,
            "proposal_block_parts_header": {
              "total": "0",
              "hash": ""
            },
            "proposal_block_parts": null,
            "proposal_pol_round": "-1",
            "proposal_pol": "__________________________________________________________________________________",
            "prevotes": "__________________________________________________________________________________",
            "precommits": "__________________________________________________________________________________",
            "last_commit_round": "1",
            "last_commit": "xxxxxxxx_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
            "catchup_commit_round": "-1",
            "catchup_commit": "__________________________________________________________________________________"
          },
          "stats": {
            "last_vote_height": "870166",
            "votes": "10657",
            "last_block_part_height": "870166",
            "block_parts": "10644"
          }
        }
      },
      {
        "node_address": "4a53d0a92deb68a0085aa0b7f100fc475af42b94@50.17.110.27:26656",
        "peer_state": {
          "round_state": {
            "height": "870167",
            "round": "0",
            "step": 1,
            "start_time": "2018-10-26T03:52:46.758307467Z",
            "proposal": false,
            "proposal_block_parts_header": {
              "total": "0",
              "hash": ""
            },
            "proposal_block_parts": null,
            "proposal_pol_round": "-1",
            "proposal_pol": "__________________________________________________________________________________",
            "prevotes": "__________________________________________________________________________________",
            "precommits": "__________________________________________________________________________________",
            "last_commit_round": "1",
            "last_commit": "xxxxxxxx_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
            "catchup_commit_round": "-1",
            "catchup_commit": "__________________________________________________________________________________"
          },
          "stats": {
            "last_vote_height": "870166",
            "votes": "29631",
            "last_block_part_height": "870166",
            "block_parts": "29030"
          }
        }
      },
      {
        "node_address": "292c1eb49a6e3d88c48adb8e42a9c1e61488cb78@54.215.221.213:26656",
        "peer_state": {
          "round_state": {
            "height": "870167",
            "round": "0",
            "step": 1,
            "start_time": "2018-10-26T03:52:46.992025408Z",
            "proposal": false,
            "proposal_block_parts_header": {
              "total": "0",
              "hash": ""
            },
            "proposal_block_parts": null,
            "proposal_pol_round": "-1",
            "proposal_pol": "__________________________________________________________________________________",
            "prevotes": "__________________________________________________________________________________",
            "precommits": "__________________________________________________________________________________",
            "last_commit_round": "1",
            "last_commit": "xxxxxxxx_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
            "catchup_commit_round": "-1",
            "catchup_commit": "__________________________________________________________________________________"
          },
          "stats": {
            "last_vote_height": "870166",
            "votes": "33917",
            "last_block_part_height": "870166",
            "block_parts": "33444"
          }
        }
      },
      {
        "node_address": "e81a57e3275340174d007636d78a28a238cbba46@192.241.154.136:26656",
        "peer_state": {
          "round_state": {
            "height": "870167",
            "round": "0",
            "step": 1,
            "start_time": "2018-10-26T03:52:46.8731486Z",
            "proposal": false,
            "proposal_block_parts_header": {
              "total": "0",
              "hash": ""
            },
            "proposal_block_parts": null,
            "proposal_pol_round": "-1",
            "proposal_pol": "__________________________________________________________________________________",
            "prevotes": "__________________________________________________________________________________",
            "precommits": "__________________________________________________________________________________",
            "last_commit_round": "1",
            "last_commit": "xxxxxxxx_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
            "catchup_commit_round": "-1",
            "catchup_commit": "__________________________________________________________________________________"
          },
          "stats": {
            "last_vote_height": "870166",
            "votes": "33926",
            "last_block_part_height": "870166",
            "block_parts": "32468"
          }
        }
      },
      {
        "node_address": "303481d4f855f227fb576c765ed88a15b67150b9@188.1.89.170:26656",
        "peer_state": {
          "round_state": {
            "height": "870167",
            "round": "0",
            "step": 1,
            "start_time": "2018-10-26T03:52:46.81040707Z",
            "proposal": false,
            "proposal_block_parts_header": {
              "total": "0",
              "hash": ""
            },
            "proposal_block_parts": null,
            "proposal_pol_round": "-1",
            "proposal_pol": "__________________________________________________________________________________",
            "prevotes": "__________________________________________________________________________________",
            "precommits": "__________________________________________________________________________________",
            "last_commit_round": "1",
            "last_commit": "xxxxxxxx_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
            "catchup_commit_round": "-1",
            "catchup_commit": "__________________________________________________________________________________"
          },
          "stats": {
            "last_vote_height": "870166",
            "votes": "33890",
            "last_block_part_height": "870166",
            "block_parts": "33104"
          }
        }
      },
      {
        "node_address": "e179345924be168def105e64daff369c27e7c47e@104.237.6.136:26656",
        "peer_state": {
          "round_state": {
            "height": "870167",
            "round": "0",
            "step": 3,
            "start_time": "2018-10-26T03:52:43.697998611Z",
            "proposal": false,
            "proposal_block_parts_header": {
              "total": "0",
              "hash": ""
            },
            "proposal_block_parts": null,
            "proposal_pol_round": "-1",
            "proposal_pol": "__________________________________________________________________________________",
            "prevotes": "__________________________________________________________________________________",
            "precommits": "__________________________________________________________________________________",
            "last_commit_round": "1",
            "last_commit": "xxxxxxxx_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
            "catchup_commit_round": "-1",
            "catchup_commit": "__________________________________________________________________________________"
          },
          "stats": {
            "last_vote_height": "870166",
            "votes": "33490",
            "last_block_part_height": "870166",
            "block_parts": "32559"
          }
        }
      },
      {
        "node_address": "1d20090454966d7c5b74f9442b1d39e9e8511ca9@91.90.148.194:26656",
        "peer_state": {
          "round_state": {
            "height": "870167",
            "round": "0",
            "step": 1,
            "start_time": "2018-10-26T03:52:46.883909771Z",
            "proposal": false,
            "proposal_block_parts_header": {
              "total": "0",
              "hash": ""
            },
            "proposal_block_parts": null,
            "proposal_pol_round": "-1",
            "proposal_pol": "__________________________________________________________________________________",
            "prevotes": "__________________________________________________________________________________",
            "precommits": "__________________________________________________________________________________",
            "last_commit_round": "1",
            "last_commit": "xxxxxxxx_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
            "catchup_commit_round": "-1",
            "catchup_commit": "__________________________________________________________________________________"
          },
          "stats": {
            "last_vote_height": "870166",
            "votes": "28817",
            "last_block_part_height": "870166",
            "block_parts": "28716"
          }
        }
      }
    ]
  }
}`,
	"genesis.json": `
{
  "jsonrpc": "2.0",
  "id": "",
  "result": {
    "genesis": {
      "genesis_time": "2018-08-21T08:00:00.123456789Z",
      "chain_id": "gaia-8001",
      "consensus_params": {
        "block_size_params": {
          "max_bytes": "22020096",
          "max_txs": "10000",
          "max_gas": "-1"
        },
        "tx_size_params": {
          "max_bytes": "10240",
          "max_gas": "-1"
        },
        "block_gossip_params": {
          "block_part_size_bytes": "65536"
        },
        "evidence_params": {
          "max_age": "100000"
        }
      },
      "validators": [
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "WnV5CjfHHyMrAPbX796iXgnfAwTP/cFL+kzz99bW+kA="
          },
          "power": "7",
          "name": "baryon.network"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "RBuRRNEzA9RA1Wrdi9PPFQJ29/n/bqN9O2tQv9Gq248="
          },
          "power": "2164",
          "name": "skoed-validator-7000"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "F2wCn9rKafNZsYZwoLGkSQIpr3rk86cjYyuhSjsjRaE="
          },
          "power": "255",
          "name": "ironfork"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "1XB37ELQL8SIu5CeeRNMByNFtf2rQf543PKoxUCBeBA="
          },
          "power": "54",
          "name": "sentinel.co"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "NQX4yKpOztKrmgBhGIC5WOALOLOq3LTpbzsN4ZLXGec="
          },
          "power": "990",
          "name": "Dokia-Interstellar-Cruiser"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "j9be+ddLChInrFz6/820/uYh4WZBzlp61klyJBDy/ZY="
          },
          "power": "100",
          "name": "redbricks7000"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "8pfpbIxBBiu88hpxS3CeRpv7kClEjl8SwVgckDNBGlE="
          },
          "power": "850",
          "name": "Staking Facilities Validator"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "pPGLc4NhNaehdoV2antWuyr0GmBVEG1NhD9NiSRrTi0="
          },
          "power": "1250",
          "name": "sikka.tech"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "Go9GXHI6SCQo2QKMxkAkgYLhfo3XrVjWLR2nE2AvYyk="
          },
          "power": "100",
          "name": "Staked"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "KTcVPCcGHjfCtHJF/8+Z6WnP8aF1KDZGP/17zjkQyok="
          },
          "power": "150",
          "name": "debt"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "b52QQPpAO9Wa5qEmOr8ZPvLho0CgbBJx+WWxK+bpaYI="
          },
          "power": "2265",
          "name": "wxvalidator0"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "QzxYamaSVNx+Idlb9bCcYsPI4Oa/zQ+rx4MN+XA2xLI="
          },
          "power": "1065",
          "name": "Bison Trails"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "ydjx2ea+PVuChrny6X2dluJwyXta+BsNQRsgHXp8fXw="
          },
          "power": "150",
          "name": "Mia"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "FSj9jMpCOdWfX0rCQ0TpnEa8/uOroQl4V4Bk3Msck3w="
          },
          "power": "10",
          "name": "trusted-validator"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "M2Iu/A8RQcmoiIEar/iV4wBjVj0I0gcd/nNnvwmnT2g="
          },
          "power": "1000",
          "name": "layover run"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "w0Xbnzcm1tLGgpUrICpotJJXKDlmctoPz/JTBjzkIfk="
          },
          "power": "10",
          "name": "ATEAM-TT"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "eeImG09hOPo1W7j7lKepN/Lx6I9GGHqVBVEKmznxACc="
          },
          "power": "1220",
          "name": "⚛️ melea-trust 🐞"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "A6GzeXUM3vsXaDAEYMSDgSKkqn9AoUYjs8empH46MGY="
          },
          "power": "1280",
          "name": "Mythos"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "Abws3eXrUFAH8LeZJIcECakPL945TTmFsBlXONOUeII="
          },
          "power": "155",
          "name": "Cosmodator-7000"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "qgws5Mm+saA71C4cdQavr1bw0UytNqJcpwPDHR0k2hk="
          },
          "power": "9",
          "name": "EddyG-cVali"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "dpYEHcQAo439YIswYnm3ytxyoP7RtorAS59+FOXiOKQ="
          },
          "power": "230",
          "name": "new_vali"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "UQHY8OYNqCwDkAm7bn+X/VqV6p3PHTRqkQGEsqnpdws="
          },
          "power": "110",
          "name": "kreios-2"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "lUMCRAeu47BsOhNvCQTQJQeB68z0/VaElC9j5gDt9y8="
          },
          "power": "1501",
          "name": "vhxnode1"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "ilDFePqWkNRNWUEoWKq42jZC5CD4rglfs7+gCuq6Qc8="
          },
          "power": "151",
          "name": "lino"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "xlO2cnii42KisAn8OcstC/3XV5+I0FlcSbWuyy5MVA8="
          },
          "power": "127",
          "name": "21e800"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "XQDVMXja3kFk5Jb47BsqJmzcDsM4lE9+r+f/J3O5Jms="
          },
          "power": "135",
          "name": "inschain_validator"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "FiEl7a96NaaDaVH5mV9EK/6tOUWL5mkkT46ARRKAWIg="
          },
          "power": "1681",
          "name": "chorusone_validator"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "LRK6flgMsHrBYkIIahegEkzlwC7gQ3dAyw0jcAd1u/k="
          },
          "power": "505",
          "name": "@fulltrust"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "MfyCioP9/yGlqUcwnL/ENi23RKPBm1O1wum1DUUUYT8="
          },
          "power": "56",
          "name": "Bity Validator"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "ISIM341M+EUYOMpwhn9gyCvRUZ06xomp/+lOwa50meQ="
          },
          "power": "800",
          "name": "corona02"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "DrgnMZqdtzRbDvDGh5290M1QtAfem1VHaY+zhq/nBWI="
          },
          "power": "50",
          "name": "FennecFox"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "VSz12Gh6kqPKSlwrZZQ0QZsXvRadOvAeDhmuboV/Lqc="
          },
          "power": "281",
          "name": "412"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "VakMQSPBEuSC9Nwuv8WWhrZVUmH31bUR4+G6pJhkgE8="
          },
          "power": "1773",
          "name": "Umbrella"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "lB+CgSbTzpHXJHUU15fJIPaPl6XF0nSfdNxBnlKJqTk="
          },
          "power": "1829",
          "name": "delegate-now"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "ViJzrw9Unk1XyDg4f0UQLFE4zfUaFh23wZZY0f31Qg0="
          },
          "power": "1200",
          "name": "CosCloud"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "RwPRoiY5C0covekqbr3VrQwxWGHioUUIf2+TOq8LIC0="
          },
          "power": "135",
          "name": "starfish"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "VBVHOLnWGptY26J0wqXoZI2Dnu96pccMb08zlsaxPCQ="
          },
          "power": "1100",
          "name": "certus.one"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "5GAxIHzu06l0n3MU8wNQrCcrrnpZR9EKe5qZMsM2h/E="
          },
          "power": "100",
          "name": "ip-10-0-3-99"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "kol7Gj60Fct4X8T1rHLJQ0z/b14UqqSae8h1e37rLL8="
          },
          "power": "90",
          "name": "VNode01"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "N3K5kDdfcKJurfaa6s2zfKgtYvz1Pagz7VWi9ZfX8yM="
          },
          "power": "1252",
          "name": "ATEAM1"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "4JoJuRfaANhdM1x3AWRo1/Cj9DH3VA+fi1SynzknV+w="
          },
          "power": "150",
          "name": "davinchcode"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "y7p9JSVZBnRxjAI9v5Pxl37hMtyuHf6B4Ghqzm6+ii0="
          },
          "power": "100",
          "name": "oleary-labs"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "R/3f7VruxWpu+2hiHlVpplTwoOou5kfQI1k/6/9H/y8="
          },
          "power": "152",
          "name": "stake.zone"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "aSbdZXtmYI31dCcZ/OgyL1O6YltrDqmkSyuhLIiIKxQ="
          },
          "power": "1998",
          "name": "codetree"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "dnFjFoTM9sP/RjQkXBK1YpYn3v5W+j0+g/OfUHS4xu8="
          },
          "power": "55",
          "name": "aether"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "omAzuJps8KX3/iOC1LjwkMPMH3c6tjfLXwCNWXRBdWw="
          },
          "power": "1455",
          "name": "Adrian Brink - Cryptium Labs"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "BaaCxmYHKJ6obIzTCdRtjw1cc8d2mUJcMbLWCjf1aLo="
          },
          "power": "3359",
          "name": "grass-fed"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "31942iQRr+ajHMCklJslIbPzQl6OplvluWuXUnHlVQ0="
          },
          "power": "606",
          "name": "newaether"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "gJDxhwAE6GeGCKQeVaNZ5is7+7MFHXtOG0UsnguKdoA="
          },
          "power": "300",
          "name": "BlissDynamics"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "GpHZ0Hr3C1hvVMkcDe/76lvWMjw5SF0JBI/D6gDkQTo="
          },
          "power": "250",
          "name": "StakeStack"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "Epd2FDKZwDybzT38Z7WB0y2jiLn9/2OLzmY3Zu18l6I="
          },
          "power": "1257",
          "name": "figment"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "t07LChLn+eoTivIraUyU0M/T4zkhY6TZkn8cAWtOLfU="
          },
          "power": "18",
          "name": "egon validator"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "tkdOxR2CIOctyc5e9IMGwnMwB420OTN3Rfb6fUMVhr8="
          },
          "power": "5",
          "name": "Mythos(private)"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "dj4b9d19oYvNJP1WMHJ1g5tzzMGMNxRnxrtCXBF6jsc="
          },
          "power": "30",
          "name": "Liviu"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "1UyFY3YJja8vrmDc/ZqQFo7ORvLdLK2X2wYeCUBZzJc="
          },
          "power": "204",
          "name": "clawmvp.xyz"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "zaRqlLb7T6OaFRImEZF1eIuzoRpDpVyvBs3ITbPuPiw="
          },
          "power": "30",
          "name": "satoshi-ledger"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "UjTvuOew2EaooduJBiYmBWeF5ai0yFJG8uio5YXpJgg="
          },
          "power": "135",
          "name": "P2P.ORG Validator"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "AJR2ex094A1nJEWQsZYjALWsrSl1/huGZ37z2ZsMrpg="
          },
          "power": "1100",
          "name": "dev"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "Oxp5JKfzexRWJravPZ9XAEF2ydyzh04362Eo4WloEBE="
          },
          "power": "31",
          "name": "Yakut"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "KfZeww1iQ6YnslfhtZ0qi8yVy0PQHj7/j0Rrkk2gz5o="
          },
          "power": "110",
          "name": "bharvest-validator(USW1)"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "k3YLQYEN2QMP6XITRsBmgb+pNGhJ5Jbg0bzUW977kK0="
          },
          "power": "90",
          "name": "TropicalMongoX"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "2JoNf1gavJ1d6XFIumO1Mki5GVMOcg58AioHksU3maE="
          },
          "power": "1130",
          "name": "bianjie"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "6lTjnB9szF+z4c0XwBrplBjK6GaaMUAcFJ/8VHwzOuE="
          },
          "power": "10",
          "name": "ATEAM-H"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "Kv+wY9i0uMGO8ukF4o+Za5uuUp/S8o9N4C1rdA22h8M="
          },
          "power": "10",
          "name": "louis"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "UZlobYgzLP7gOVwBtMRo5S3yrMPlZWIXYEwXCs7dhIA="
          },
          "power": "359",
          "name": "kreios"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "+Uc1tDCsc8rM+mPQ79C8Mp3Y7ajlL2Lb5AjYboPqlQE="
          },
          "power": "5",
          "name": "chungkueiblock"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "kCfC+M4tOz3Js2IFJw331BkWcW+cdpSJXFRkSakmFZU="
          },
          "power": "40",
          "name": "INVIERNO"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "QEMDwUVoyJT7MNfOYKa25xU+Lnsz/ciH8rFUri4diLI="
          },
          "power": "90",
          "name": "SVNode01"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "LSDd6ab46sHxwJSrg5YLpsPG2o6EcsZ3rDikpHzMNmI="
          },
          "power": "210",
          "name": "w1m3l"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "0HqB2x6x5HzeozpHatePECw07x1UcDdSz8kQGNznnA8="
          },
          "power": "1001",
          "name": "block3.community"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "UBTju7UZfXLVPPYb1a8gPZ69BeCv2Fho7YVo2EUbxKc="
          },
          "power": "135",
          "name": "7768"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "jj0Y/Fy8JSJR3g+PHU6Ce0ecYwHGUVJ4bVyR7WwcyLI="
          },
          "power": "1151",
          "name": "nuevax"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "oG8Q5o+SN4wqMLvlIfVgQPnsQzNEKeH0D/XGM8JlGrY="
          },
          "power": "2512",
          "name": "wancloud"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "Xi7nIgj4PqVXrpKLfJhcyxyVY1d3HRo72sKKPDmuU78="
          },
          "power": "150",
          "name": "kittyfish"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "G1F5j4VjmFrlOBHtqqOAa7q2hMfUm6lnzctnVFhROZ0="
          },
          "power": "150",
          "name": "⚛️ cosmos-trust.com 🐞 🎬  🌈 🦄 🌟 🚀 "
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "3wRufybSUsTMnUeQkP74uJNDRKeM8jBLAS64T0BRfpY="
          },
          "power": "180",
          "name": "mpaxeNode"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "MRCeBDANSjH6IsxO0z6tRe+xqoZvIGhdfl1t+SXGUpM="
          },
          "power": "1160",
          "name": "jlandrews"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "b9RSkt+WmMxVHQExQH0IMPpnR9zDAaJwz/mv1gtyRVY="
          },
          "power": "81",
          "name": "smartpesa"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "2qtEBT+Tc+SD2wJsdrVMHXrBKfvesxtmtSKDK5fXwA0="
          },
          "power": "144",
          "name": "jjangg96"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "BB4a/Xh5z+dkGCRlF+pSGC3iDOoDrFse/xzQAtmxMF4="
          },
          "power": "1149",
          "name": "chainflow08"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "66j9af4xDJSblMLS+mFbp7d8TaFGu0FOo+0MwEYm2lE="
          },
          "power": "1210",
          "name": "bkcm"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "G5xqaI1M/ckd6k5tLzWkGF/MdGZ4/xmJThP4AjksrFk="
          },
          "power": "29",
          "name": "basement"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "CuPin/hLM2tK6KvcUlOi/Brm5xNi3zq6UPgpLFPSzt8="
          },
          "power": "28",
          "name": "NodeBreaker"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "ePxcUeuPx3UiqfCPLioJjYwDgk/E7HGhatxfzl4/vI0="
          },
          "power": "1340",
          "name": "shensi-cosmos-validator"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "aUViYC2znC55sleHfmsIN9hZ45SbYPbDcYA0gVzglsc="
          },
          "power": "1100",
          "name": "Gold2"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "wwgoAZrmpxQRYXHDXRUBX4kX7IBTpDsvuyxFfKuhEZI="
          },
          "power": "20",
          "name": "fourthstate"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "cCoFsZzKZ9SQZbHe4NueVObIezP6ts0tRTZ/aN96dig="
          },
          "power": "150",
          "name": "Dori"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "HjSC7VkhKih6xMhudlqfaFE8ZZnP8RKJPv4iqR7RhcE="
          },
          "power": "392",
          "name": "firstblock.io"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "IS4skJndcdUJqSxMPkMmOD4rGaySvA8JJZk/TBxN7dQ="
          },
          "power": "32",
          "name": "hallophen"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "TwzOJ4GcN+ZTswub4R8488SrKeWXjY/PaqCF5neXJig="
          },
          "power": "400",
          "name": "kochacolaj"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "S8s6fdAQNQ3bN9SNVAsHB/j8uv1CM1roxeLesL+fh4g="
          },
          "power": "129",
          "name": "luigi001"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "8K3clCjVU33BTIpUhdahGmu++WxHj4NUE9krCRkk++s="
          },
          "power": "1008",
          "name": "lunamint"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "K4kLogLtZxqrYSqRVJfrFm9tUG+Tc3QWXWIewnAgI9w="
          },
          "power": "147",
          "name": "abcin"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "pggq0WUTlsiA8cBiptRgxlw4WUG2GmXfLQxkbMbCBPk="
          },
          "power": "155",
          "name": "wecosmos"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "0ZPHmftStU+DSqvnIleNyaFmBQsgWJwPC7vk4uF5u00="
          },
          "power": "2456",
          "name": "sentinel.co"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "ENAVynNXVpj/IdYx9kCPKaPs4bWSxRIHNlmS9QiDuZQ="
          },
          "power": "1155",
          "name": "iqlusion"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "5XCMLGNTGkMh9/JsUzi/8Xl7Y0cdDe38stX+zvBLXuE="
          },
          "power": "974",
          "name": "bharvest-validator(KR1)"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "JREG9eji9gfJR2b61qMu8jNZ835ezPY2A4o+zjYMfDg="
          },
          "power": "107",
          "name": "kamon-moniker"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "PxJbo5FKA6mXtgwclRQVNIjOCQK3Q7WkLQrvM9lYbGI="
          },
          "power": "156",
          "name": "space4"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "oL/QCr7LEOivyTqpGrmwVd1r+hYI2WB5+kSVzpDMxx4="
          },
          "power": "700",
          "name": "gruberx"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "H0SIA/BU6Xj8oT5bQkvLpEITN3CqFLbMeBcQ72NZrAE="
          },
          "power": "135",
          "name": "4455"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "WRsXnLz3gf8o4lYYeCZjAXgPU1cdmOOYPdy7aY63iIA="
          },
          "power": "100",
          "name": "ColmenaCowork_SVQ"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "P9RgE4RMQT/aHap2oICpwpgKeBAwxPUwuU9zIffKFNM="
          },
          "power": "1231",
          "name": "Nylira Validator"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "4DEDoU/RsHMPS54GzgwkWnW5zPQfMt9aInFFc3GyfA8="
          },
          "power": "506",
          "name": "crytter"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "DsTbM0AgHfhSUKvOGkxudDOY3ojYT6bifhpelqHs8+s="
          },
          "power": "100",
          "name": "BFF-Validator-7000"
        },
        {
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "O95l52FybUIUBkaYeGHCLRNkSI5zAewb4SxzS0LOb50="
          },
          "power": "115",
          "name": "loafer"
        }
      ],
      "app_hash": "",
      "app_state": {
        "accounts": [
          {
            "address": "cosmosaccaddr1qz9pauujxsypyu3q9n0t6mydsfwev50j8g8ras",
            "coins": [
              {
                "denom": "steak",
                "amount": "11"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1qrc3ed8tnz6vc24ftmnht8efs5ufhjmrjkds4x",
            "coins": [
              {
                "denom": "@MarceldeveloperToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "251"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1qtkc837fpfprvr2fcmuw6hgkesen4pxnslkr4z",
            "coins": [
              {
                "denom": "steak",
                "amount": "30"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1qv68kxwmarzefrcgnu4trdrznslzdtra7gfxhy",
            "coins": [
              {
                "denom": "steak",
                "amount": "3"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1qvut6w8v5fccephuv502qhmvhn7hy6cnddrdyf",
            "coins": [
              {
                "denom": "steak",
                "amount": "4"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1qw4eww34ug66edg9mgsapgcgjuqcpyqxv77f7f",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1qsvdcgf7zyue0usl2tt3ptn0x6lzs0ua6y5dgs",
            "coins": [
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1q3lcu430n2mjar6q52s2y7z2hjt7j5dy7t7w2d",
            "coins": [
              {
                "denom": "steak",
                "amount": "20"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1qjzjfn55hygaak9l9x04z792mexce2zd2gvkzk",
            "coins": [
              {
                "denom": "mycoin5Token",
                "amount": "100"
              },
              {
                "denom": "steak",
                "amount": "72"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1qkpevcffatqksc7zw28mzrqd0hwkus5qrlsmpt",
            "coins": [
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1qk2q3l28t6qhzyc7p832ep8uwd0ypstgqla8yx",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1qkc3tghc3fms7eye7vtu0g0370epr4jkje2ne7",
            "coins": [
              {
                "denom": "liangpingToken",
                "amount": "10"
              },
              {
                "denom": "skoed-validator-7000Token",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "5000"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1qkllv8f6qakkw3hk9dqvytys090lk6twsyv8vf",
            "coins": [
              {
                "denom": "ironforkToken",
                "amount": "1000"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1qexep37ryqkj6rvaah5us24tve46hac755vru3",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1qe82wd64ncuh8z05kmwpqhc5vkq4wsxqv67cvz",
            "coins": [
              {
                "denom": "steak",
                "amount": "4"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1q6uza2sy8wux9ws33dtfax6j3vn307f7lrv7t5",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1qmw6d8uh3rs9uj8c0tztznu9x4zezekq89635d",
            "coins": [
              {
                "denom": "steak",
                "amount": "5"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1qarnnunjrsd66yuz65ugsl7pm3v3hxrxd5jvaq",
            "coins": [
              {
                "denom": "steak",
                "amount": "385"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1pp58pzp26zqf4sscwdzpns9c5prpkcr7j2v0nx",
            "coins": [
              {
                "denom": "steak",
                "amount": "19"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1pzlud6lg8w9phcwetc5aqp24eflshtv4xlxthf",
            "coins": [
              {
                "denom": "coscloudToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "151"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1pynxhek6ls6xcwptuwt6cz94lezz6wu6j062gq",
            "coins": [
              {
                "denom": "steak",
                "amount": "30"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1p9wdfa8kt2khg2kgj74l7zwcveg9nkp82le3g8",
            "coins": [
              {
                "denom": "steak",
                "amount": "20"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1p054wqkus4rjajwh3e079xuk4g7l4zm5t74kce",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1pj6vfgetplkcuq0k9m98dpej6zew3pdest842v",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1pngeljxmaecunw2uw246vejvg8pezrpc7ltg4m",
            "coins": [
              {
                "denom": "steak",
                "amount": "56"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1p56gv748xfd74qek5e637vhcr6tyjd9ukqfecc",
            "coins": [
              {
                "denom": "dokia-capitalToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "35"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1ph64lhnqwp4tjf4y70y45u5e7cnv2djfg9nkgr",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1pc63alzz2qdgxk7f3mhr56xly8xz5cs47k294m",
            "coins": [
              {
                "denom": "steak",
                "amount": "4"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1pm0gx3lk46gf8vyj5my9w5tk06cuq66ll77ugj",
            "coins": [
              {
                "denom": "redbricks7000Token",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "150"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1pmntq5en2rgtr5rzr4e304efrve4lr43z32y5s",
            "coins": [
              {
                "denom": "Staking Facilities ValidatorToken",
                "amount": "1000"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1pag7xgpl8njlaksvp2ur5un3htg85vcrxcp5rs",
            "coins": [
              {
                "denom": "ravenclubToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "150"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1p7jg2t7xrjmxpu94tp9wawmnx64ds8skfmsqe0",
            "coins": [
              {
                "denom": "steak",
                "amount": "15"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1p75wzde5z2ataefw65j3mtrm25kha779yfrvt6",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1zpx36km7tk5cksyzxgvcp6g552p3uhwx84r53q",
            "coins": [
              {
                "denom": "sikka.techToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "200"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1zpdmfsq9p32sxgqmxfvvczndxg0s7l506a6n0x",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1z8mltyca36vv5fud4xgtntjzekn20tz5auut72",
            "coins": [
              {
                "denom": "steak",
                "amount": "2"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1zvyjt8aw6xgqfxmh7mguvntudrz4llln2dyaam",
            "coins": [
              {
                "denom": "steak",
                "amount": "19"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1zv9ayyk8l5ptzuuur8wt5up73pptzm2sv4spj5",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1zds6r7jyxyxcpf05r5yyyy3u8q2rvj9kkc6vcv",
            "coins": [
              {
                "denom": "StakedToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "60"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1zn5hc3rtpd4twuwaemj550gx5hmh9vvez2jmla",
            "coins": [
              {
                "denom": "steak",
                "amount": "20"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1z5ujex3hy53c3kwlaktgxwqf00lraud6qna5fy",
            "coins": [
              {
                "denom": "steak",
                "amount": "2"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1z49r5y3yk6kpz6wy64cr5tu0lung2cgg3j4tn5",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1zkupr83hrzkn3up5elktzcq3tuft8nxseu9x80",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1zagese3pa7k002epyy8cpt783pe2p2n5qprug6",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1rxxcpkmsngd0azkh3n2467m66ls4rwq52yuv27",
            "coins": [
              {
                "denom": "liangpingToken",
                "amount": "990"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1rg7udx3zvn3ux8l2j0l7ysyskum79972fgndpz",
            "coins": [
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1r2hjwvyudlmxlr3dqdq6r72khcc7yej9zew48d",
            "coins": [
              {
                "denom": "steak",
                "amount": "67"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1rvjcqjnxmsq3nlhv4cwjnhfjag508qkh6r62tr",
            "coins": [
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1rvm0em6w3qkzcwnzf9hkqvksujl895dfww4ecn",
            "coins": [
              {
                "denom": "steak",
                "amount": "405"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1rww4u22qqf2rhw02fww63t539d3ppeur3ywff3",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1rwmtl5nqz9frczcp8pgs95vak2mqka4546j6p4",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1r05wuvemjzl25n3lspfku03e069cs5xsulfusn",
            "coins": [
              {
                "denom": "steak",
                "amount": "35"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1rj74vaqm0xkxl5cjjam63mayh4x6at3m379ulv",
            "coins": [
              {
                "denom": "MiaToken",
                "amount": "1000"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1rnjapk7nw3ts6h6lsse3vm9d8mjjqqrx9mvc09",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1r5epqszx4supvesf4csxxlaxn4947lc38cjkcw",
            "coins": [
              {
                "denom": "steak",
                "amount": "30"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1r4lzzzf683ec70fxfd9a3jmczxxkdxpfafz4ar",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1rcp29q3hpd246n6qak7jluqep4v006cdj2zse3",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1resera5tks46p0gshs8svvjk34fz9vhsmhpztg",
            "coins": [
              {
                "denom": "steak",
                "amount": "20"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1yqq6up8zyakmf6t7gg5txnyatqakj7ckkt8jfu",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1yqpa5vhs8nq9mguyd2j5k5kn7dlqlgmjl4h67x",
            "coins": [
              {
                "denom": "steak",
                "amount": "6"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1ypk9spe0lsta44f3ay6rqrzygh90ekayxcxcfn",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1yz2t5qxggyq0ffkfhw6v59xnaqackkw0009wjy",
            "coins": [
              {
                "denom": "steak",
                "amount": "70"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1yfd9a4sqm8jm7lq932jc6fu8k86f394tgc2sku",
            "coins": [
              {
                "denom": "steak",
                "amount": "370"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1yfseqtj5sjhzz2q2ym09jym4h4nc4yevae0jp2",
            "coins": [
              {
                "denom": "meleatrustToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1y2z20pwqu5qpclque3pqkguruvheum2djtzjw3",
            "coins": [
              {
                "denom": "pbostrom/Mythos-1Token",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1y2jfcttj9w9l4ztdaea9059e7y3trzc7w0cp4u",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1yvawtylze64y5t6x7c9e6nzf6num9s0hd762ze",
            "coins": [
              {
                "denom": "steak",
                "amount": "2"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1yd0rklq45zg89fywr89ccutlcwp9kehhh0z03k",
            "coins": [
              {
                "denom": "Cosmodator-7000Token",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "2131"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1y0qavya7zxwrn0qe43f6nxmm48kaavgpvuxge6",
            "coins": [
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1y3pmvvwjqrtxa3hatw33ylp8hz5ge0cw4rmt0u",
            "coins": [
              {
                "denom": "cwgoesToken",
                "amount": "101"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1y3fydat73q569cceaccpuytm3y4tdpv3n0fhdv",
            "coins": [
              {
                "denom": "steak",
                "amount": "80"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1y32vvanvz88ajamugxrgl7vnvf0vtmjg5wxg7f",
            "coins": [
              {
                "denom": "steak",
                "amount": "60"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1yj87w5yjftyar3mhuefg569q4v26rfp4ntacc5",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1ychw8dj7r37j9mlpd4j22azqzwzdryqkejf0ac",
            "coins": [
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1ye7jc0hldu3k7nz42zcg2yvy7lzttx9q5yf3w9",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1y6d5g4s75s9m8xemrs6d9am2cef4nm3ku0s5lv",
            "coins": [
              {
                "denom": "steak",
                "amount": "5"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1y63usu5k67zr37ha5p30pgzawv4g74qtmvwsyp",
            "coins": [
              {
                "denom": "steak",
                "amount": "20"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1ymwytdwfu39uxpxmkjenfnyv3mkskehf97hwu8",
            "coins": [
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1yu7a5agy0ypchjcj467ej6jys468makfsvgg9e",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1yaxhcrpu9a009aavd86rrewlsmxcxh6lycazky",
            "coins": [
              {
                "denom": "steak",
                "amount": "2"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1yanwmhh6hxldue3szsm36k74n0xxrz3u8wllex",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1y7n3emnw6mkxutvs26rczynpvq84ddqu332hlv",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1y75427njr6vrfhd5tljw4gpffezgmmwx3slz2m",
            "coins": [
              {
                "denom": "steak",
                "amount": "11"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1ylz6tqpgz0flpl6h5szmj6mmzyr598rq7heuye",
            "coins": null
          },
          {
            "address": "cosmosaccaddr19q6y4mk3x668vm2jx0pu6e0jfy5a39jz0qwvxt",
            "coins": [
              {
                "denom": "2400bpsToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "40"
              }
            ]
          },
          {
            "address": "cosmosaccaddr19pkt3aeaprl8dt9x5c77u9unkrndfd3ze40tj9",
            "coins": null
          },
          {
            "address": "cosmosaccaddr19rlhdmplhnmvy3jmcd835awask4l46nvp3pt4g",
            "coins": [
              {
                "denom": "steak",
                "amount": "40"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1998p0xkdwvul952ydavnx88tmwkhlfa0vhrngj",
            "coins": [
              {
                "denom": "ianstreamToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "50"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1990ceh7wh50j2z6wq08ekzez7yr6hpwtpap9ku",
            "coins": [
              {
                "denom": "steak",
                "amount": "2"
              }
            ]
          },
          {
            "address": "cosmosaccaddr19grl7vglzuduusjcnatly5wgrg0v5vqhs356f3",
            "coins": [
              {
                "denom": "irisToken",
                "amount": "40"
              }
            ]
          },
          {
            "address": "cosmosaccaddr192yhf7f7f9304cy6cu6np3r8a3m9yqzqfeu9je",
            "coins": [
              {
                "denom": "broadleaf7000Token",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr19d8m2wg4glf4rfaaz96pdxxhq8tx4e98zrhuqu",
            "coins": [
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr19dfkpwfj3mls96ny2ulgq7cf0dh5xs5a05pn6e",
            "coins": [
              {
                "denom": "steak",
                "amount": "2"
              }
            ]
          },
          {
            "address": "cosmosaccaddr19dw2kndn6shyx968zc3k4ah5wa05z0zhy8gsh8",
            "coins": [
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr193vn0gk3nsmjkxwz78gce8e8mkmagmvulpg5jt",
            "coins": [
              {
                "denom": "jackToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "50"
              }
            ]
          },
          {
            "address": "cosmosaccaddr193e492vdft33xu3td0wmkc9kswemcyysvkz8x7",
            "coins": null
          },
          {
            "address": "cosmosaccaddr19nqkn77ya8fxcjgmcdzepecypvx270avjfyjy8",
            "coins": null
          },
          {
            "address": "cosmosaccaddr19n4s8fx9hr0j2dttr4fg48wr7kpxxste7evjzk",
            "coins": [
              {
                "denom": "steak",
                "amount": "240"
              }
            ]
          },
          {
            "address": "cosmosaccaddr19ncpp07tuqqgu2kjp4wq2zlt3s0q5f9ywvhwuw",
            "coins": null
          },
          {
            "address": "cosmosaccaddr195qlnqc2mxsxnfvcljqr5ssfjjymn0usk35v75",
            "coins": [
              {
                "denom": "steak",
                "amount": "2"
              }
            ]
          },
          {
            "address": "cosmosaccaddr194pux2wkhuusey6x4wwaqm0wh26cleh06u28lr",
            "coins": [
              {
                "denom": "steak",
                "amount": "5"
              }
            ]
          },
          {
            "address": "cosmosaccaddr19h5dc2c76h4ryvl9zl9373mk52ka77vy5g48gs",
            "coins": [
              {
                "denom": "steak",
                "amount": "19"
              }
            ]
          },
          {
            "address": "cosmosaccaddr19cemj59p7mvmnfukya2a328dsh6nd7yg3hdzt0",
            "coins": [
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr19uhnhct0p45ek6qxp3cjjrjtz4pacwcsgzvpuj",
            "coins": [
              {
                "denom": "steak",
                "amount": "450"
              },
              {
                "denom": "vhxnode1Token",
                "amount": "1000"
              }
            ]
          },
          {
            "address": "cosmosaccaddr197p65z7ge5g55r68thvw4l5e43gnm70jhu5g75",
            "coins": [
              {
                "denom": "ritter-rammToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "50"
              }
            ]
          },
          {
            "address": "cosmosaccaddr19ljp6mlsfuy52r7d4c7r8mqfm0cxw5yktkqt80",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr19l4ytmyrf0234r20n0u3av830adsljw2jck7jl",
            "coins": [
              {
                "denom": "steak",
                "amount": "4"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1xzv7cxt0d99hg8c67dg7g7hketsnjpd40hmqpf",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1x9zs84mzzcdncyxh84k99sn3qhqw5ry8r8zwy7",
            "coins": [
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1xxln0menrvzkjdjtqaqf7juya2dtchezqj9fc7",
            "coins": [
              {
                "denom": "steak",
                "amount": "20"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1xf56zlpjs9f5w3ygy4mq907ea60pqw6z3xk0x3",
            "coins": [
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1xtkfzzty24l66a0f9guqt4w8k0p66fpat7rlf2",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1xvzw5nz4mp5s5cat7rancc248v5wpecrlqtw6n",
            "coins": [
              {
                "denom": "steak",
                "amount": "2"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1xvt4e7xd0j9dwv2w83g50tpcltsl90h5dfnz6h",
            "coins": [
              {
                "denom": "21e800Token",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "48"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1xdp4pls4ryvchq2n8v0cpmtwsprvyh8wvg563q",
            "coins": [
              {
                "denom": "sheiudToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "50"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1xs8gy05pgxtsddx42az9lnrpu3ah0sg09krawq",
            "coins": [
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1x3k0xzhrrxk95el2f8a8w4pcn03lq5mzfgjmvc",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1xhzp9marf3zjtjn4hl73xkv4stgsjhj2xyws0r",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1xupfqk73y7rmc6qdgv7rtjy8wsngvt2g2t59t3",
            "coins": [
              {
                "denom": "inschain_validatorToken",
                "amount": "1000"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1xuwcmdvef5xudnyqa9gh5vqlq6ff0meu87zjxh",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1xa8z3sza70wrt06yfhlxq0ujg6fa7ld8h7m6xz",
            "coins": [
              {
                "denom": "steak",
                "amount": "35"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1xlkah6d2pcr8rchunmn6qdmdptr02844eee6le",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1xla8lu6qmk8j0wwlxzd9txclurglrp4r34khhz",
            "coins": [
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr18zwvgsrl3gc8y0fhzdev378u9gn37d9024njy0",
            "coins": null
          },
          {
            "address": "cosmosaccaddr18z3pnjgdtt337z8g5drfts7r3fm6n9a0896h0r",
            "coins": [
              {
                "denom": "coinoneToken",
                "amount": "1000"
              }
            ]
          },
          {
            "address": "cosmosaccaddr18raps2wmm63n2gldsrqc5x5z023p559l7r5vv6",
            "coins": null
          },
          {
            "address": "cosmosaccaddr182rhfxwz8ycec8sjq5frdg8qpwcnwrclu2gnya",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1825jrcgu3keurjlnpsujx9y8lrzgddjfsmax0n",
            "coins": null
          },
          {
            "address": "cosmosaccaddr182ujqw3r8p5fffjqkf0rxzj29pg5q96nxd2khq",
            "coins": [
              {
                "denom": "UmbrellaToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr18tkxwzhyn4a5ss63vttwpespvm7kxm895m9ppf",
            "coins": [
              {
                "denom": "steak",
                "amount": "20"
              }
            ]
          },
          {
            "address": "cosmosaccaddr18thamkhnj9wz8pa4nhnp9rldprgant57ryzag7",
            "coins": [
              {
                "denom": "steak",
                "amount": "380"
              }
            ]
          },
          {
            "address": "cosmosaccaddr18wfz5vj26y079surms5sm6avjtezzspfvqs6g4",
            "coins": [
              {
                "denom": "shensiToken",
                "amount": "1000"
              }
            ]
          },
          {
            "address": "cosmosaccaddr18w20rry3cclx23z2r70j8cd8an79kudja6s6gj",
            "coins": [
              {
                "denom": "steak",
                "amount": "20"
              }
            ]
          },
          {
            "address": "cosmosaccaddr18sqscwa5g4q6quc97eggwqk4wyzjyvrmvtq7f2",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr183ytuac46uh2q7jw4r8phvueljn7pfy3gfu3tv",
            "coins": [
              {
                "denom": "steak",
                "amount": "8"
              }
            ]
          },
          {
            "address": "cosmosaccaddr183g7an5mvj347annxtsyw05etu6ezx400ykte5",
            "coins": [
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr183tl5h22xfz8e8n5dp3fphu55kltfxxtr54vyw",
            "coins": null
          },
          {
            "address": "cosmosaccaddr18jtgkmw6ahnxd5lllp6256kuex2u2wsc5qt2u8",
            "coins": [
              {
                "denom": "steak",
                "amount": "35"
              }
            ]
          },
          {
            "address": "cosmosaccaddr18ja9lhun5le78j36s8rs2je6z4pzc6far7xfuz",
            "coins": null
          },
          {
            "address": "cosmosaccaddr185n9jln2g4w79w39m0g85x2tns6uekqg99tq7u",
            "coins": [
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr18cswjuy9qzwcjnkmvp32ryyht0rlaueazynegn",
            "coins": null
          },
          {
            "address": "cosmosaccaddr18cut539f6w9xewnz0fcxa95784pg4mk5ghk3fz",
            "coins": [
              {
                "denom": "steak",
                "amount": "30"
              }
            ]
          },
          {
            "address": "cosmosaccaddr18m09d56pg5p2660de4sjfezpd8ud6jfghndfnt",
            "coins": [
              {
                "denom": "starfishToken",
                "amount": "1000"
              }
            ]
          },
          {
            "address": "cosmosaccaddr18up28pmkspg2r8uv879fdlrxg0gpxnnj5nn0gf",
            "coins": null
          },
          {
            "address": "cosmosaccaddr18u2sqnuetfnkugp59e9pgyv2dpuvkkxmmsc7m8",
            "coins": [
              {
                "denom": "dooroomeeToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "50"
              }
            ]
          },
          {
            "address": "cosmosaccaddr187en8ww5322c3e6y6c67daxugs8f93dldyw73g",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1gq0qecxs8xdaarrqxxazwavwxm7qz5jzs5anvt",
            "coins": [
              {
                "denom": "certus.oneToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "50"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1g8xwmm3tp7pgdrmjgt4fm7d26anrrnle4yavuw",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1gghhdtx2sceafvgg8dry5sqrvc8srxghm4qqsy",
            "coins": [
              {
                "denom": "gregToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "50"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1gg6natrtt5lf02xwr06ujcczvavl54wgljuaut",
            "coins": [
              {
                "denom": "mining-shipToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "50"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1gjwzk6vcl89hdyl95agsvqng7kuf47nlhnkdx3",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1g4q87u438qdh2c8ue4dzdr9ldrqrs77ptm9c70",
            "coins": [
              {
                "denom": "colony-finderToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "50"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1g6sc5t9t68vmcj3alk7dfqr54tvastpxac28k6",
            "coins": [
              {
                "denom": "VNode01Token",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "50"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1gujxamhwhfh6ndeuc96u4m9y2hrhsnurce786l",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1g7re7m8fq0qaumpwztmgptrhw9ravqa4py0qsk",
            "coins": [
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1fzqwytcm25gsxg67l2z3rtmskpx75u6l5a2mq7",
            "coins": [
              {
                "denom": "steak",
                "amount": "5"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1frdjhnkgvlq5s0v5unf9njzlshdtsruv4z3sey",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1fyevhdmxpkz8stjxggv9wejf9caed02vyp8y5p",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1f9mxupaw2sx78hren09pga6ndc5z3urszq94c7",
            "coins": [
              {
                "denom": "steak",
                "amount": "5"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1fxfr3e277d3a7jkqqry963ktep8fftpz5uyxv7",
            "coins": [
              {
                "denom": "steak",
                "amount": "35"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1fxtfxtx4y46svzs3qtavlcnn0vtnlryad29jgy",
            "coins": [
              {
                "denom": "coin6Token",
                "amount": "2"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1ff94l3uh7gans5qjgj7j3tjfwsg5sl4ytq25ad",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1ftrggumlvkesqny90par3gc9lp70yqhtlqucsk",
            "coins": [
              {
                "denom": "steak",
                "amount": "43"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1fskmzyt2hr8usr3s65dq3rfur3fy6g2hjp23tm",
            "coins": [
              {
                "denom": "ATEAM1Token",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1fnklmrcpwnlv0pl85vg4h7evpl6wjdnvhacts7",
            "coins": [
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1f4sah2y8r995wsazt2dw909akhdjvr4f9zv28r",
            "coins": [
              {
                "denom": "steak",
                "amount": "5"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1fhfcs5npydv8s96wrk9p5ychctslu92t4n5qd4",
            "coins": [
              {
                "denom": "davinchcodeToken",
                "amount": "1000"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1fmk62wmcufp8dgemdrcxme88lxxnyxq9a93stc",
            "coins": [
              {
                "denom": "steak",
                "amount": "15"
              }
            ]
          },
          {
            "address": "cosmosaccaddr12qprqyxtxxlu7unncd9rrhmer5emg2pmrw27qk",
            "coins": null
          },
          {
            "address": "cosmosaccaddr12qnaj7mwyjjdw4506s2d47j7cfdtkg3jy9vdqh",
            "coins": [
              {
                "denom": "steak",
                "amount": "25"
              }
            ]
          },
          {
            "address": "cosmosaccaddr12pn4p6r5wjpsep9kn655ng7fh59yez7t0rahru",
            "coins": [
              {
                "denom": "oleary-labsToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "50"
              }
            ]
          },
          {
            "address": "cosmosaccaddr12zpkulxv7kn59sgf0tpf24qhqzxsvf3gamkl7g",
            "coins": [
              {
                "denom": "mwnode1Token",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "9"
              }
            ]
          },
          {
            "address": "cosmosaccaddr12y7e5t78tw7d9z20c2rfvl4ne3jll9nuvk6s7l",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr12xmjdkyf6stle05tdls28u4aelzep0rhzhllyc",
            "coins": [
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr128ty3kzhcepjacu4q0xjgq60qa3zz8na3jl793",
            "coins": [
              {
                "denom": "nuevaxToken",
                "amount": "1"
              },
              {
                "denom": "stake.zoneToken",
                "amount": "1000"
              }
            ]
          },
          {
            "address": "cosmosaccaddr12fm2n7pkmtqk7untyhs9my749jyzef3vhne95q",
            "coins": [
              {
                "denom": "steak",
                "amount": "6426"
              }
            ]
          },
          {
            "address": "cosmosaccaddr12wnqvqg79s9jqrul0lva3h72rfmewm7jprrcp5",
            "coins": [
              {
                "denom": "Nemea7000Token",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "50"
              }
            ]
          },
          {
            "address": "cosmosaccaddr120skmenn2a0ra8y6zassrxvnfc5rlme8rqarvs",
            "coins": [
              {
                "denom": "aetherToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "600"
              }
            ]
          },
          {
            "address": "cosmosaccaddr12ceualfg92x7du73mtcv0zya4nxvq3tl2m52uz",
            "coins": [
              {
                "denom": "cryptiumToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "119"
              }
            ]
          },
          {
            "address": "cosmosaccaddr126ayk3hse5zvk9gxfmpsjr9565ef72pv9g20yx",
            "coins": [
              {
                "denom": "grass-fedToken",
                "amount": "1000"
              }
            ]
          },
          {
            "address": "cosmosaccaddr12uggfcy84ghj4sepeyd95ym2wahw2l73fhcju3",
            "coins": [
              {
                "denom": "steak",
                "amount": "3"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1tqgmgng39tars3vgsmvdsk6n2yu04l66vs05pt",
            "coins": [
              {
                "denom": "steak",
                "amount": "40"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1tq0zwzyc88l2enrlhtzw0he8rm24xfd5s9aeer",
            "coins": [
              {
                "denom": "D2R-validator-7000Token",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "50"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1tywpgw32j0fntw4mjdm9l29j8xp3faj3ph0ylm",
            "coins": [
              {
                "denom": "steak",
                "amount": "11"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1t9rqvvys4zh300x93a6fr3h064gp9y8al7m4en",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1t99umqef9m3yscftgm6jh6gacxqad7t5xw93xg",
            "coins": [
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1t9nxlu834au3u7mltyf8udzw7kzsf0wntx7e4w",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1t9l4ztfwk0mp3za3x38yfqkernuclxwgdqzme6",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1txmm2ya9tdf378a9awnglzp7mwj0qsmsm3y397",
            "coins": [
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1t8lqjkvy0w0fzs3kv4s00s9cnlqqwepmnux9g5",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1tgwg9ewvmhsf8l7v28ugqmf07s5e2dpxj9qpxz",
            "coins": [
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1tsysfw53f36ap440e5ljf6hnc3ueuwal9vmcx4",
            "coins": [
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1t3afuz2kt99jz4pe5k4vjvkdmldn2x0lqzv83w",
            "coins": [
              {
                "denom": "BlissDynamicsToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "9"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1t3lvwnp0rmfyc96c0qh6qa67nplpdhytr5dd59",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1tc6jhusrwymrus63rl7yrk37z5jdksnytahuj4",
            "coins": [
              {
                "denom": "steak",
                "amount": "870"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1vqhlc253yxx3457lr4grawaqhu8xuphm45eld5",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1vrc7zpg5teawwuzkfh6t7c6sy353sukhlarxpa",
            "coins": [
              {
                "denom": "figmentToken",
                "amount": "900"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1vv9d030rnupxs2nrcnujhwzksy8fxz5r7t5n94",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1v04pxgxepjx5jxcayndexnwlhumy0jpsfcn5e4",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1vsg7a4q6vrcwlqsfm26je3xzzepwnjg2pss5vk",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1v4d9c3wy6n7h4udj4lkumc0l99x95cml44uyv8",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1vk5dwhn8m3sht9xrartj9kuv098t00spkqphcd",
            "coins": [
              {
                "denom": "steak",
                "amount": "21"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1vmdgku2g0n5p2ef995r4fugu99ze9e5me9kh4d",
            "coins": [
              {
                "denom": "bmen-companyToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "50"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1vuc8ex5f02z4hg5jm4lztrhv9j8cyt6cy0m4ey",
            "coins": [
              {
                "denom": "steak",
                "amount": "2"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1v7rge7kq75ey2f62vktv36ackzsfell3u6wrw8",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1v7hutth05td5242x4y0nh9gq03qumdmz3tayay",
            "coins": [
              {
                "denom": "steak",
                "amount": "20"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1vlmpl003lzm4lq0vr94227l2u02n2y490gsc89",
            "coins": [
              {
                "denom": "steak",
                "amount": "5"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1dq7jt3pn7mrce2twac907jue7hjd0p0rgt3qnq",
            "coins": [
              {
                "denom": "sparkpool-validator-02Token",
                "amount": "1000"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1drycruvuzjn4d6hx39hxz6jsz6hwkf09a8cyxy",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1d2t6h36cvqh6clmfk6w3u7spehv7ql7alcpfr4",
            "coins": [
              {
                "denom": "steak",
                "amount": "2"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1dw4dajgyff0pgmtjdv8t7hqp5va7xdadwpp7n6",
            "coins": [
              {
                "denom": "steak",
                "amount": "4"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1ds3yvnzup07u4j778xmnrh624j096fma3hw4nz",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1dkm8z36m96k8aukg5rxnfpw46dt5mnmvsfuze2",
            "coins": [
              {
                "denom": "steak",
                "amount": "5"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1dcgqsyam6d0am5z2m7dufp7s9k2px2kcqaswxt",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1dazry9zqmwduf7svv4extay64jnlg2yghfk482",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1d79807agakq8k5aeykva98gzem60f9k2kpetm6",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1dl9mu7uya6mfsd80wyw2rgjmwjtk5ajjf8mxcz",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1dlgd8tksf635zntchgc3kh60fkgpe9sstkql0l",
            "coins": [
              {
                "denom": "steak",
                "amount": "3"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1dl0hk7064k5wp43zkllmd0p9pgu2mnra898gyh",
            "coins": [
              {
                "denom": "steak",
                "amount": "20"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1wz3smpf4grhd7f82e3u3a8lk32kvq7s2raxudx",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1wzelzhl5a0hm403sgy8ydy7pnmwjjx0thu78et",
            "coins": [
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1wxf0ck9h2e8p2wmecxtep6cefhexsp4kzc8fxy",
            "coins": [
              {
                "denom": "finalityToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "50"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1wx33m9dvglryga0saey0pr99ja0klhcfrwaw7l",
            "coins": [
              {
                "denom": "snaticoToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "50"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1wf9nug554ugpw4t0wnlppxc6crl2n02qr8v3cd",
            "coins": [
              {
                "denom": "steak",
                "amount": "50"
              },
              {
                "denom": "sunny-mintorToken",
                "amount": "1000"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1wf3sncgk7s2ykamrhy4etf09y94rrrg4k73wwr",
            "coins": [
              {
                "denom": "steak",
                "amount": "30"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1w2j930e3ueg8f365ff2t92m5lnp856fe47f3hy",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1wd965wxzyl2krp67sdwddy4ep48xhsrdrq63rd",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1wnwr086pvkv2pqa5khznunm93ac52u677zgqe6",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1wk0t6na230vxhf6ddresw2c40k5y3ayrww0s2m",
            "coins": [
              {
                "denom": "P2P.ORG ValidatorToken",
                "amount": "1000"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1whqnxfj5fakx2tmjsja8y9ev7fqdzdldaesmx2",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1whxd48da3r56n8eecym8zg0c6xmf35fn2myart",
            "coins": [
              {
                "denom": "devToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "50"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1whd2u62u3zk7w09yuq4cz97fueq8e96qp69xz0",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1wesup293xcfjze34u98j05ffu04cssy32gjv3u",
            "coins": [
              {
                "denom": "steak",
                "amount": "5"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1w7n4zcftpv4yvsjjlegmcgk25kslvy0xdpk03v",
            "coins": [
              {
                "denom": "steak",
                "amount": "3"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1wlqylrnsjyxtex97fq762tt3xx570rv6ydqpja",
            "coins": [
              {
                "denom": "steak",
                "amount": "2"
              }
            ]
          },
          {
            "address": "cosmosaccaddr10qp8kqwm2cuql0hw2az5mngpmw5xm9ee32exlp",
            "coins": [
              {
                "denom": "steak",
                "amount": "50"
              },
              {
                "denom": "xiaochinaToken",
                "amount": "1000"
              }
            ]
          },
          {
            "address": "cosmosaccaddr10z0xk35lmfwfzk3tgzljvvr4nvfaj4778y5kgz",
            "coins": [
              {
                "denom": "steak",
                "amount": "7"
              }
            ]
          },
          {
            "address": "cosmosaccaddr10yx3zn39va2rdvcn76cljlzn7zmlyecc7wdtvq",
            "coins": [
              {
                "denom": "steak",
                "amount": "25"
              }
            ]
          },
          {
            "address": "cosmosaccaddr10t65mmpnd4r9uzq7hqachgzw0ct9uqea9r5365",
            "coins": [
              {
                "denom": "steak",
                "amount": "5"
              }
            ]
          },
          {
            "address": "cosmosaccaddr100ztxxmwkqr44jts3a543xhme7vjychzyagz3g",
            "coins": null
          },
          {
            "address": "cosmosaccaddr10s62wd8qca4pnhdk602xqnvpy0ck8uzll4hrsf",
            "coins": [
              {
                "denom": "steak",
                "amount": "4"
              }
            ]
          },
          {
            "address": "cosmosaccaddr10j2c9vspxg0rxfy2am7h4pkswdga3jzwk0x0nv",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr105fhlukl4s2a9h7kv86lf69ayy0mv6y7quw8v5",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr10505nl7yftsme9jk2glhjhta7w0475uva87paj",
            "coins": [
              {
                "denom": "ForboleToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "29"
              }
            ]
          },
          {
            "address": "cosmosaccaddr104e3vqacz2kncw979dczttp5jwts3sacv3ksz7",
            "coins": null
          },
          {
            "address": "cosmosaccaddr10h3v6phzsfuw6yvwpdvv0a4q3ha9zkq67gl8w6",
            "coins": [
              {
                "denom": "steak",
                "amount": "60"
              }
            ]
          },
          {
            "address": "cosmosaccaddr106ccejx2h7hwzj9xxaz3hju22yyq5mfy4rydav",
            "coins": null
          },
          {
            "address": "cosmosaccaddr10mfxvxjga3rxp5w93rx8jdxtt68j5qk0wv02hn",
            "coins": null
          },
          {
            "address": "cosmosaccaddr10atd0mjtdt7uzk0jgr32c32semnh3xvu4nx77p",
            "coins": [
              {
                "denom": "steak",
                "amount": "5"
              }
            ]
          },
          {
            "address": "cosmosaccaddr10lur643c66xdn4tsty4lgzwhq9etmhj850sc7j",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1sqec04r28jfevkzfkdj4ql2qzr2zwrmg78qzj8",
            "coins": [
              {
                "denom": "lambda-mixerToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "50"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1szlc8n7c2drvhxcaw0cs4h03jrstzjkl9yamj7",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1syhzpjwx6vv3erv2myq7txrjhrp95hrhgcl242",
            "coins": [
              {
                "denom": "TropicalMongoXToken",
                "amount": "1000"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1st4262sfqk7jqdwp2s7gk2umm5yhexhjrew53z",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1svu9nw80fzurmevlku9mzeqxrr8tpl5550034a",
            "coins": [
              {
                "denom": "steak",
                "amount": "5"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1sv7kdmv2haprv3c5ap62aar68d2apc9a7fwmnu",
            "coins": [
              {
                "denom": "steak",
                "amount": "11"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1swydkj5u3jd4jwe7uygu4479zgs4qg6v4ds3c9",
            "coins": [
              {
                "denom": "steak",
                "amount": "45"
              },
              {
                "denom": "stereo-watcherToken",
                "amount": "1000"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1s0p83zm5uxjppf8x9ugyp6hpta9kagl44ke7as",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1s4w7w56ymm5fnm25nhdvae5yuc5792zxfynq9h",
            "coins": [
              {
                "denom": "steak",
                "amount": "13"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1shuqhpl273t96yg6nnqvyfeewj3ew3mdcwvcnu",
            "coins": [
              {
                "denom": "irisToken",
                "amount": "960"
              },
              {
                "denom": "steak",
                "amount": "18"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1smg58zu028xkq53c8avp2kqhzcaf8wgq5jzrww",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1s76r4v04tt4fq6cvr39y6wy5zla9p6kdmz0wkr",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr13q937pwglh24knwa2v23ml0kkpl9vwzjmfmj3q",
            "coins": [
              {
                "denom": "iaspirationiToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "50"
              }
            ]
          },
          {
            "address": "cosmosaccaddr13q0fl9qakvt5zkh6qurlrur4xq7gullkv9f5zq",
            "coins": null
          },
          {
            "address": "cosmosaccaddr13psmeywshfyfcg3tr5l4qm968smtkcu8um0nqn",
            "coins": null
          },
          {
            "address": "cosmosaccaddr13rt2suxfgefcquy9pveuc5uxfayy5rscr4vld7",
            "coins": null
          },
          {
            "address": "cosmosaccaddr13fw4ff9u2qftsjj39mda54ufrpjs77q0ke6m49",
            "coins": [
              {
                "denom": "steak",
                "amount": "33"
              }
            ]
          },
          {
            "address": "cosmosaccaddr13fet8nre9ra2rs7gtw035evk2254pgw9mpn79z",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr13tqmv9rrlqy4rm9kq6en48vsauwh9gjv0vv4sv",
            "coins": null
          },
          {
            "address": "cosmosaccaddr13dre54wd7md0wg77x5fj5scj5u20e6stvdr829",
            "coins": [
              {
                "denom": "steak",
                "amount": "8"
              }
            ]
          },
          {
            "address": "cosmosaccaddr13ddmgank8chffte8zgnf7e8eqr95tc26elk069",
            "coins": null
          },
          {
            "address": "cosmosaccaddr135dz5hdtvk3z8vl0zyy5l223kjanv0gudu4905",
            "coins": [
              {
                "denom": "SaiKrishnaToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "50"
              }
            ]
          },
          {
            "address": "cosmosaccaddr135shy2er3weyxf44cp4q0whmc4rkuqcup3389j",
            "coins": [
              {
                "denom": "steak",
                "amount": "20"
              }
            ]
          },
          {
            "address": "cosmosaccaddr134muwl6a0n8suc5gruj6gh2ka74wkh4kqh8vm4",
            "coins": [
              {
                "denom": "figmentToken",
                "amount": "4"
              }
            ]
          },
          {
            "address": "cosmosaccaddr13kftrc7h866cnwxuephx9rgj2hdfd6dyuc4zsu",
            "coins": null
          },
          {
            "address": "cosmosaccaddr13cds7hwqyq9ja2hsv4sg7glq9arlk43gcl3cek",
            "coins": null
          },
          {
            "address": "cosmosaccaddr13lsvgpu9lptwg0ef3c8ykqf2x35yjvndfqeqkj",
            "coins": [
              {
                "denom": "steak",
                "amount": "2"
              }
            ]
          },
          {
            "address": "cosmosaccaddr13lcax0d224sch59zz76m7n60nw0qfzty957l3s",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1jp66q0yftj6rhwqs2n9uxnmfm78et8gfplacpz",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1j9w65cmt5s5kw9430jsv87yd4dgmmq6w60f3tl",
            "coins": [
              {
                "denom": "steak",
                "amount": "9"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1j2frwt2vq2xer7f060wpsu3y3f63uys2w9lx2e",
            "coins": [
              {
                "denom": "steak",
                "amount": "90"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1jtt6al0acr8h0uvq489rt9zx9lnau7rlcu30pt",
            "coins": [
              {
                "denom": "JColToken",
                "amount": "1000"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1jt77t00gucffqhrq7lh6vwmwzrwgw2fpvnhwk7",
            "coins": [
              {
                "denom": "steak",
                "amount": "50"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1jvh448tvk368k4md83ys7auledclek0vfpckz2",
            "coins": [
              {
                "denom": "ramihanToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "50"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1jdrxh34d00rm8ydclhxgkzmgp55y67m54umags",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1jdehygp2p8z4f2wyf9t5q6f6mfqex66mfhfwdg",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1jw554408yw2h438q200jyuqgln76xh4ax0q4s0",
            "coins": [
              {
                "denom": "TruNodeToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "132"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1jsvgewy7h62v3q43m0l347wlcwyhd4un5q8aa3",
            "coins": [
              {
                "denom": "COSMODROMEToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "50"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1jsch8k385shvse6j5dfx20483qym5uhq76xpjf",
            "coins": [
              {
                "denom": "SVNode01Token",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "50"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1j4k7p2yqn6d5kfjqg0ductjmr48l75pw8tp6w8",
            "coins": [
              {
                "denom": "steak",
                "amount": "20"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1jkmn6asju47zuuzrf8rjt7sllaj5cx4kueyv8p",
            "coins": [
              {
                "denom": "w1m3lToken",
                "amount": "1000"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1jhsuj64hgtuptm07safql27hwzm9jht50zxzre",
            "coins": [
              {
                "denom": "steak",
                "amount": "20"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1jh3grawl62juvx5p8fz5xsy9hpw8w3mngqafe4",
            "coins": [
              {
                "denom": "block3.communityToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "115"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1jc0nw5s49pxdjg2nh2ew7rm9r7c566mh5z4a93",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1jck6gp4mqy33sk6a0fr5c8maq53hf4245v3mgg",
            "coins": [
              {
                "denom": "7768Token",
                "amount": "1000"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1je98ejgpvv2uzku3t6h5cydyr9hzva2cqhf3v5",
            "coins": [
              {
                "denom": "steak",
                "amount": "5"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1jegatcczazsuj86at74pwcv5ew0q29k8r2h3nn",
            "coins": [
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1j68chllpyc0myr5fjhpw5uzmfntpgdlhuurk9d",
            "coins": [
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1j6uaeggxkn0t4n4n8zna0g9vqgfj6jw6ejwyv4",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1ja98tyxchsde6v3l02thp4fc35eymg88kts070",
            "coins": [
              {
                "denom": "steak",
                "amount": "2"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1nqjlq0r3d27dq4mujfjd83wyy94zx746jdnzk8",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1npga2sl3xt8kn0st2fdjg2x73zj2s0xg2znumn",
            "coins": [
              {
                "denom": "steak",
                "amount": "2"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1nzh0av9pyd36j07mjwug9eddycywzng4d0843j",
            "coins": [
              {
                "denom": "steak",
                "amount": "50"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1nypejllw6gku8cgzccpdp07ae0j442xphdrlen",
            "coins": [
              {
                "denom": "steak",
                "amount": "130"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1nyje8wp0w9t9c6r7ycx75p29zpxv8hrw0fu0mf",
            "coins": [
              {
                "denom": "steak",
                "amount": "20"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1n9qsfp3x09rhhxptzrgmdlucqhg2ce95fm3fw8",
            "coins": [
              {
                "denom": "juelianshanaToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "50"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1n8yt3faem889lpr8x7f0pe0wth8003aclldlr3",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1ngqq8lsadzerg66wcgv3hx86gaefedjhpyvakh",
            "coins": [
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1n2jad3hurvgfxamr2ee0mmv5tsckyxmesqz294",
            "coins": [
              {
                "denom": "steak",
                "amount": "3"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1n0l69ax723z2n060eeh357c0vsz2kagf06vscz",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1nnyel6v0kx8zmfh9edmre3ua4dt9306cfxsxgd",
            "coins": [
              {
                "denom": "spptest1Token",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "50"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1nn6mjdtcutmmpjcapkecg5muqrgp7td6s0gre2",
            "coins": [
              {
                "denom": "steak",
                "amount": "30"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1n54p3gh6vrqarcdqxd8xqvpt5p6slnqhta55v6",
            "coins": [
              {
                "denom": "steak",
                "amount": "5"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1nkgjjd9nvhamh8gp973y5vfmze3dsa550t8dkx",
            "coins": [
              {
                "denom": "steak",
                "amount": "2"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1nhnhk0q53lfmwhq58268302zxz3trpgrwgr02w",
            "coins": [
              {
                "denom": "steak",
                "amount": "30"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1na3wr7ujdp3qdej6l5y0k4znzrkaz77t2yjaqf",
            "coins": [
              {
                "denom": "nuevaxToken",
                "amount": "999"
              }
            ]
          },
          {
            "address": "cosmosaccaddr15pzn94awwmcff4m9cmju84t59t0uj5m65fwtl6",
            "coins": [
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr15rrxdflhq4744qt9m5ysstq3zpykhacf908eez",
            "coins": [
              {
                "denom": "wancloudsentryToken",
                "amount": "1000"
              }
            ]
          },
          {
            "address": "cosmosaccaddr15yqm96jm0mhlluerr2ze0u369vqjdhfye2f9e9",
            "coins": null
          },
          {
            "address": "cosmosaccaddr15xlevh5an4gqnyvjs7yv5ccpp9rvyrcfemteey",
            "coins": [
              {
                "denom": "steak",
                "amount": "20"
              }
            ]
          },
          {
            "address": "cosmosaccaddr15twdq9t9pnmr2y5ynccqfjdkj5wt5vzsq0hp0r",
            "coins": [
              {
                "denom": "steak",
                "amount": "20"
              }
            ]
          },
          {
            "address": "cosmosaccaddr15dt7zafut73my95xclacfjldschxjnesye6385",
            "coins": [
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr15w2rengajq9js8hu57kjw88dly5vy7gsqedn0n",
            "coins": [
              {
                "denom": "kittyfishToken",
                "amount": "1000"
              }
            ]
          },
          {
            "address": "cosmosaccaddr15wkl5rq6uk8qdnpehdkp7clng255e06z2p9lfk",
            "coins": null
          },
          {
            "address": "cosmosaccaddr15wmreh6mryt49ly2v3e5gaw8ssv49xmzqrxqax",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr15nhcasamcyav7gs34nezq0r0g9zyswrzh6gt3f",
            "coins": null
          },
          {
            "address": "cosmosaccaddr15k7ws7jddz6qfnnclg5h27q4gpjeyamkgupeaw",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr15klp8pypkd8htppzt6pguej57yyvp5p442khcu",
            "coins": [
              {
                "denom": "meteor-discoverToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "50"
              }
            ]
          },
          {
            "address": "cosmosaccaddr15mzck4rm5w55sfgr9kgexgtna22v608907t8y6",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr15u9ve7fz8fqaf7r2u3p4f9ru4ze47pau5cxgcg",
            "coins": [
              {
                "denom": "mpaxeNodeToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "80"
              }
            ]
          },
          {
            "address": "cosmosaccaddr157eu0k3fvnk4d3rpgtyz2jujf8shlk8mxeqjkk",
            "coins": null
          },
          {
            "address": "cosmosaccaddr157mg9hnhchfrqvk3enrvmvj29yhmlwf759xrgw",
            "coins": [
              {
                "denom": "jlandrewsToken",
                "amount": "1000"
              }
            ]
          },
          {
            "address": "cosmosaccaddr14q7h5s8w0fgruyv3qwcauaxpy7cmu289hd8hsh",
            "coins": [
              {
                "denom": "steak",
                "amount": "20"
              }
            ]
          },
          {
            "address": "cosmosaccaddr14zfph0h8rexsca6gg6jkwqup3sgl6mwj6eu4e6",
            "coins": [
              {
                "denom": "smartpesaToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "70"
              }
            ]
          },
          {
            "address": "cosmosaccaddr14dwnmm6n7tjdpeylpwsdsatdl0umm75dfkqcpa",
            "coins": [
              {
                "denom": "gazua1Token",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "30"
              }
            ]
          },
          {
            "address": "cosmosaccaddr14sqp5mqtmfjdk9khae6lkfytwj9crtunpzr8f7",
            "coins": [
              {
                "denom": "steak",
                "amount": "5"
              }
            ]
          },
          {
            "address": "cosmosaccaddr14e774gfzt5l9ka766ehfgu6n5hgy9f3sehzyv8",
            "coins": [
              {
                "denom": "cwgoesToken",
                "amount": "800"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1kzp74plx94h8mp90vun33nmdw3qck7mecs6hdy",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1k9pqxd8fxqlk52uwfxnlsexqj6xnmw5swhss45",
            "coins": [
              {
                "denom": "jjangg96Token",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "120"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1kdrp0qh5enr4vc3xvl0ks6wf583tg0nc30z5k8",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1kwftgznylzerft2ksgkzvvfn5rfpy4nk2ye8na",
            "coins": [
              {
                "denom": "chainflow08Token",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1k0lw4fzp2adfh30ranet6mfujhzdywd6l22nag",
            "coins": [
              {
                "denom": "steak",
                "amount": "19"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1k5ah4mpy8ykzp2mpcvapnyv6gemp35zuyhw8sy",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1kcksf3e9z3dalfck9ugmcdjlxcu44s2krjt6k2",
            "coins": [
              {
                "denom": "cwgoesToken",
                "amount": "99"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1k7u0ysqy89jkt7n9804cl4vzy45afa5jqevx6p",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1klslzz2n2nqvdaa6gwflwpka7gc60vvmh3k450",
            "coins": [
              {
                "denom": "idoor7000Token",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "50"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1kl4460nu3vpgkh7rqhnf5jcsr5ckdsm6v9eyv6",
            "coins": [
              {
                "denom": "figmentToken",
                "amount": "50"
              },
              {
                "denom": "steak",
                "amount": "35"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1hqclw9xcymkhtrazzq9x576v9khhe6j9ja3jp0",
            "coins": [
              {
                "denom": "steak",
                "amount": "3"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1hydk58dtxr87ml34xqc07ph2h66yt2f0rrhyt8",
            "coins": [
              {
                "denom": "steak",
                "amount": "30"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1hfygmryre7r8m9pfqmc85y7kw7ejphmp59k2x8",
            "coins": [
              {
                "denom": "bkcmToken",
                "amount": "1000"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1hf2965c3zewcqapaey4mchyga88nd2rpqlldk9",
            "coins": [
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1hf0qr0z36v38geace28mfwyayv2r5r8kl9st8p",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1h2faa4f2vcg7rjss6y5wnswnrw4mh82hf4wrcp",
            "coins": [
              {
                "denom": "steak",
                "amount": "13"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1hwyayqqsm4nv570ux8u7dla0xn4ad08gydppyj",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1h4q0xkdg30cl9vw0u8ejm0rs337dszy98gnd4a",
            "coins": [
              {
                "denom": "joltz-secureware.ioToken",
                "amount": "1000"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1hhcffh2fjf5w42q7lmlcpsrg4agslmm2ercuqk",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1hhu6pvea5l934c059jxjv3ccfan6m7jy39zfs6",
            "coins": [
              {
                "denom": "steak",
                "amount": "30"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1h6el6v9rn0unkw2fh9s86juvg6fsvcmy8wy72s",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1hm5f52equcd4f268uxf4xhwljpnljvdyqal3dv",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1hmkg328mkqen9kc7z3j07hhu5ep9qh64q8kkfh",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1h7yr9pm3ffhq5djx75ycmumvpwlxhulss3e2yv",
            "coins": [
              {
                "denom": "steak",
                "amount": "30"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1cze0acxhmpx26awlc98h2v33ay95k232yenkqx",
            "coins": [
              {
                "denom": "steak",
                "amount": "37"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1cypmdvszcd9kd3jenmqxd03cpceql8rtuvxftp",
            "coins": [
              {
                "denom": "Gold2Token",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "35"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1cy5t4e62z4h7r74cn7vkylx77n5s4hds2xeud6",
            "coins": [
              {
                "denom": "steak",
                "amount": "20"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1cyayv35tv47t829mel6jm0vqzmrdhf3jq87tqg",
            "coins": [
              {
                "denom": "bucksterToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "50"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1cgvwwnfeu7jv3wxsy6znr2dqnhs5xc6cjpfmxw",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1c2dkaym4teqw6jl9jkq8eu8nf4wzn2lgf4ydyt",
            "coins": [
              {
                "denom": "faucetToken",
                "amount": "10000000"
              },
              {
                "denom": "steak",
                "amount": "9935501"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1csk8r7uahekf9a5qf5axsl00fjvemhkrfu8kep",
            "coins": [
              {
                "denom": "steak",
                "amount": "20"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1cnf4nztxfrtwcfztfzz3sreylfkxe2xh5z307x",
            "coins": [
              {
                "denom": "steak",
                "amount": "3"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1ckr3s08lm6mddnueaffaplnn4eet3c4n6h8zzh",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1cut9le04243e3fjq6jc8jysup8lpl3ngwa4ntx",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1c7r8w934puc05af2yssgwsdu3w4vgp6a5p2xaz",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1ep87dml7wrc0t6d4rh05np5evtqwlxy0kstpvy",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1ezlvjcfq5s6fzrzp0rgrfudkte280dy528v93m",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1eyd6g0h5h96heegs42djrptlq9zflx3q2cndmp",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1eg88jsn38mrael6cqu7d2u8j6dynya7fv2p2tl",
            "coins": [
              {
                "denom": "DoriToken",
                "amount": "1000"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1ev9gnq7z9dga7mr0nf403942hnuj39r8dx8m40",
            "coins": [
              {
                "denom": "steak",
                "amount": "20"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1e078whe3xwhagy2hjjsht36tq3uehz3v2cs2n5",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1e3qm8jd9357zhdemwnaafmf0wy3f4yqmd307c2",
            "coins": [
              {
                "denom": "firstblock.ioToken",
                "amount": "1000"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1e32c8x0dcgv0kz6qr7xceerne2gznt5s8e8tr7",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1e3kz5kwptnwaslhnynq7jlrnya4sswswe4xprl",
            "coins": [
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1ejhdq2gpj20schz5n2r7lyxt6pwp2usahaf3wz",
            "coins": [
              {
                "denom": "steak",
                "amount": "30"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1ehhkppuyz9mz9qeevqyvhrl2kyqrc0vuhz2436",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1emaa7mwgpnpmc7yptm728ytp9quamsvu9rk4hp",
            "coins": [
              {
                "denom": "kochacolajToken",
                "amount": "1000"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1eu22vuernyw7yemd29lnkmd4x2l5grvx9kh6he",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1eamypvjgl8hjfa4gysujkt780f0qv3c5q5a3tt",
            "coins": [
              {
                "denom": "steak",
                "amount": "60"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1elpemwnp36jjcltv80hz76kcrk4zmllg0p6wk7",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr16rcrnftjyl2mctz78825ng8tx5ss22jf6jcp9l",
            "coins": [
              {
                "denom": "luigi001Token",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "37"
              }
            ]
          },
          {
            "address": "cosmosaccaddr16yrgs9d5gdu0erduz0rnd5et6u5ux0n9uyxh3g",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr16yxdmg8w5vc038cg0myh0w32s4h74ch59j67u6",
            "coins": [
              {
                "denom": "steak",
                "amount": "2"
              }
            ]
          },
          {
            "address": "cosmosaccaddr16yk2su2ndeetfqvra5cwwcf3sk59yjagyz99a8",
            "coins": [
              {
                "denom": "steak",
                "amount": "130"
              }
            ]
          },
          {
            "address": "cosmosaccaddr16xae9kqqllsjak3ntxmlgn4s3jmvvyeqcwzgen",
            "coins": null
          },
          {
            "address": "cosmosaccaddr168q7hkcnwahkugh02m0dr293hpzgj9w3z68r2d",
            "coins": [
              {
                "denom": "steak",
                "amount": "70"
              }
            ]
          },
          {
            "address": "cosmosaccaddr162k6wg0ufzy6pe079m3evt0tqrjfrrath5qw0d",
            "coins": [
              {
                "denom": "steak",
                "amount": "5"
              }
            ]
          },
          {
            "address": "cosmosaccaddr16dfl4ruvt54qyhyf26fwmle62jnywp7mu9akt7",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr160836cl5pqyy4mmk2k5u5v5xwcuztc8fksrk3k",
            "coins": null
          },
          {
            "address": "cosmosaccaddr160wqlavl0sa4fzm6lgm92cdc0lgzpzhcu7sl7q",
            "coins": [
              {
                "denom": "steak",
                "amount": "4"
              }
            ]
          },
          {
            "address": "cosmosaccaddr163rcl6f2vfl6x9svkl49edqhdl6agpfzncgupl",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr163emaq8ahwp25pmfqrjaz3c7wm9vtytcwmlphg",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1640dsytftqn2uxk4prlfxjvn46nnp8wxt5s0ye",
            "coins": null
          },
          {
            "address": "cosmosaccaddr164jntjfk9zs8x29mc27qansfwvjqs60gj6ermu",
            "coins": [
              {
                "denom": "lunamintToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "280"
              }
            ]
          },
          {
            "address": "cosmosaccaddr16hvayg2lv7rmqz7rjsd3fnq3y34j6r97dxrecf",
            "coins": null
          },
          {
            "address": "cosmosaccaddr16m99y8vsmyenxt4d93qft4gs0czzk6np80y775",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1670l5up6e5fddvlc027yvvlvedrzyx0mmsl622",
            "coins": [
              {
                "denom": "cosmosthecatToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "50"
              }
            ]
          },
          {
            "address": "cosmosaccaddr16lz93dqwq6rmn3eflugjnz5exr3jpspe3akwug",
            "coins": [
              {
                "denom": "steak",
                "amount": "80"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1m9nu6haug2v25nsrudv0xevxs52zyjzqld3drr",
            "coins": [
              {
                "denom": "steak",
                "amount": "209"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1mx36cmexzz6dxntf6lxpsdsf8u7kuzerh07m4s",
            "coins": [
              {
                "denom": "steak",
                "amount": "110"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1mf66nhhpqjtqt4pp6uqzu933h9wkkrc4z7l8wv",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1m2pl4vq3g8t4am0zk8qme08dnwe2p95e37l33s",
            "coins": [
              {
                "denom": "steak",
                "amount": "40"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1mtez787cl6phgkylneq9wjj3d6letn0247lhf5",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1mvqtcm5a8pf8q5xv50ag6k445032fz7qaax5ge",
            "coins": [
              {
                "denom": "steak",
                "amount": "18"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1mvsaaudvvhkfxnkjev3r9g73nrjejvdfc36zz9",
            "coins": [
              {
                "denom": "steak",
                "amount": "8"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1mvmdrtuxjttv90wssvh6xreeguvq8tksm789m5",
            "coins": [
              {
                "denom": "steak",
                "amount": "40"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1mwmcpq5nhaen8zyy86mrrept2gmp0z5peapkhu",
            "coins": [
              {
                "denom": "abcinToken",
                "amount": "1000"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1m0u78hks3hcf5c6szvlly7x40dkattyal0wy2w",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1mndgxh7y3ayxy25erzuke0xgngac3pw8txvy8a",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1m7rtacq9ygng2zf2y4z3mer8hdtacg4hyspmrh",
            "coins": [
              {
                "denom": "figmentToken",
                "amount": "46"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1upnvzj3zwqu7vnd7w5mtay6ugn482jerv028nz",
            "coins": [
              {
                "denom": "steak",
                "amount": "3"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1uga4nuresex5u8ajjh2pcr39l0s9hszdkp843j",
            "coins": [
              {
                "denom": "steak",
                "amount": "50"
              },
              {
                "denom": "wingmanToken",
                "amount": "1000"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1uffeh2wq0lrvwm7z77c595v040q4kk4g3v8vte",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1usm7kdwff58zjyz3wy89qemetrqehklala7mkm",
            "coins": [
              {
                "denom": "steak",
                "amount": "2"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1u374sn4c4arc8nsnpqdaqm2s5n70gpg6xvw4ny",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1unqmzc0w20qnpu94n7xs4z9nrteh8flxag028u",
            "coins": [
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1u40w3tq6w9cycgdpy4s692e92k03wnt8824zhp",
            "coins": [
              {
                "denom": "steak",
                "amount": "5"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1uknhjkw775zp5xkztc4clxr240ut59ek5seerx",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1ukcuqpqw3eq505wkqd2adgn8ugewxr6jtakngs",
            "coins": [
              {
                "denom": "daefreecaToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "50"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1uep3zha7srhafrggd5dn3gp0l08xua0x24yzj4",
            "coins": [
              {
                "denom": "steak",
                "amount": "70"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1u6ddcsjueax884l3tfrs66497c7g86sk3vfmqj",
            "coins": [
              {
                "denom": "steak",
                "amount": "160"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1umaajfgap5ef6yxvk5706kwk8j08l7wh6h9fp2",
            "coins": [
              {
                "denom": "iqlusionToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "88399"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1ul3ef4trvjfmgaptu70ed4fjlky2aa49t4pfle",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1uluae4kaxq3wxd698t256a8luv2tf0ngy40qp3",
            "coins": [
              {
                "denom": "steak",
                "amount": "200"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1artqm53zt7j7znm8cpgqzgp9jhlx0xrtle2pjr",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1arlpxs2ftf5hgetqxxkvd7mqdc28mmaqclyv4y",
            "coins": [
              {
                "denom": "bharvestToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "50"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1ay37rp2pc3kjarg7a322vu3sa8j9puahqavh2n",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1a8ppz2vh0zjfh8vaqckp67ct0ajm9crv0t3v0g",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1a8ct2evn3944h2fww346nsmc22z9rlmv7g6flk",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1agfqwpcxpl8s8pl3uhxhw7h0r05an60dmlgukx",
            "coins": [
              {
                "denom": "steak",
                "amount": "18"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1ad44rugqnfazzxqep43svkvd7jj4d9e38ynn0n",
            "coins": [
              {
                "denom": "steak",
                "amount": "3"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1anuuffusmq5ng3rhlndhnacy45et30jqygtn67",
            "coins": [
              {
                "denom": "BarytToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "50"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1akrn2ekh8p0vmy6wkfqdxsr7h5j2q7vj05xh7f",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1a694yvd5c2fuwd79p0qsfsdvnpykqz250l35zf",
            "coins": [
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1aaumwh9rc9ksh7dwuxux43jd93e7t4u88g5kg4",
            "coins": [
              {
                "denom": "steak",
                "amount": "20"
              }
            ]
          },
          {
            "address": "cosmosaccaddr17rqsh3fw6rehnwzarz90jkqtt5fupmh50gy556",
            "coins": [
              {
                "denom": "space4Token",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr17gvlvfpsfl6hffn5u2hahk22un4ynpykc44tat",
            "coins": [
              {
                "denom": "gruberxToken",
                "amount": "1000"
              }
            ]
          },
          {
            "address": "cosmosaccaddr17wg97pzzlswd4p9rejdavp87rv4vdugmm8pwkf",
            "coins": [
              {
                "denom": "steak",
                "amount": "10"
              }
            ]
          },
          {
            "address": "cosmosaccaddr170yylzwpzx9c6er2mn086xtdy80w2dmwp9wa3c",
            "coins": null
          },
          {
            "address": "cosmosaccaddr17n0y36k4430006yc6m2r709t9uxk73yptwsrsg",
            "coins": null
          },
          {
            "address": "cosmosaccaddr17nlusdvrk34fq65jemy3umfjfwaxfzv4asyl60",
            "coins": [
              {
                "denom": "4455Token",
                "amount": "1000"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1lq0mvtgnwe8mp0096un0v8ztcsj8ad92t2cwrq",
            "coins": [
              {
                "denom": "infinite-castingToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "45"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1lzy0ztj9cz7lemlev99u0upulqjd5cu6nk3t2n",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1l9jt8xgejkcm4sulhwj8c83ftprz5p9lyq4605",
            "coins": [
              {
                "denom": "cosmosToken",
                "amount": "1000"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1lvmxcdhua2kt3qd86wnnej5m6gm4max4um5sx2",
            "coins": [
              {
                "denom": "steak",
                "amount": "2"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1lw2d0qljstqzp6fqhsd5jazackgad9y2e3u2n8",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1l0qw5znfd6e8pshpjvyghjjzyr4l6ymla080lt",
            "coins": [
              {
                "denom": "nyliraToken",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "48"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1l3agyg568dsk2r9v7f35wge3yr85ampza7cdfj",
            "coins": [
              {
                "denom": "windmillToken",
                "amount": "1000"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1lnsr37qy64k4026gak3hvt8j2yh6c0ctxu46dw",
            "coins": [
              {
                "denom": "steak",
                "amount": "13"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1l529pzwyqyh7c5gg2wvmnx86qve7hyzr5mlk85",
            "coins": [
              {
                "denom": "steak",
                "amount": "2"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1lhjta6nt0lewj05m8444tuyhalkrffgpm7njpp",
            "coins": [
              {
                "denom": "BFF-Validator-7000Token",
                "amount": "1000"
              },
              {
                "denom": "steak",
                "amount": "1322"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1lc4jjw8zveuuj3njhy70zl23vgmpds2x24efca",
            "coins": [
              {
                "denom": "steak",
                "amount": "250"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1ledtq3q59uws8x78fjp30jvyulgjxl6yrx9x7j",
            "coins": [
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1lmazkxl9dtetn5lxzl45ut3ywwn94xvs3l4z8f",
            "coins": null
          },
          {
            "address": "cosmosaccaddr1lufq4294z7mmv8pghu9c30jmrrr9629lnwgd2u",
            "coins": [
              {
                "denom": "steak",
                "amount": "40"
              }
            ]
          },
          {
            "address": "cosmosaccaddr1lamtaccmxtak3w87g2dc0dlmecn9r0xgumkgtz",
            "coins": [
              {
                "denom": "steak",
                "amount": "1"
              }
            ]
          }
        ],
        "stake": {
          "pool": {
            "loose_tokens": "212234692368559902773014687327484492499334064496342017769139923035458728144919945137085919281749652786862586180209457322113759521105809070531211470506853218958050414866880167260657863739551652192137986049110725956052905704758818959982320658920794675668552064072238944093207349839633292480904346056107796224020080112445238478513757301551906370650373478169897541122069802602945534940145786304533068814943533930972906449983847914418613883243454229191697608448325572106833582710351535919362394483029380480542586486261407813018972054510265531490657883920535958690038587946293781570140135933217290963138230954485933118636731702849516112455184967089435748052528036914591916846061696750735527083841046512146774731083493285079420185153635615506205679139707366838712372605550276090588769737701435671243899824765182112838167666029900569234536970608040574716511086893100048478224094190486876917185986809531136797511860121445977859007948872468785789703779795691443108195616502424916434207502486427828788956840836992496531085759517722416727807667390924847734758804566179091636514959392799069188372721367831331110458081955835928443892419095932693400385273950885920091488321433751751535289836946545305590199876648995139778938250490114289510274183655998132281709023181865632542357298066269777564343753095882413548444231877206287627063978602318337637493451025792225507132732684643687486970003518524413564431921084394163011620507576101798101868687109423014977363066898839160979876401484528151191566119497417614711893271336419362034551327787658513622548663272288015026248873461521204921520317957869695103087669105075589633490514906981823236481320900674887471684246836161077527577918361885107730942590068766549236646992402788358566097270158463105833136968000705430256339804514651753256418561907619221566022910606606764721518296229753733641382514940120234521365586915145709181638223891153820569480497157851844433542586180525308478072757485213896182698413125792187729794843162037178476565265231421389422260592759821033883/20891080884821288214936513078767214469995841625290057084591977940219676920201103573603743654086273216540116817461098873206255880046908414217131166077880682528389281138014445604719519941588711798520441515463960396551523169294059582356600171764959826348969443952432853391059045461602343063648470338757779668479373347477478521355103694977501171737505983368487928796369306792077357030331966812678021793060446683766024779586855353966404049083602927499437752904965368719257394516323420051267652055024786806061807282124874257024889193873082268069292215810584394051776563161580278635823211761332288683787297085577465315444622264536480307169606763243105908257500502993503793484760386558210386840141656073106458220614126868493299191464939098381355626390008684156692722414660824992906786009750166861660454111502952749111939065695767011842636683869084062347771270413109022400836402699141165569131339421912716913002878566963211095013746548273413182056902584622759826444565509891551815878120259377745017492491285251896611745779513234465971017710455120244774552433479282660133754670813719046534492669403007823721365077824812739657392186031098902304403076796004295479892678612582906711521817557244571481951756843512724064309431463095282564288969492675967736221948943329034584116924461238339562647788542136485763482574600083575883466642001774801289514924627144010449754930076548579082867620591947286500534079824430946876218758498309184347779714955120422161085340625362053168273912668361524974414018116452843496772223549567360891625083368517506494362533454336269091585886586953590300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
            "bonded_tokens": "31284000023/500000",
            "inflation_last_time": "2018-08-21T08:00:00.123456789Z",
            "inflation": "1/5",
            "date_last_commission_reset": "0",
            "prev_bonded_shares": "0"
          },
          "params": {
            "inflation_rate_change": "13/100",
            "inflation_max": "1/5",
            "inflation_min": "7/100",
            "goal_bonded": "67/100",
            "unbonding_time": "10000",
            "max_validators": 300,
            "bond_denom": "steak"
          },
          "validators": [
            {
              "owner": "cosmosaccaddr1qv68kxwmarzefrcgnu4trdrznslzdtra7gfxhy",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "WnV5CjfHHyMrAPbX796iXgnfAwTP/cFL+kzz99bW+kA="
              },
              "revoked": false,
              "status": 2,
              "tokens": "7",
              "delegator_shares": "7",
              "description": {
                "moniker": "baryon.network",
                "identity": "EA1BCB68BEA98362",
                "website": "https://www.baryon.network",
                "details": "Baryons are Cosmos building blocks"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1qvut6w8v5fccephuv502qhmvhn7hy6cnddrdyf",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "U9EhwT1J9AHv3N/Dy4c/ADGLj5J050pucJ9ycTGvweg="
              },
              "revoked": false,
              "status": 0,
              "tokens": "99",
              "delegator_shares": "110",
              "description": {
                "moniker": "AT-Validator",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1qjzjfn55hygaak9l9x04z792mexce2zd2gvkzk",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "2RYn/WYgNvvDb0NLcdJ70bDiuAVH+tLH7VWH2M59i1c="
              },
              "revoked": false,
              "status": 0,
              "tokens": "787/10",
              "delegator_shares": "548/3",
              "description": {
                "moniker": "#🔥WuTiao.Com🔥",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1qk2q3l28t6qhzyc7p832ep8uwd0ypstgqla8yx",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "tU8IqlII8g8ngcmtMVG/k/8SmunfJJBkAoOisN+qcKE="
              },
              "revoked": false,
              "status": 0,
              "tokens": "482/5",
              "delegator_shares": "119",
              "description": {
                "moniker": "JB-The-Valid",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1qkc3tghc3fms7eye7vtu0g0370epr4jkje2ne7",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "RBuRRNEzA9RA1Wrdi9PPFQJ29/n/bqN9O2tQv9Gq248="
              },
              "revoked": false,
              "status": 2,
              "tokens": "2164",
              "delegator_shares": "2164",
              "description": {
                "moniker": "skoed-validator-7000",
                "identity": "D889FE03C0655B11",
                "website": "https://skoed.io",
                "details": "Skoed IO"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1qkllv8f6qakkw3hk9dqvytys090lk6twsyv8vf",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "F2wCn9rKafNZsYZwoLGkSQIpr3rk86cjYyuhSjsjRaE="
              },
              "revoked": false,
              "status": 2,
              "tokens": "255",
              "delegator_shares": "255",
              "description": {
                "moniker": "ironfork",
                "identity": "",
                "website": "http://gameofthrones.wikia.com/wiki/Iron_Throne",
                "details": "Aegon the Dragon"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1qarnnunjrsd66yuz65ugsl7pm3v3hxrxd5jvaq",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "1XB37ELQL8SIu5CeeRNMByNFtf2rQf543PKoxUCBeBA="
              },
              "revoked": false,
              "status": 2,
              "tokens": "54",
              "delegator_shares": "60",
              "description": {
                "moniker": "sentinel.co",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "Sentinel Validator 1 on gaia-7005"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1pj6vfgetplkcuq0k9m98dpej6zew3pdest842v",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "r8ObA8qeSscl54T9I2SrdH5uUFxjbqYs2IjtU1ZxnAA="
              },
              "revoked": false,
              "status": 0,
              "tokens": "283/10",
              "delegator_shares": "35",
              "description": {
                "moniker": "Yakut",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1p56gv748xfd74qek5e637vhcr6tyjd9ukqfecc",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "NQX4yKpOztKrmgBhGIC5WOALOLOq3LTpbzsN4ZLXGec="
              },
              "revoked": false,
              "status": 2,
              "tokens": "990",
              "delegator_shares": "1100",
              "description": {
                "moniker": "Dokia-Interstellar-Cruiser",
                "identity": "4AD8231561095EBA",
                "website": "https://ion.dokia.capital",
                "details": "NodeOperator-Aurel (@awrelllRo)"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1pm0gx3lk46gf8vyj5my9w5tk06cuq66ll77ugj",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "j9be+ddLChInrFz6/820/uYh4WZBzlp61klyJBDy/ZY="
              },
              "revoked": false,
              "status": 2,
              "tokens": "100",
              "delegator_shares": "100",
              "description": {
                "moniker": "redbricks7000",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1pmntq5en2rgtr5rzr4e304efrve4lr43z32y5s",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "8pfpbIxBBiu88hpxS3CeRpv7kClEjl8SwVgckDNBGlE="
              },
              "revoked": false,
              "status": 2,
              "tokens": "850",
              "delegator_shares": "850",
              "description": {
                "moniker": "Staking Facilities Validator",
                "identity": "6B0DF6793DE1FB1F",
                "website": "http://stakingfacilities.com/cosmos",
                "details": "Hello Cosmos"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1p75wzde5z2ataefw65j3mtrm25kha779yfrvt6",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "Wz36ZKKPAL/yPW+d5mZ+xHb7w3QlAZCquplGKbGL+yk="
              },
              "revoked": false,
              "status": 0,
              "tokens": "333/5",
              "delegator_shares": "74",
              "description": {
                "moniker": "liangping",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1zpx36km7tk5cksyzxgvcp6g552p3uhwx84r53q",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "pPGLc4NhNaehdoV2antWuyr0GmBVEG1NhD9NiSRrTi0="
              },
              "revoked": false,
              "status": 2,
              "tokens": "1250",
              "delegator_shares": "1250",
              "description": {
                "moniker": "sikka.tech",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1ztc70khr9aghgwmuxap5y9ufavpxutjtkvhpe7",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "7vURd1EqPVfI/1viF/+LNm9lwaBAN/tSuoJ+Od972z0="
              },
              "revoked": false,
              "status": 0,
              "tokens": "90",
              "delegator_shares": "100",
              "description": {
                "moniker": "Project Bohr",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1zds6r7jyxyxcpf05r5yyyy3u8q2rvj9kkc6vcv",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "Go9GXHI6SCQo2QKMxkAkgYLhfo3XrVjWLR2nE2AvYyk="
              },
              "revoked": false,
              "status": 2,
              "tokens": "100",
              "delegator_shares": "100",
              "description": {
                "moniker": "Staked",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1zkupr83hrzkn3up5elktzcq3tuft8nxseu9x80",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "KTcVPCcGHjfCtHJF/8+Z6WnP8aF1KDZGP/17zjkQyok="
              },
              "revoked": false,
              "status": 2,
              "tokens": "150",
              "delegator_shares": "150",
              "description": {
                "moniker": "debt",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1rvm0em6w3qkzcwnzf9hkqvksujl895dfww4ecn",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "b52QQPpAO9Wa5qEmOr8ZPvLho0CgbBJx+WWxK+bpaYI="
              },
              "revoked": false,
              "status": 2,
              "tokens": "2265",
              "delegator_shares": "7550/3",
              "description": {
                "moniker": "wxvalidator0",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "actwo"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1rww4u22qqf2rhw02fww63t539d3ppeur3ywff3",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "QzxYamaSVNx+Idlb9bCcYsPI4Oa/zQ+rx4MN+XA2xLI="
              },
              "revoked": false,
              "status": 2,
              "tokens": "1065",
              "delegator_shares": "1065",
              "description": {
                "moniker": "Bison Trails",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1r05wuvemjzl25n3lspfku03e069cs5xsulfusn",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "tg8QOFWDiDL4dAlRnOoD1v6u9gOPFC2+rkf1YCm/SB4="
              },
              "revoked": false,
              "status": 0,
              "tokens": "41/10",
              "delegator_shares": "5",
              "description": {
                "moniker": "@fulltrust",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1rj74vaqm0xkxl5cjjam63mayh4x6at3m379ulv",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "ydjx2ea+PVuChrny6X2dluJwyXta+BsNQRsgHXp8fXw="
              },
              "revoked": false,
              "status": 2,
              "tokens": "150",
              "delegator_shares": "150",
              "description": {
                "moniker": "Mia",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1r4lzzzf683ec70fxfd9a3jmczxxkdxpfafz4ar",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "FSj9jMpCOdWfX0rCQ0TpnEa8/uOroQl4V4Bk3Msck3w="
              },
              "revoked": false,
              "status": 2,
              "tokens": "10",
              "delegator_shares": "10",
              "description": {
                "moniker": "trusted-validator",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1rcp29q3hpd246n6qak7jluqep4v006cdj2zse3",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "M2Iu/A8RQcmoiIEar/iV4wBjVj0I0gcd/nNnvwmnT2g="
              },
              "revoked": false,
              "status": 2,
              "tokens": "1000",
              "delegator_shares": "1000",
              "description": {
                "moniker": "layover run",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1yqpa5vhs8nq9mguyd2j5k5kn7dlqlgmjl4h67x",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "15LL69u+d5RdR8NA/5187ycW7SstSmrwiu1sk/JvhCU="
              },
              "revoked": false,
              "status": 0,
              "tokens": "9/5",
              "delegator_shares": "2",
              "description": {
                "moniker": "art_Fvalid",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1yfd9a4sqm8jm7lq932jc6fu8k86f394tgc2sku",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "w0Xbnzcm1tLGgpUrICpotJJXKDlmctoPz/JTBjzkIfk="
              },
              "revoked": false,
              "status": 2,
              "tokens": "10",
              "delegator_shares": "10",
              "description": {
                "moniker": "ATEAM-TT",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1yfseqtj5sjhzz2q2ym09jym4h4nc4yevae0jp2",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "eeImG09hOPo1W7j7lKepN/Lx6I9GGHqVBVEKmznxACc="
              },
              "revoked": false,
              "status": 2,
              "tokens": "1220",
              "delegator_shares": "1220",
              "description": {
                "moniker": "⚛️ melea-trust 🐞",
                "identity": "4BE49EABAA41B8BF",
                "website": "http://cosmos-trust.com/",
                "details": "Hola from Catalonia, Spain = to the Cosmos! Telegram user= (@Meleatrust)"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1y2z20pwqu5qpclque3pqkguruvheum2djtzjw3",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "A6GzeXUM3vsXaDAEYMSDgSKkqn9AoUYjs8empH46MGY="
              },
              "revoked": false,
              "status": 2,
              "tokens": "1280",
              "delegator_shares": "1280",
              "description": {
                "moniker": "Mythos",
                "identity": "2E9FDF34351A5112",
                "website": "https://mythos.services/",
                "details": "Validator run by Mythos"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1yd0rklq45zg89fywr89ccutlcwp9kehhh0z03k",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "Abws3eXrUFAH8LeZJIcECakPL945TTmFsBlXONOUeII="
              },
              "revoked": false,
              "status": 2,
              "tokens": "155",
              "delegator_shares": "155",
              "description": {
                "moniker": "Cosmodator-7000",
                "identity": "",
                "website": "https://cosmodator.io",
                "details": "Blockchain Cosmo Validator"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1y3fydat73q569cceaccpuytm3y4tdpv3n0fhdv",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "qgws5Mm+saA71C4cdQavr1bw0UytNqJcpwPDHR0k2hk="
              },
              "revoked": false,
              "status": 2,
              "tokens": "9",
              "delegator_shares": "10",
              "description": {
                "moniker": "EddyG-cVali",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1y32vvanvz88ajamugxrgl7vnvf0vtmjg5wxg7f",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "e9SYHFlpZz134w539AzQBlsw+Wk0N59eUkER6n/PcgU="
              },
              "revoked": false,
              "status": 0,
              "tokens": "18",
              "delegator_shares": "20",
              "description": {
                "moniker": "noroute",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1y6d5g4s75s9m8xemrs6d9am2cef4nm3ku0s5lv",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "+pZYup+40uW60mwWqLTxdjuu+wRXPq3c9Z+0fTDVvpg="
              },
              "revoked": false,
              "status": 0,
              "tokens": "369/10",
              "delegator_shares": "41",
              "description": {
                "moniker": "baryon.network",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1y7n3emnw6mkxutvs26rczynpvq84ddqu332hlv",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "oyA3rg/uuKNv4xASzi1tmmkn+FuVAstd+wZjDVYCirU="
              },
              "revoked": false,
              "status": 0,
              "tokens": "45",
              "delegator_shares": "50",
              "description": {
                "moniker": "kr_test",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1ylz6tqpgz0flpl6h5szmj6mmzyr598rq7heuye",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "+CUXpZTFWwT7jG4DCPezLBtv1+lAgdJd+36rRb7QnmM="
              },
              "revoked": false,
              "status": 0,
              "tokens": "180",
              "delegator_shares": "20000/81",
              "description": {
                "moniker": "Bajun",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr19q6y4mk3x668vm2jx0pu6e0jfy5a39jz0qwvxt",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "eNKen3yMRu2SDLEu3V+U2LXsyka3Na7sp99V7AuDP+s="
              },
              "revoked": false,
              "status": 0,
              "tokens": "243/5",
              "delegator_shares": "60",
              "description": {
                "moniker": "blackhole",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr19pkt3aeaprl8dt9x5c77u9unkrndfd3ze40tj9",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "BqlZxJkG3oUh7x14+BVpiXchKI9fbQ4G2llagRJithw="
              },
              "revoked": false,
              "status": 0,
              "tokens": "41/10",
              "delegator_shares": "5",
              "description": {
                "moniker": "agaoye",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr19rlhdmplhnmvy3jmcd835awask4l46nvp3pt4g",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "RpY+WLYnDD2KYvSvfqnSjsNE2592f55L61rdbcsXRDk="
              },
              "revoked": false,
              "status": 0,
              "tokens": "27",
              "delegator_shares": "30",
              "description": {
                "moniker": "ATEAM-T",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr19dfkpwfj3mls96ny2ulgq7cf0dh5xs5a05pn6e",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "ncrdAm41M/iC+6ROlnGvtwYZ2UHioDeQl2Hc6okEDew="
              },
              "revoked": false,
              "status": 0,
              "tokens": "27/10",
              "delegator_shares": "3",
              "description": {
                "moniker": "Forbole-1",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr193vn0gk3nsmjkxwz78gce8e8mkmagmvulpg5jt",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "tOEqjO2t51PEgO9Tv0B7qM0yPmy1n5tMa3Beg0tp3ns="
              },
              "revoked": false,
              "status": 0,
              "tokens": "8271/10",
              "delegator_shares": "9190/9",
              "description": {
                "moniker": "divislight",
                "identity": "6A0D65E29A4CBC8E",
                "website": "https://jackzampolin.com",
                "details": "Its all about the signatures baby"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr19n4s8fx9hr0j2dttr4fg48wr7kpxxste7evjzk",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "dpYEHcQAo439YIswYnm3ytxyoP7RtorAS59+FOXiOKQ="
              },
              "revoked": false,
              "status": 2,
              "tokens": "230",
              "delegator_shares": "230",
              "description": {
                "moniker": "new_vali",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr19ncpp07tuqqgu2kjp4wq2zlt3s0q5f9ywvhwuw",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "UQHY8OYNqCwDkAm7bn+X/VqV6p3PHTRqkQGEsqnpdws="
              },
              "revoked": false,
              "status": 2,
              "tokens": "110",
              "delegator_shares": "110",
              "description": {
                "moniker": "kreios-2",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr195qlnqc2mxsxnfvcljqr5ssfjjymn0usk35v75",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "92HX8Xt4jN7KFMnkOCoSshw2ihCQEE+ers9EvGKhxL4="
              },
              "revoked": false,
              "status": 0,
              "tokens": "41/10",
              "delegator_shares": "5",
              "description": {
                "moniker": "lowpeople",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr194pux2wkhuusey6x4wwaqm0wh26cleh06u28lr",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "ocujtNWoFsmh22NmGX8U7X0uNoRuUg2L0/lMZIGIQbE="
              },
              "revoked": false,
              "status": 0,
              "tokens": "203/10",
              "delegator_shares": "25",
              "description": {
                "moniker": "baryon.network",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr19h5dc2c76h4ryvl9zl9373mk52ka77vy5g48gs",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "cJv3rBtHyIGTsVQ0ZWXdjIL/dBQwn8N9ml2aT8e3/jM="
              },
              "revoked": false,
              "status": 0,
              "tokens": "117/10",
              "delegator_shares": "13",
              "description": {
                "moniker": "gridworkz",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr19uhnhct0p45ek6qxp3cjjrjtz4pacwcsgzvpuj",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "lUMCRAeu47BsOhNvCQTQJQeB68z0/VaElC9j5gDt9y8="
              },
              "revoked": false,
              "status": 2,
              "tokens": "1501",
              "delegator_shares": "1501",
              "description": {
                "moniker": "vhxnode1",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1xvzw5nz4mp5s5cat7rancc248v5wpecrlqtw6n",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "ilDFePqWkNRNWUEoWKq42jZC5CD4rglfs7+gCuq6Qc8="
              },
              "revoked": false,
              "status": 2,
              "tokens": "151",
              "delegator_shares": "151",
              "description": {
                "moniker": "lino",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1xvt4e7xd0j9dwv2w83g50tpcltsl90h5dfnz6h",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "xlO2cnii42KisAn8OcstC/3XV5+I0FlcSbWuyy5MVA8="
              },
              "revoked": false,
              "status": 2,
              "tokens": "127",
              "delegator_shares": "127",
              "description": {
                "moniker": "21e800",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1xupfqk73y7rmc6qdgv7rtjy8wsngvt2g2t59t3",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "XQDVMXja3kFk5Jb47BsqJmzcDsM4lE9+r+f/J3O5Jms="
              },
              "revoked": false,
              "status": 2,
              "tokens": "135",
              "delegator_shares": "150",
              "description": {
                "moniker": "inschain_validator",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1xuwcmdvef5xudnyqa9gh5vqlq6ff0meu87zjxh",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "FiEl7a96NaaDaVH5mV9EK/6tOUWL5mkkT46ARRKAWIg="
              },
              "revoked": false,
              "status": 2,
              "tokens": "1681",
              "delegator_shares": "1681",
              "description": {
                "moniker": "chorusone_validator",
                "identity": "",
                "website": "https://www.chorus.one",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1xa8z3sza70wrt06yfhlxq0ujg6fa7ld8h7m6xz",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "LRK6flgMsHrBYkIIahegEkzlwC7gQ3dAyw0jcAd1u/k="
              },
              "revoked": false,
              "status": 2,
              "tokens": "505",
              "delegator_shares": "505",
              "description": {
                "moniker": "@fulltrust",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1xlkah6d2pcr8rchunmn6qdmdptr02844eee6le",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "MfyCioP9/yGlqUcwnL/ENi23RKPBm1O1wum1DUUUYT8="
              },
              "revoked": false,
              "status": 2,
              "tokens": "56",
              "delegator_shares": "56",
              "description": {
                "moniker": "Bity Validator",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr18zwvgsrl3gc8y0fhzdev378u9gn37d9024njy0",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "ISIM341M+EUYOMpwhn9gyCvRUZ06xomp/+lOwa50meQ="
              },
              "revoked": false,
              "status": 2,
              "tokens": "800",
              "delegator_shares": "8000/9",
              "description": {
                "moniker": "corona02",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr18z3pnjgdtt337z8g5drfts7r3fm6n9a0896h0r",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "nsofE1FmSr1TiDR0gfnxfMDQ8o2pC+1NE7Oa9ceztSg="
              },
              "revoked": false,
              "status": 0,
              "tokens": "342",
              "delegator_shares": "380",
              "description": {
                "moniker": "coinone",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr18raps2wmm63n2gldsrqc5x5z023p559l7r5vv6",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "DrgnMZqdtzRbDvDGh5290M1QtAfem1VHaY+zhq/nBWI="
              },
              "revoked": false,
              "status": 2,
              "tokens": "50",
              "delegator_shares": "50",
              "description": {
                "moniker": "FennecFox",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1825jrcgu3keurjlnpsujx9y8lrzgddjfsmax0n",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "VSz12Gh6kqPKSlwrZZQ0QZsXvRadOvAeDhmuboV/Lqc="
              },
              "revoked": false,
              "status": 2,
              "tokens": "281",
              "delegator_shares": "2810/9",
              "description": {
                "moniker": "412",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr182ujqw3r8p5fffjqkf0rxzj29pg5q96nxd2khq",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "VakMQSPBEuSC9Nwuv8WWhrZVUmH31bUR4+G6pJhkgE8="
              },
              "revoked": false,
              "status": 2,
              "tokens": "1773",
              "delegator_shares": "1773",
              "description": {
                "moniker": "Umbrella",
                "identity": "A530AC4D75991FE2",
                "website": "",
                "details": "Umbrella Validator"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr18thamkhnj9wz8pa4nhnp9rldprgant57ryzag7",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "lB+CgSbTzpHXJHUU15fJIPaPl6XF0nSfdNxBnlKJqTk="
              },
              "revoked": false,
              "status": 2,
              "tokens": "18291/10",
              "delegator_shares": "9463951/4191",
              "description": {
                "moniker": "delegate-now",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr185n9jln2g4w79w39m0g85x2tns6uekqg99tq7u",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "ViJzrw9Unk1XyDg4f0UQLFE4zfUaFh23wZZY0f31Qg0="
              },
              "revoked": false,
              "status": 2,
              "tokens": "1200",
              "delegator_shares": "1200",
              "description": {
                "moniker": "CosCloud",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr18m09d56pg5p2660de4sjfezpd8ud6jfghndfnt",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "RwPRoiY5C0covekqbr3VrQwxWGHioUUIf2+TOq8LIC0="
              },
              "revoked": false,
              "status": 2,
              "tokens": "135",
              "delegator_shares": "150",
              "description": {
                "moniker": "starfish",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1gq0qecxs8xdaarrqxxazwavwxm7qz5jzs5anvt",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "VBVHOLnWGptY26J0wqXoZI2Dnu96pccMb08zlsaxPCQ="
              },
              "revoked": false,
              "status": 2,
              "tokens": "1100",
              "delegator_shares": "1100",
              "description": {
                "moniker": "certus.one",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1g8xwmm3tp7pgdrmjgt4fm7d26anrrnle4yavuw",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "5GAxIHzu06l0n3MU8wNQrCcrrnpZR9EKe5qZMsM2h/E="
              },
              "revoked": false,
              "status": 2,
              "tokens": "100",
              "delegator_shares": "100",
              "description": {
                "moniker": "ip-10-0-3-99",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1g6sc5t9t68vmcj3alk7dfqr54tvastpxac28k6",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "kol7Gj60Fct4X8T1rHLJQ0z/b14UqqSae8h1e37rLL8="
              },
              "revoked": false,
              "status": 2,
              "tokens": "90",
              "delegator_shares": "100",
              "description": {
                "moniker": "VNode01",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1fyevhdmxpkz8stjxggv9wejf9caed02vyp8y5p",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "v/tjGnK7dQ2Ek+/8q3/fxVlKucDnK8SVukLlXLVK29U="
              },
              "revoked": false,
              "status": 0,
              "tokens": "9",
              "delegator_shares": "10",
              "description": {
                "moniker": "threez",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1ftrggumlvkesqny90par3gc9lp70yqhtlqucsk",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "T3ITN3B9p3HuIo4m8G3zzLmmR+JmPyAat+L2eYG3/sw="
              },
              "revoked": false,
              "status": 0,
              "tokens": "2106",
              "delegator_shares": "2340",
              "description": {
                "moniker": "bdongerNext",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1fskmzyt2hr8usr3s65dq3rfur3fy6g2hjp23tm",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "N3K5kDdfcKJurfaa6s2zfKgtYvz1Pagz7VWi9ZfX8yM="
              },
              "revoked": false,
              "status": 2,
              "tokens": "1252",
              "delegator_shares": "1252",
              "description": {
                "moniker": "ATEAM1",
                "identity": "E503F770E04D5684",
                "website": "https://nodeateam.com",
                "details": "ATEAM1"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1fhfcs5npydv8s96wrk9p5ychctslu92t4n5qd4",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "4JoJuRfaANhdM1x3AWRo1/Cj9DH3VA+fi1SynzknV+w="
              },
              "revoked": false,
              "status": 2,
              "tokens": "150",
              "delegator_shares": "150",
              "description": {
                "moniker": "davinchcode",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr12qprqyxtxxlu7unncd9rrhmer5emg2pmrw27qk",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "E0KYUg0eXCnwZ1GL2vZ0214HGqbydUQ7sowpwS5VpXM="
              },
              "revoked": false,
              "status": 0,
              "tokens": "243/10",
              "delegator_shares": "30",
              "description": {
                "moniker": "cm-gaia-1",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr12pn4p6r5wjpsep9kn655ng7fh59yez7t0rahru",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "y7p9JSVZBnRxjAI9v5Pxl37hMtyuHf6B4Ghqzm6+ii0="
              },
              "revoked": false,
              "status": 2,
              "tokens": "100",
              "delegator_shares": "100",
              "description": {
                "moniker": "oleary-labs",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr128ty3kzhcepjacu4q0xjgq60qa3zz8na3jl793",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "R/3f7VruxWpu+2hiHlVpplTwoOou5kfQI1k/6/9H/y8="
              },
              "revoked": false,
              "status": 2,
              "tokens": "152",
              "delegator_shares": "152",
              "description": {
                "moniker": "stake.zone",
                "identity": "0A888728046018EC",
                "website": "http://stake.zone",
                "details": "operated by nuevax"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr12fm2n7pkmtqk7untyhs9my749jyzef3vhne95q",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "aSbdZXtmYI31dCcZ/OgyL1O6YltrDqmkSyuhLIiIKxQ="
              },
              "revoked": false,
              "status": 2,
              "tokens": "1998",
              "delegator_shares": "1998",
              "description": {
                "moniker": "codetree",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "codetree gaia node"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr120skmenn2a0ra8y6zassrxvnfc5rlme8rqarvs",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "dnFjFoTM9sP/RjQkXBK1YpYn3v5W+j0+g/OfUHS4xu8="
              },
              "revoked": false,
              "status": 2,
              "tokens": "55",
              "delegator_shares": "55",
              "description": {
                "moniker": "aether",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr12ceualfg92x7du73mtcv0zya4nxvq3tl2m52uz",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "omAzuJps8KX3/iOC1LjwkMPMH3c6tjfLXwCNWXRBdWw="
              },
              "revoked": false,
              "status": 2,
              "tokens": "1455",
              "delegator_shares": "1455",
              "description": {
                "moniker": "Adrian Brink - Cryptium Labs",
                "identity": "F61053D3FBD06353",
                "website": "https://cryptium.ch",
                "details": "Secure and Available Validation out of the Swiss Alps brought to you by Cryptium Labs; @adrianbrink, @awasunyin, @cwgoes."
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr126ayk3hse5zvk9gxfmpsjr9565ef72pv9g20yx",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "BaaCxmYHKJ6obIzTCdRtjw1cc8d2mUJcMbLWCjf1aLo="
              },
              "revoked": false,
              "status": 2,
              "tokens": "33593/10",
              "delegator_shares": "302672930/65691",
              "description": {
                "moniker": "grass-fed",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1t9rqvvys4zh300x93a6fr3h064gp9y8al7m4en",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "31942iQRr+ajHMCklJslIbPzQl6OplvluWuXUnHlVQ0="
              },
              "revoked": false,
              "status": 2,
              "tokens": "606",
              "delegator_shares": "606",
              "description": {
                "moniker": "newaether",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1t9nxlu834au3u7mltyf8udzw7kzsf0wntx7e4w",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "JH/rHJsE3h137vSNSsmY1X/5NU/1BkgPhahSsmSgp+Y="
              },
              "revoked": false,
              "status": 0,
              "tokens": "8/5",
              "delegator_shares": "2",
              "description": {
                "moniker": "spark-test",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1tsysfw53f36ap440e5ljf6hnc3ueuwal9vmcx4",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "QRDgIC3dwcj0pqN95IShVJ5aNrnV/nvr19KkCr6koWo="
              },
              "revoked": false,
              "status": 0,
              "tokens": "9/5",
              "delegator_shares": "2",
              "description": {
                "moniker": "spark-yuan",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1t3afuz2kt99jz4pe5k4vjvkdmldn2x0lqzv83w",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "gJDxhwAE6GeGCKQeVaNZ5is7+7MFHXtOG0UsnguKdoA="
              },
              "revoked": false,
              "status": 2,
              "tokens": "300",
              "delegator_shares": "300",
              "description": {
                "moniker": "BlissDynamics",
                "identity": "2A60089536521E12",
                "website": "https://cosmos.network",
                "details": "Prosper with Bliss Dynamics"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1t3lvwnp0rmfyc96c0qh6qa67nplpdhytr5dd59",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "GpHZ0Hr3C1hvVMkcDe/76lvWMjw5SF0JBI/D6gDkQTo="
              },
              "revoked": false,
              "status": 2,
              "tokens": "250",
              "delegator_shares": "250",
              "description": {
                "moniker": "StakeStack",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1vrc7zpg5teawwuzkfh6t7c6sy353sukhlarxpa",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "Epd2FDKZwDybzT38Z7WB0y2jiLn9/2OLzmY3Zu18l6I="
              },
              "revoked": false,
              "status": 2,
              "tokens": "1257",
              "delegator_shares": "1257",
              "description": {
                "moniker": "figment",
                "identity": "",
                "website": "https://figment.network",
                "details": "Figment Networks Inc."
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1v4d9c3wy6n7h4udj4lkumc0l99x95cml44uyv8",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "LUYsagcVciSK95n6WEIbKhsTKD4tgAbBz7kND3ZBpVo="
              },
              "revoked": false,
              "status": 0,
              "tokens": "13071/20",
              "delegator_shares": "1005",
              "description": {
                "moniker": "suyuuuuu",
                "identity": "D86942F16972FFE7",
                "website": "https://en.wikipedia.org/wiki/Istanbul",
                "details": "Istanbul"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 3,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1vk5dwhn8m3sht9xrartj9kuv098t00spkqphcd",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "t07LChLn+eoTivIraUyU0M/T4zkhY6TZkn8cAWtOLfU="
              },
              "revoked": false,
              "status": 2,
              "tokens": "18",
              "delegator_shares": "18",
              "description": {
                "moniker": "egon validator",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1dq7jt3pn7mrce2twac907jue7hjd0p0rgt3qnq",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "uEWWDBwFW+/BpTCvNCLW7AP98hndBukzSbrwCb7sooo="
              },
              "revoked": false,
              "status": 0,
              "tokens": "11/5",
              "delegator_shares": "3",
              "description": {
                "moniker": "spark",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1dkm8z36m96k8aukg5rxnfpw46dt5mnmvsfuze2",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "tkdOxR2CIOctyc5e9IMGwnMwB420OTN3Rfb6fUMVhr8="
              },
              "revoked": false,
              "status": 2,
              "tokens": "5",
              "delegator_shares": "5",
              "description": {
                "moniker": "Mythos(private)",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1dl9mu7uya6mfsd80wyw2rgjmwjtk5ajjf8mxcz",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "dj4b9d19oYvNJP1WMHJ1g5tzzMGMNxRnxrtCXBF6jsc="
              },
              "revoked": false,
              "status": 2,
              "tokens": "30",
              "delegator_shares": "30",
              "description": {
                "moniker": "Liviu",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1dl0hk7064k5wp43zkllmd0p9pgu2mnra898gyh",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "ParbWCAIUA/JGPjgapGrRhYzKy6FQ1HpaJcAQwGGn1s="
              },
              "revoked": false,
              "status": 0,
              "tokens": "45",
              "delegator_shares": "50",
              "description": {
                "moniker": "EU_test",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1wf3sncgk7s2ykamrhy4etf09y94rrrg4k73wwr",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "1UyFY3YJja8vrmDc/ZqQFo7ORvLdLK2X2wYeCUBZzJc="
              },
              "revoked": false,
              "status": 2,
              "tokens": "409/2",
              "delegator_shares": "2045/9",
              "description": {
                "moniker": "clawmvp.xyz",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1w2j930e3ueg8f365ff2t92m5lnp856fe47f3hy",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "P8GO78Sebgek8Zh9bnnsl1lVycl0dV7g3BufsXPEipQ="
              },
              "revoked": false,
              "status": 0,
              "tokens": "53/10",
              "delegator_shares": "10",
              "description": {
                "moniker": "chainflow09",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1wnwr086pvkv2pqa5khznunm93ac52u677zgqe6",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "zaRqlLb7T6OaFRImEZF1eIuzoRpDpVyvBs3ITbPuPiw="
              },
              "revoked": false,
              "status": 2,
              "tokens": "30",
              "delegator_shares": "30",
              "description": {
                "moniker": "satoshi-ledger",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1wk0t6na230vxhf6ddresw2c40k5y3ayrww0s2m",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "UjTvuOew2EaooduJBiYmBWeF5ai0yFJG8uio5YXpJgg="
              },
              "revoked": false,
              "status": 2,
              "tokens": "135",
              "delegator_shares": "150",
              "description": {
                "moniker": "P2P.ORG Validator",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1whxd48da3r56n8eecym8zg0c6xmf35fn2myart",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "AJR2ex094A1nJEWQsZYjALWsrSl1/huGZ37z2ZsMrpg="
              },
              "revoked": false,
              "status": 2,
              "tokens": "1100",
              "delegator_shares": "1100",
              "description": {
                "moniker": "dev",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1whd2u62u3zk7w09yuq4cz97fueq8e96qp69xz0",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "t6nC3SGVoHDhRGw0xNpBHsD9GIcTs+1aCXkXjWXYXOY="
              },
              "revoked": false,
              "status": 0,
              "tokens": "18",
              "delegator_shares": "20",
              "description": {
                "moniker": "liangping-bj",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1wesup293xcfjze34u98j05ffu04cssy32gjv3u",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "HQA+mntabk10g6Us3YQ1I9U4jfv9LPGGXMmTCVbPB+Y="
              },
              "revoked": false,
              "status": 0,
              "tokens": "99/10",
              "delegator_shares": "11",
              "description": {
                "moniker": "teamSoNicX",
                "identity": "54D70B791F6F9035",
                "website": "https://sites.google.com/view/sonicx",
                "details": "ⓣSNX_teamSoNicX"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr100ztxxmwkqr44jts3a543xhme7vjychzyagz3g",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "Oxp5JKfzexRWJravPZ9XAEF2ydyzh04362Eo4WloEBE="
              },
              "revoked": false,
              "status": 2,
              "tokens": "31",
              "delegator_shares": "31",
              "description": {
                "moniker": "Yakut",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr10505nl7yftsme9jk2glhjhta7w0475uva87paj",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "2p8s/pRBZPjYWKKMlR7AOXypDzDmPo762iXlKpCwtco="
              },
              "revoked": false,
              "status": 0,
              "tokens": "1818/5",
              "delegator_shares": "4040/9",
              "description": {
                "moniker": "Forbole",
                "identity": "4A5D9C100A76D9A8",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr106ccejx2h7hwzj9xxaz3hju22yyq5mfy4rydav",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "KfZeww1iQ6YnslfhtZ0qi8yVy0PQHj7/j0Rrkk2gz5o="
              },
              "revoked": false,
              "status": 2,
              "tokens": "110",
              "delegator_shares": "110",
              "description": {
                "moniker": "bharvest-validator(USW1)",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr10mfxvxjga3rxp5w93rx8jdxtt68j5qk0wv02hn",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "i5w40hJcCfKwigRBYtlI5Lcej4YAosroQbREZq2Xla0="
              },
              "revoked": false,
              "status": 0,
              "tokens": "9",
              "delegator_shares": "10",
              "description": {
                "moniker": "lowpeople2",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr10atd0mjtdt7uzk0jgr32c32semnh3xvu4nx77p",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "xIS+GYtNUOsWEbVq1BtXciEVRVVRDVJogfuyauE0H90="
              },
              "revoked": false,
              "status": 0,
              "tokens": "9/2",
              "delegator_shares": "5",
              "description": {
                "moniker": "deadlock",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1syhzpjwx6vv3erv2myq7txrjhrp95hrhgcl242",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "k3YLQYEN2QMP6XITRsBmgb+pNGhJ5Jbg0bzUW977kK0="
              },
              "revoked": false,
              "status": 2,
              "tokens": "90",
              "delegator_shares": "90",
              "description": {
                "moniker": "TropicalMongoX",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1st4262sfqk7jqdwp2s7gk2umm5yhexhjrew53z",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "Bawa9CE/IGarJDSWdA7gXLJjBxtnCWk7BX94H3Zyhd0="
              },
              "revoked": false,
              "status": 0,
              "tokens": "9",
              "delegator_shares": "10",
              "description": {
                "moniker": "bharvest-validator(FF1)",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1shuqhpl273t96yg6nnqvyfeewj3ew3mdcwvcnu",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "2JoNf1gavJ1d6XFIumO1Mki5GVMOcg58AioHksU3maE="
              },
              "revoked": false,
              "status": 2,
              "tokens": "1130",
              "delegator_shares": "1130",
              "description": {
                "moniker": "bianjie",
                "identity": "B536B81DDA4323BB",
                "website": "https://bianjie.ai",
                "details": "Power house from China"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr13q0fl9qakvt5zkh6qurlrur4xq7gullkv9f5zq",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "6lTjnB9szF+z4c0XwBrplBjK6GaaMUAcFJ/8VHwzOuE="
              },
              "revoked": false,
              "status": 2,
              "tokens": "10",
              "delegator_shares": "10",
              "description": {
                "moniker": "ATEAM-H",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr13psmeywshfyfcg3tr5l4qm968smtkcu8um0nqn",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "Kv+wY9i0uMGO8ukF4o+Za5uuUp/S8o9N4C1rdA22h8M="
              },
              "revoked": false,
              "status": 2,
              "tokens": "10",
              "delegator_shares": "10",
              "description": {
                "moniker": "louis",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr13rt2suxfgefcquy9pveuc5uxfayy5rscr4vld7",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "UZlobYgzLP7gOVwBtMRo5S3yrMPlZWIXYEwXCs7dhIA="
              },
              "revoked": false,
              "status": 2,
              "tokens": "359",
              "delegator_shares": "359",
              "description": {
                "moniker": "kreios",
                "identity": "A0D6EFF27E1AA89D",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr13fw4ff9u2qftsjj39mda54ufrpjs77q0ke6m49",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "+Uc1tDCsc8rM+mPQ79C8Mp3Y7ajlL2Lb5AjYboPqlQE="
              },
              "revoked": false,
              "status": 2,
              "tokens": "5",
              "delegator_shares": "5",
              "description": {
                "moniker": "chungkueiblock",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1j9w65cmt5s5kw9430jsv87yd4dgmmq6w60f3tl",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "GTnDCQ1qvYXfljMQD9ekCruCnVzY5Vfhww/YNCKkmFA="
              },
              "revoked": false,
              "status": 0,
              "tokens": "4/5",
              "delegator_shares": "1",
              "description": {
                "moniker": "jst002",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1jtt6al0acr8h0uvq489rt9zx9lnau7rlcu30pt",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "L0I4JoDfktbDWe0fCDL/nQlBPkF5mNgqamnM5JKJ1Uc="
              },
              "revoked": false,
              "status": 0,
              "tokens": "135",
              "delegator_shares": "150",
              "description": {
                "moniker": "JCol",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1jt77t00gucffqhrq7lh6vwmwzrwgw2fpvnhwk7",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "kCfC+M4tOz3Js2IFJw331BkWcW+cdpSJXFRkSakmFZU="
              },
              "revoked": false,
              "status": 2,
              "tokens": "81/2",
              "delegator_shares": "45",
              "description": {
                "moniker": "INVIERNO",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1jdrxh34d00rm8ydclhxgkzmgp55y67m54umags",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "HbzTC9Xsz7IdhiqRImwpzVlOJ37xCElJ5yKd9hcXSYs="
              },
              "revoked": false,
              "status": 0,
              "tokens": "9",
              "delegator_shares": "10",
              "description": {
                "moniker": "baryon.network-x",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1jw554408yw2h438q200jyuqgln76xh4ax0q4s0",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "SNwrT1B+A4g6TY7x0QzVrmVbcbl3cHXzXdD1tFHxLNo="
              },
              "revoked": false,
              "status": 0,
              "tokens": "81",
              "delegator_shares": "100",
              "description": {
                "moniker": "TruNode",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1jsch8k385shvse6j5dfx20483qym5uhq76xpjf",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "QEMDwUVoyJT7MNfOYKa25xU+Lnsz/ciH8rFUri4diLI="
              },
              "revoked": false,
              "status": 2,
              "tokens": "90",
              "delegator_shares": "100",
              "description": {
                "moniker": "SVNode01",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1jkmn6asju47zuuzrf8rjt7sllaj5cx4kueyv8p",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "LSDd6ab46sHxwJSrg5YLpsPG2o6EcsZ3rDikpHzMNmI="
              },
              "revoked": false,
              "status": 2,
              "tokens": "210",
              "delegator_shares": "210",
              "description": {
                "moniker": "w1m3l",
                "identity": "CB48E188D37C680D",
                "website": "https://wimelpos.com",
                "details": "From south of Spain, Seville"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1jh3grawl62juvx5p8fz5xsy9hpw8w3mngqafe4",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "0HqB2x6x5HzeozpHatePECw07x1UcDdSz8kQGNznnA8="
              },
              "revoked": false,
              "status": 2,
              "tokens": "1001",
              "delegator_shares": "1001",
              "description": {
                "moniker": "block3.community",
                "identity": "EC7A1E9392E0F7CA",
                "website": "http://block3.community/",
                "details": "The best way to learn, is to build."
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1jck6gp4mqy33sk6a0fr5c8maq53hf4245v3mgg",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "UBTju7UZfXLVPPYb1a8gPZ69BeCv2Fho7YVo2EUbxKc="
              },
              "revoked": false,
              "status": 2,
              "tokens": "135",
              "delegator_shares": "150",
              "description": {
                "moniker": "7768",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1je98ejgpvv2uzku3t6h5cydyr9hzva2cqhf3v5",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "yvHIU57th/log9naiq3Y5AKqDDTbsGNG1SRhp5p4Go8="
              },
              "revoked": false,
              "status": 0,
              "tokens": "9/2",
              "delegator_shares": "5",
              "description": {
                "moniker": "RockChain",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1ja98tyxchsde6v3l02thp4fc35eymg88kts070",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "wdUJrATbIN7C9pgzDzaBZ4WANQzQL+/r+36yXqDdeCQ="
              },
              "revoked": false,
              "status": 0,
              "tokens": "99/10",
              "delegator_shares": "11",
              "description": {
                "moniker": "fullnode04",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1nnyel6v0kx8zmfh9edmre3ua4dt9306cfxsxgd",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "q5ezDn4DcWFPWvMayPJI35nXr//jjF8fGHsuiHjpDcU="
              },
              "revoked": false,
              "status": 0,
              "tokens": "90",
              "delegator_shares": "100",
              "description": {
                "moniker": "spptest1",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1na3wr7ujdp3qdej6l5y0k4znzrkaz77t2yjaqf",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "jj0Y/Fy8JSJR3g+PHU6Ce0ecYwHGUVJ4bVyR7WwcyLI="
              },
              "revoked": false,
              "status": 2,
              "tokens": "1151",
              "delegator_shares": "1151",
              "description": {
                "moniker": "nuevax",
                "identity": "0A888728046018EC",
                "website": "http://stake.zone",
                "details": "operated by nuevax"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr15rrxdflhq4744qt9m5ysstq3zpykhacf908eez",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "oG8Q5o+SN4wqMLvlIfVgQPnsQzNEKeH0D/XGM8JlGrY="
              },
              "revoked": false,
              "status": 2,
              "tokens": "2512",
              "delegator_shares": "2512",
              "description": {
                "moniker": "wancloud",
                "identity": "D185B18E845FEA40",
                "website": "https://www.wancloud.io",
                "details": "wanxiang blockchain - wancloud"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr15w2rengajq9js8hu57kjw88dly5vy7gsqedn0n",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "Xi7nIgj4PqVXrpKLfJhcyxyVY1d3HRo72sKKPDmuU78="
              },
              "revoked": false,
              "status": 2,
              "tokens": "150",
              "delegator_shares": "150",
              "description": {
                "moniker": "kittyfish",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr15wkl5rq6uk8qdnpehdkp7clng255e06z2p9lfk",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "JFh7R/vhQU/SW43OPRcFtPKdCHowiYrISAl+V+bv3V8="
              },
              "revoked": false,
              "status": 0,
              "tokens": "18",
              "delegator_shares": "20",
              "description": {
                "moniker": "vps",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr15mzck4rm5w55sfgr9kgexgtna22v608907t8y6",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "G1F5j4VjmFrlOBHtqqOAa7q2hMfUm6lnzctnVFhROZ0="
              },
              "revoked": false,
              "status": 2,
              "tokens": "150",
              "delegator_shares": "150",
              "description": {
                "moniker": "⚛️ cosmos-trust.com 🐞 🎬  🌈 🦄 🌟 🚀 ",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr15u9ve7fz8fqaf7r2u3p4f9ru4ze47pau5cxgcg",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "3wRufybSUsTMnUeQkP74uJNDRKeM8jBLAS64T0BRfpY="
              },
              "revoked": false,
              "status": 2,
              "tokens": "180",
              "delegator_shares": "200",
              "description": {
                "moniker": "mpaxeNode",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr157eu0k3fvnk4d3rpgtyz2jujf8shlk8mxeqjkk",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "EaeIIAIlSnmV19W5TfouZ3Kuosamdbjd9+n7gmPTJF4="
              },
              "revoked": false,
              "status": 0,
              "tokens": "47/2",
              "delegator_shares": "290/9",
              "description": {
                "moniker": "Stone",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr157mg9hnhchfrqvk3enrvmvj29yhmlwf759xrgw",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "MRCeBDANSjH6IsxO0z6tRe+xqoZvIGhdfl1t+SXGUpM="
              },
              "revoked": false,
              "status": 2,
              "tokens": "1160",
              "delegator_shares": "1160",
              "description": {
                "moniker": "jlandrews",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr14zfph0h8rexsca6gg6jkwqup3sgl6mwj6eu4e6",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "b9RSkt+WmMxVHQExQH0IMPpnR9zDAaJwz/mv1gtyRVY="
              },
              "revoked": false,
              "status": 2,
              "tokens": "81",
              "delegator_shares": "100",
              "description": {
                "moniker": "smartpesa",
                "identity": "2D5AA38DB12E96D1",
                "website": "https://www.smartpesa.io",
                "details": "Thorsten Neumann"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr14sqp5mqtmfjdk9khae6lkfytwj9crtunpzr8f7",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "cAbKlLG8cvAuUoyTvCCzTP5w2MCqnjUgjR1FyZihz30="
              },
              "revoked": false,
              "status": 0,
              "tokens": "2021/10",
              "delegator_shares": "2245/9",
              "description": {
                "moniker": "shensi",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr14e774gfzt5l9ka766ehfgu6n5hgy9f3sehzyv8",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "bZSEBDNIOr0xJ/PxaAScJIyG6hqFtryBAMNwghAOTTU="
              },
              "revoked": false,
              "status": 0,
              "tokens": "81",
              "delegator_shares": "100",
              "description": {
                "moniker": "cwgoes",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1k9pqxd8fxqlk52uwfxnlsexqj6xnmw5swhss45",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "2qtEBT+Tc+SD2wJsdrVMHXrBKfvesxtmtSKDK5fXwA0="
              },
              "revoked": false,
              "status": 2,
              "tokens": "144",
              "delegator_shares": "1600/9",
              "description": {
                "moniker": "jjangg96",
                "identity": "3EAA521EDB813768",
                "website": "",
                "details": "BigFan of Cosmos"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1kwftgznylzerft2ksgkzvvfn5rfpy4nk2ye8na",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "BB4a/Xh5z+dkGCRlF+pSGC3iDOoDrFse/xzQAtmxMF4="
              },
              "revoked": false,
              "status": 2,
              "tokens": "1149",
              "delegator_shares": "1149",
              "description": {
                "moniker": "chainflow08",
                "identity": "81D443FA08A4A926",
                "website": "https://chainflow.io/cosmos-validators-intro",
                "details": "#BUIDL'ing Chainflow Cosmos Validators for the Community"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1k5ah4mpy8ykzp2mpcvapnyv6gemp35zuyhw8sy",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "PV9ZicPQGtOsRT3UnuY+4YGbHUIHwaVGc1JsHhcOM/c="
              },
              "revoked": false,
              "status": 0,
              "tokens": "36",
              "delegator_shares": "40",
              "description": {
                "moniker": "chorus_test",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1hydk58dtxr87ml34xqc07ph2h66yt2f0rrhyt8",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "aunzA3JTNtbn4EW3t914GMHjogqHQV6vvsm3Z7dJ07o="
              },
              "revoked": false,
              "status": 0,
              "tokens": "45",
              "delegator_shares": "50",
              "description": {
                "moniker": "tokyo_test",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1hfygmryre7r8m9pfqmc85y7kw7ejphmp59k2x8",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "66j9af4xDJSblMLS+mFbp7d8TaFGu0FOo+0MwEYm2lE="
              },
              "revoked": false,
              "status": 2,
              "tokens": "12099/10",
              "delegator_shares": "4033/3",
              "description": {
                "moniker": "bkcm",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1h4q0xkdg30cl9vw0u8ejm0rs337dszy98gnd4a",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "2gOiMAdnIdND4cA75E7naQdyyIYDAdcjF3uO6OiEZlU="
              },
              "revoked": false,
              "status": 0,
              "tokens": "1191/10",
              "delegator_shares": "147",
              "description": {
                "moniker": "joltz-secureware.io",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1h6el6v9rn0unkw2fh9s86juvg6fsvcmy8wy72s",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "G5xqaI1M/ckd6k5tLzWkGF/MdGZ4/xmJThP4AjksrFk="
              },
              "revoked": false,
              "status": 2,
              "tokens": "29",
              "delegator_shares": "29",
              "description": {
                "moniker": "basement",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1hm5f52equcd4f268uxf4xhwljpnljvdyqal3dv",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "CuPin/hLM2tK6KvcUlOi/Brm5xNi3zq6UPgpLFPSzt8="
              },
              "revoked": false,
              "status": 2,
              "tokens": "281/10",
              "delegator_shares": "347/9",
              "description": {
                "moniker": "NodeBreaker",
                "identity": "89D35987D8289AAC",
                "website": "https://1q2.me",
                "details": "I will inspect all the validator's security for cosmos space."
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1hmkg328mkqen9kc7z3j07hhu5ep9qh64q8kkfh",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "yW/UYA7l6xOwS6/Upa4cZCiaPPR8m7ih9NL4LDj3kt0="
              },
              "revoked": false,
              "status": 0,
              "tokens": "99/10",
              "delegator_shares": "11",
              "description": {
                "moniker": "TEAM_Korea_No.1",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1h7yr9pm3ffhq5djx75ycmumvpwlxhulss3e2yv",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "ePxcUeuPx3UiqfCPLioJjYwDgk/E7HGhatxfzl4/vI0="
              },
              "revoked": false,
              "status": 2,
              "tokens": "1340",
              "delegator_shares": "1340",
              "description": {
                "moniker": "shensi-cosmos-validator",
                "identity": "6F7EB61673529227",
                "website": "https://validator.shensi.com",
                "details": "Shensi Public Chain"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1cze0acxhmpx26awlc98h2v33ay95k232yenkqx",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "4t8TLw6CWORPo3EecYgZpgPeNgC0UQxZVq6xuBxuwwI="
              },
              "revoked": false,
              "status": 0,
              "tokens": "54/5",
              "delegator_shares": "12",
              "description": {
                "moniker": "eon-sent-01",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "First Time I am becoming a validator"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1cypmdvszcd9kd3jenmqxd03cpceql8rtuvxftp",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "aUViYC2znC55sleHfmsIN9hZ45SbYPbDcYA0gVzglsc="
              },
              "revoked": false,
              "status": 2,
              "tokens": "1100",
              "delegator_shares": "1100",
              "description": {
                "moniker": "Gold2",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1cy5t4e62z4h7r74cn7vkylx77n5s4hds2xeud6",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "wwgoAZrmpxQRYXHDXRUBX4kX7IBTpDsvuyxFfKuhEZI="
              },
              "revoked": false,
              "status": 2,
              "tokens": "20",
              "delegator_shares": "20",
              "description": {
                "moniker": "fourthstate",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1cyayv35tv47t829mel6jm0vqzmrdhf3jq87tqg",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "u4GEbsk9IEF56V1am5dRtAWXz4iFQkO03FVL87BZXIM="
              },
              "revoked": false,
              "status": 0,
              "tokens": "81",
              "delegator_shares": "100",
              "description": {
                "moniker": "buckster",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1eg88jsn38mrael6cqu7d2u8j6dynya7fv2p2tl",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "cCoFsZzKZ9SQZbHe4NueVObIezP6ts0tRTZ/aN96dig="
              },
              "revoked": false,
              "status": 2,
              "tokens": "150",
              "delegator_shares": "150",
              "description": {
                "moniker": "Dori",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1e3qm8jd9357zhdemwnaafmf0wy3f4yqmd307c2",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "HjSC7VkhKih6xMhudlqfaFE8ZZnP8RKJPv4iqR7RhcE="
              },
              "revoked": false,
              "status": 2,
              "tokens": "392",
              "delegator_shares": "392",
              "description": {
                "moniker": "firstblock.io",
                "identity": "23D9B8528FC93D58",
                "website": "https://firstblock.io/",
                "details": "You Delegate. We Validate."
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1e32c8x0dcgv0kz6qr7xceerne2gznt5s8e8tr7",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "IS4skJndcdUJqSxMPkMmOD4rGaySvA8JJZk/TBxN7dQ="
              },
              "revoked": false,
              "status": 2,
              "tokens": "32",
              "delegator_shares": "32",
              "description": {
                "moniker": "hallophen",
                "identity": "[do-not-modify]",
                "website": "altheamesh.com",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1ejhdq2gpj20schz5n2r7lyxt6pwp2usahaf3wz",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "s3yBiatOnBzzPwfAvCj8Np5UkPx9aw/iVAc6xtBFMPo="
              },
              "revoked": false,
              "status": 0,
              "tokens": "45/2",
              "delegator_shares": "25",
              "description": {
                "moniker": "wancloud",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1ehhkppuyz9mz9qeevqyvhrl2kyqrc0vuhz2436",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "z4VrDRaACHkmiCZr5fZmCo5awgV3Bi16HMDYjU9EKPY="
              },
              "revoked": false,
              "status": 0,
              "tokens": "18",
              "delegator_shares": "20",
              "description": {
                "moniker": "monochord",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1emaa7mwgpnpmc7yptm728ytp9quamsvu9rk4hp",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "TwzOJ4GcN+ZTswub4R8488SrKeWXjY/PaqCF5neXJig="
              },
              "revoked": false,
              "status": 2,
              "tokens": "400",
              "delegator_shares": "400",
              "description": {
                "moniker": "kochacolaj",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr16rcrnftjyl2mctz78825ng8tx5ss22jf6jcp9l",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "S8s6fdAQNQ3bN9SNVAsHB/j8uv1CM1roxeLesL+fh4g="
              },
              "revoked": false,
              "status": 2,
              "tokens": "64650023/500000",
              "delegator_shares": "4428528062450529/30825011500000",
              "description": {
                "moniker": "luigi001",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr16yxdmg8w5vc038cg0myh0w32s4h74ch59j67u6",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "AFnJW89yDemEVNQFHXagLxvI3FCsvBlw5l4yfZujbH4="
              },
              "revoked": false,
              "status": 0,
              "tokens": "27/10",
              "delegator_shares": "3",
              "description": {
                "moniker": "Forbole-2",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr16dfl4ruvt54qyhyf26fwmle62jnywp7mu9akt7",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "P8qibtjNrOXNb8HlESQEPpdGeZrN43xmEfer7NJQ/2A="
              },
              "revoked": false,
              "status": 0,
              "tokens": "18",
              "delegator_shares": "20",
              "description": {
                "moniker": "magoo",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr163emaq8ahwp25pmfqrjaz3c7wm9vtytcwmlphg",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "NoXstzpWjOU8AFUUCd2190pmOzGRFMnC8nGuHLJhFrk="
              },
              "revoked": false,
              "status": 0,
              "tokens": "18",
              "delegator_shares": "20",
              "description": {
                "moniker": "0xMrManlu-Validator",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1640dsytftqn2uxk4prlfxjvn46nnp8wxt5s0ye",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "puPhAssdpqKTl8h1eZqSScYPjSSQJA/EMEnFLh3GyjE="
              },
              "revoked": false,
              "status": 0,
              "tokens": "117",
              "delegator_shares": "130",
              "description": {
                "moniker": "wancloud02",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr164jntjfk9zs8x29mc27qansfwvjqs60gj6ermu",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "8K3clCjVU33BTIpUhdahGmu++WxHj4NUE9krCRkk++s="
              },
              "revoked": false,
              "status": 2,
              "tokens": "1008",
              "delegator_shares": "1008",
              "description": {
                "moniker": "lunamint",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr16hvayg2lv7rmqz7rjsd3fnq3y34j6r97dxrecf",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "72qgkZiyI788usRQ/tGmbs3TzU9t4quP+WX7rxPX/7I="
              },
              "revoked": false,
              "status": 0,
              "tokens": "89/2",
              "delegator_shares": "55",
              "description": {
                "moniker": "spark",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1mtez787cl6phgkylneq9wjj3d6letn0247lhf5",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "sQfjU1RODJ/QxQLhZ3AabdaH3aNLz3PaERf2kwEFWRg="
              },
              "revoked": false,
              "status": 0,
              "tokens": "309",
              "delegator_shares": "1030",
              "description": {
                "moniker": "wxvali_slave",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 3,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1mvmdrtuxjttv90wssvh6xreeguvq8tksm789m5",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "9eFY/cE8j9uIbxZIjys8e6O0T5E3xKvllw7EFBETY2k="
              },
              "revoked": false,
              "status": 0,
              "tokens": "1579/10",
              "delegator_shares": "650/3",
              "description": {
                "moniker": "cm-gaia-2",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1mwmcpq5nhaen8zyy86mrrept2gmp0z5peapkhu",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "K4kLogLtZxqrYSqRVJfrFm9tUG+Tc3QWXWIewnAgI9w="
              },
              "revoked": false,
              "status": 2,
              "tokens": "147",
              "delegator_shares": "147",
              "description": {
                "moniker": "abcin",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1uffeh2wq0lrvwm7z77c595v040q4kk4g3v8vte",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "pggq0WUTlsiA8cBiptRgxlw4WUG2GmXfLQxkbMbCBPk="
              },
              "revoked": false,
              "status": 2,
              "tokens": "155",
              "delegator_shares": "155",
              "description": {
                "moniker": "wecosmos",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1uknhjkw775zp5xkztc4clxr240ut59ek5seerx",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "8onIV6eJHhGmQObsD7HRUdG/6OC2ogxCRIdLyh1HXxg="
              },
              "revoked": false,
              "status": 0,
              "tokens": "45",
              "delegator_shares": "50",
              "description": {
                "moniker": "east_test",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1u6ddcsjueax884l3tfrs66497c7g86sk3vfmqj",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "0ZPHmftStU+DSqvnIleNyaFmBQsgWJwPC7vk4uF5u00="
              },
              "revoked": false,
              "status": 2,
              "tokens": "4913/2",
              "delegator_shares": "24565/9",
              "description": {
                "moniker": "sentinel.co",
                "identity": "[do-not-modify]",
                "website": "https://sentinel.co/",
                "details": "Sentinel-validator-2 on gaia-7005"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1umaajfgap5ef6yxvk5706kwk8j08l7wh6h9fp2",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "ENAVynNXVpj/IdYx9kCPKaPs4bWSxRIHNlmS9QiDuZQ="
              },
              "revoked": false,
              "status": 2,
              "tokens": "1155",
              "delegator_shares": "1155",
              "description": {
                "moniker": "iqlusion",
                "identity": "7843656AA22520BC",
                "website": "https://www.iqlusion.io",
                "details": "None"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1ul3ef4trvjfmgaptu70ed4fjlky2aa49t4pfle",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "5XCMLGNTGkMh9/JsUzi/8Xl7Y0cdDe38stX+zvBLXuE="
              },
              "revoked": false,
              "status": 2,
              "tokens": "4871/5",
              "delegator_shares": "9742/9",
              "description": {
                "moniker": "bharvest-validator(KR1)",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1uluae4kaxq3wxd698t256a8luv2tf0ngy40qp3",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "jnLS3Z0zv0IG1jPaHZH+DUmh0zcZv2u4jFhK28/LtpU="
              },
              "revoked": false,
              "status": 0,
              "tokens": "4/5",
              "delegator_shares": "1",
              "description": {
                "moniker": "liangping-bj",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1artqm53zt7j7znm8cpgqzgp9jhlx0xrtle2pjr",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "yLIfj0Nm4P1dqnhE4qfh8YYKns1+Cs25WkE7NSrnmtk="
              },
              "revoked": false,
              "status": 0,
              "tokens": "9",
              "delegator_shares": "10",
              "description": {
                "moniker": "art_Fvalid",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1agfqwpcxpl8s8pl3uhxhw7h0r05an60dmlgukx",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "PZcgDAjOlJunLVEj896yzbMUKiQzSj4k9Xeqt7votvo="
              },
              "revoked": false,
              "status": 0,
              "tokens": "9/10",
              "delegator_shares": "1",
              "description": {
                "moniker": "baryon.network",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1aaumwh9rc9ksh7dwuxux43jd93e7t4u88g5kg4",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "JREG9eji9gfJR2b61qMu8jNZ835ezPY2A4o+zjYMfDg="
              },
              "revoked": false,
              "status": 2,
              "tokens": "107",
              "delegator_shares": "107",
              "description": {
                "moniker": "kamon-moniker",
                "identity": "A6B9D417A70E5C33",
                "website": "https://deroris.net/",
                "details": "welcome to Deroris net"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr17rqsh3fw6rehnwzarz90jkqtt5fupmh50gy556",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "PxJbo5FKA6mXtgwclRQVNIjOCQK3Q7WkLQrvM9lYbGI="
              },
              "revoked": false,
              "status": 2,
              "tokens": "156",
              "delegator_shares": "156",
              "description": {
                "moniker": "space4",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr17gvlvfpsfl6hffn5u2hahk22un4ynpykc44tat",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "oL/QCr7LEOivyTqpGrmwVd1r+hYI2WB5+kSVzpDMxx4="
              },
              "revoked": false,
              "status": 2,
              "tokens": "700",
              "delegator_shares": "700",
              "description": {
                "moniker": "gruberx",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr17nlusdvrk34fq65jemy3umfjfwaxfzv4asyl60",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "H0SIA/BU6Xj8oT5bQkvLpEITN3CqFLbMeBcQ72NZrAE="
              },
              "revoked": false,
              "status": 2,
              "tokens": "135",
              "delegator_shares": "150",
              "description": {
                "moniker": "4455",
                "identity": "",
                "website": "",
                "details": ""
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1lzy0ztj9cz7lemlev99u0upulqjd5cu6nk3t2n",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "PyfMXdTSemRHebEEXNzkf808ePsIX6c8gjdF3IpQ6uo="
              },
              "revoked": false,
              "status": 0,
              "tokens": "17/2",
              "delegator_shares": "95/9",
              "description": {
                "moniker": "gkrizek validator",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1l9jt8xgejkcm4sulhwj8c83ftprz5p9lyq4605",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "WRsXnLz3gf8o4lYYeCZjAXgPU1cdmOOYPdy7aY63iIA="
              },
              "revoked": false,
              "status": 2,
              "tokens": "100",
              "delegator_shares": "100",
              "description": {
                "moniker": "ColmenaCowork_SVQ",
                "identity": "",
                "website": "http://www.coworkingcolmena.com/",
                "details": "MISTAKES are proof that you are TRYING!"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1lvmxcdhua2kt3qd86wnnej5m6gm4max4um5sx2",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "0XvxGzOAer6H/VZOQCx7CH0/q0BqetMpoArXPhi6p50="
              },
              "revoked": false,
              "status": 0,
              "tokens": "12/5",
              "delegator_shares": "3",
              "description": {
                "moniker": "SteakSquad",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1lw2d0qljstqzp6fqhsd5jazackgad9y2e3u2n8",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "I2ILaY31gonxYQhlIk4PFUIN+Pk7+9dDTK1C/s+Vcb0="
              },
              "revoked": false,
              "status": 0,
              "tokens": "9/2",
              "delegator_shares": "5",
              "description": {
                "moniker": "snatico",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1l0qw5znfd6e8pshpjvyghjjzyr4l6ymla080lt",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "P9RgE4RMQT/aHap2oICpwpgKeBAwxPUwuU9zIffKFNM="
              },
              "revoked": false,
              "status": 2,
              "tokens": "1231",
              "delegator_shares": "1231",
              "description": {
                "moniker": "Nylira Validator",
                "identity": "6A0D65E29A4CBC8E",
                "website": "https://nylira.net",
                "details": "Peng Zhong \u003cCDO of Tendermint\u003e"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1l3agyg568dsk2r9v7f35wge3yr85ampza7cdfj",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "4DEDoU/RsHMPS54GzgwkWnW5zPQfMt9aInFFc3GyfA8="
              },
              "revoked": false,
              "status": 2,
              "tokens": "2528/5",
              "delegator_shares": "5056/9",
              "description": {
                "moniker": "crytter",
                "identity": "C1FBCF0D1850E5DD",
                "website": "https://cosmos.crytter.io/",
                "details": "Woof woof!"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1lhjta6nt0lewj05m8444tuyhalkrffgpm7njpp",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "DsTbM0AgHfhSUKvOGkxudDOY3ojYT6bifhpelqHs8+s="
              },
              "revoked": false,
              "status": 2,
              "tokens": "100",
              "delegator_shares": "100",
              "description": {
                "moniker": "BFF-Validator-7000",
                "identity": "",
                "website": "https://blockchain-friendly-family.io",
                "details": "Friendly blockchain"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            },
            {
              "owner": "cosmosaccaddr1lmazkxl9dtetn5lxzl45ut3ywwn94xvs3l4z8f",
              "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "O95l52FybUIUBkaYeGHCLRNkSI5zAewb4SxzS0LOb50="
              },
              "revoked": false,
              "status": 2,
              "tokens": "115",
              "delegator_shares": "115",
              "description": {
                "moniker": "loafer",
                "identity": "[do-not-modify]",
                "website": "[do-not-modify]",
                "details": "[do-not-modify]"
              },
              "bond_height": "0",
              "bond_intra_tx_counter": 0,
              "proposer_reward_pool": null,
              "commission": "0",
              "commission_max": "0",
              "commission_change_rate": "0",
              "commission_change_today": "0",
              "prev_bonded_tokens": "0"
            }
          ],
          "bonds": [
            {
              "delegator_addr": "cosmosaccaddr1qz9pauujxsypyu3q9n0t6mydsfwev50j8g8ras",
              "validator_addr": "cosmosaccaddr19h5dc2c76h4ryvl9zl9373mk52ka77vy5g48gs",
              "shares": "12",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1qz9pauujxsypyu3q9n0t6mydsfwev50j8g8ras",
              "validator_addr": "cosmosaccaddr100ztxxmwkqr44jts3a543xhme7vjychzyagz3g",
              "shares": "17",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1qv68kxwmarzefrcgnu4trdrznslzdtra7gfxhy",
              "validator_addr": "cosmosaccaddr1qv68kxwmarzefrcgnu4trdrznslzdtra7gfxhy",
              "shares": "7",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1qvut6w8v5fccephuv502qhmvhn7hy6cnddrdyf",
              "validator_addr": "cosmosaccaddr1qvut6w8v5fccephuv502qhmvhn7hy6cnddrdyf",
              "shares": "110",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1qsvdcgf7zyue0usl2tt3ptn0x6lzs0ua6y5dgs",
              "validator_addr": "cosmosaccaddr1l0qw5znfd6e8pshpjvyghjjzyr4l6ymla080lt",
              "shares": "1",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1qjzjfn55hygaak9l9x04z792mexce2zd2gvkzk",
              "validator_addr": "cosmosaccaddr1qjzjfn55hygaak9l9x04z792mexce2zd2gvkzk",
              "shares": "548/3",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1qk2q3l28t6qhzyc7p832ep8uwd0ypstgqla8yx",
              "validator_addr": "cosmosaccaddr1qk2q3l28t6qhzyc7p832ep8uwd0ypstgqla8yx",
              "shares": "119",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1qkc3tghc3fms7eye7vtu0g0370epr4jkje2ne7",
              "validator_addr": "cosmosaccaddr1qkc3tghc3fms7eye7vtu0g0370epr4jkje2ne7",
              "shares": "2161",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1qkllv8f6qakkw3hk9dqvytys090lk6twsyv8vf",
              "validator_addr": "cosmosaccaddr1qkllv8f6qakkw3hk9dqvytys090lk6twsyv8vf",
              "shares": "250",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1qexep37ryqkj6rvaah5us24tve46hac755vru3",
              "validator_addr": "cosmosaccaddr1rcp29q3hpd246n6qak7jluqep4v006cdj2zse3",
              "shares": "30",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1qe82wd64ncuh8z05kmwpqhc5vkq4wsxqv67cvz",
              "validator_addr": "cosmosaccaddr1y6d5g4s75s9m8xemrs6d9am2cef4nm3ku0s5lv",
              "shares": "6",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1qmw6d8uh3rs9uj8c0tztznu9x4zezekq89635d",
              "validator_addr": "cosmosaccaddr16rcrnftjyl2mctz78825ng8tx5ss22jf6jcp9l",
              "shares": "5",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1qarnnunjrsd66yuz65ugsl7pm3v3hxrxd5jvaq",
              "validator_addr": "cosmosaccaddr1qarnnunjrsd66yuz65ugsl7pm3v3hxrxd5jvaq",
              "shares": "60",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1pj6vfgetplkcuq0k9m98dpej6zew3pdest842v",
              "validator_addr": "cosmosaccaddr1pj6vfgetplkcuq0k9m98dpej6zew3pdest842v",
              "shares": "10",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1p56gv748xfd74qek5e637vhcr6tyjd9ukqfecc",
              "validator_addr": "cosmosaccaddr1p56gv748xfd74qek5e637vhcr6tyjd9ukqfecc",
              "shares": "1100",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1pc63alzz2qdgxk7f3mhr56xly8xz5cs47k294m",
              "validator_addr": "cosmosaccaddr1xvt4e7xd0j9dwv2w83g50tpcltsl90h5dfnz6h",
              "shares": "2",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1pc63alzz2qdgxk7f3mhr56xly8xz5cs47k294m",
              "validator_addr": "cosmosaccaddr18thamkhnj9wz8pa4nhnp9rldprgant57ryzag7",
              "shares": "20/9",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1pm0gx3lk46gf8vyj5my9w5tk06cuq66ll77ugj",
              "validator_addr": "cosmosaccaddr1pm0gx3lk46gf8vyj5my9w5tk06cuq66ll77ugj",
              "shares": "100",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1pmntq5en2rgtr5rzr4e304efrve4lr43z32y5s",
              "validator_addr": "cosmosaccaddr1pmntq5en2rgtr5rzr4e304efrve4lr43z32y5s",
              "shares": "350",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1p7jg2t7xrjmxpu94tp9wawmnx64ds8skfmsqe0",
              "validator_addr": "cosmosaccaddr18thamkhnj9wz8pa4nhnp9rldprgant57ryzag7",
              "shares": "50/9",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1p75wzde5z2ataefw65j3mtrm25kha779yfrvt6",
              "validator_addr": "cosmosaccaddr1p75wzde5z2ataefw65j3mtrm25kha779yfrvt6",
              "shares": "49",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1zpx36km7tk5cksyzxgvcp6g552p3uhwx84r53q",
              "validator_addr": "cosmosaccaddr1zpx36km7tk5cksyzxgvcp6g552p3uhwx84r53q",
              "shares": "150",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1ztc70khr9aghgwmuxap5y9ufavpxutjtkvhpe7",
              "validator_addr": "cosmosaccaddr1ztc70khr9aghgwmuxap5y9ufavpxutjtkvhpe7",
              "shares": "100",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1zvyjt8aw6xgqfxmh7mguvntudrz4llln2dyaam",
              "validator_addr": "cosmosaccaddr16rcrnftjyl2mctz78825ng8tx5ss22jf6jcp9l",
              "shares": "1",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1zds6r7jyxyxcpf05r5yyyy3u8q2rvj9kkc6vcv",
              "validator_addr": "cosmosaccaddr1zds6r7jyxyxcpf05r5yyyy3u8q2rvj9kkc6vcv",
              "shares": "100",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1zkupr83hrzkn3up5elktzcq3tuft8nxseu9x80",
              "validator_addr": "cosmosaccaddr1zkupr83hrzkn3up5elktzcq3tuft8nxseu9x80",
              "shares": "150",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1zagese3pa7k002epyy8cpt783pe2p2n5qprug6",
              "validator_addr": "cosmosaccaddr1xlkah6d2pcr8rchunmn6qdmdptr02844eee6le",
              "shares": "30",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1r2hjwvyudlmxlr3dqdq6r72khcc7yej9zew48d",
              "validator_addr": "cosmosaccaddr1ja98tyxchsde6v3l02thp4fc35eymg88kts070",
              "shares": "3",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1rvm0em6w3qkzcwnzf9hkqvksujl895dfww4ecn",
              "validator_addr": "cosmosaccaddr1rvm0em6w3qkzcwnzf9hkqvksujl895dfww4ecn",
              "shares": "7540/3",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1rww4u22qqf2rhw02fww63t539d3ppeur3ywff3",
              "validator_addr": "cosmosaccaddr1rww4u22qqf2rhw02fww63t539d3ppeur3ywff3",
              "shares": "995",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1r05wuvemjzl25n3lspfku03e069cs5xsulfusn",
              "validator_addr": "cosmosaccaddr1r05wuvemjzl25n3lspfku03e069cs5xsulfusn",
              "shares": "5",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1rj74vaqm0xkxl5cjjam63mayh4x6at3m379ulv",
              "validator_addr": "cosmosaccaddr1rj74vaqm0xkxl5cjjam63mayh4x6at3m379ulv",
              "shares": "150",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1r4lzzzf683ec70fxfd9a3jmczxxkdxpfafz4ar",
              "validator_addr": "cosmosaccaddr1r4lzzzf683ec70fxfd9a3jmczxxkdxpfafz4ar",
              "shares": "10",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1rcp29q3hpd246n6qak7jluqep4v006cdj2zse3",
              "validator_addr": "cosmosaccaddr1rcp29q3hpd246n6qak7jluqep4v006cdj2zse3",
              "shares": "970",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1yqpa5vhs8nq9mguyd2j5k5kn7dlqlgmjl4h67x",
              "validator_addr": "cosmosaccaddr1yqpa5vhs8nq9mguyd2j5k5kn7dlqlgmjl4h67x",
              "shares": "2",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1ypk9spe0lsta44f3ay6rqrzygh90ekayxcxcfn",
              "validator_addr": "cosmosaccaddr1rww4u22qqf2rhw02fww63t539d3ppeur3ywff3",
              "shares": "70",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1yfd9a4sqm8jm7lq932jc6fu8k86f394tgc2sku",
              "validator_addr": "cosmosaccaddr1yfd9a4sqm8jm7lq932jc6fu8k86f394tgc2sku",
              "shares": "10",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1yfseqtj5sjhzz2q2ym09jym4h4nc4yevae0jp2",
              "validator_addr": "cosmosaccaddr1yfseqtj5sjhzz2q2ym09jym4h4nc4yevae0jp2",
              "shares": "1115",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1yfseqtj5sjhzz2q2ym09jym4h4nc4yevae0jp2",
              "validator_addr": "cosmosaccaddr15mzck4rm5w55sfgr9kgexgtna22v608907t8y6",
              "shares": "85",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1y2z20pwqu5qpclque3pqkguruvheum2djtzjw3",
              "validator_addr": "cosmosaccaddr1y2z20pwqu5qpclque3pqkguruvheum2djtzjw3",
              "shares": "250",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1yd0rklq45zg89fywr89ccutlcwp9kehhh0z03k",
              "validator_addr": "cosmosaccaddr1yd0rklq45zg89fywr89ccutlcwp9kehhh0z03k",
              "shares": "155",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1y3fydat73q569cceaccpuytm3y4tdpv3n0fhdv",
              "validator_addr": "cosmosaccaddr1y3fydat73q569cceaccpuytm3y4tdpv3n0fhdv",
              "shares": "10",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1y32vvanvz88ajamugxrgl7vnvf0vtmjg5wxg7f",
              "validator_addr": "cosmosaccaddr1y32vvanvz88ajamugxrgl7vnvf0vtmjg5wxg7f",
              "shares": "20",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1yj87w5yjftyar3mhuefg569q4v26rfp4ntacc5",
              "validator_addr": "cosmosaccaddr18thamkhnj9wz8pa4nhnp9rldprgant57ryzag7",
              "shares": "100/9",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1y6d5g4s75s9m8xemrs6d9am2cef4nm3ku0s5lv",
              "validator_addr": "cosmosaccaddr1y6d5g4s75s9m8xemrs6d9am2cef4nm3ku0s5lv",
              "shares": "35",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1y7n3emnw6mkxutvs26rczynpvq84ddqu332hlv",
              "validator_addr": "cosmosaccaddr1y7n3emnw6mkxutvs26rczynpvq84ddqu332hlv",
              "shares": "50",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1ylz6tqpgz0flpl6h5szmj6mmzyr598rq7heuye",
              "validator_addr": "cosmosaccaddr1ylz6tqpgz0flpl6h5szmj6mmzyr598rq7heuye",
              "shares": "20000/81",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr19q6y4mk3x668vm2jx0pu6e0jfy5a39jz0qwvxt",
              "validator_addr": "cosmosaccaddr19q6y4mk3x668vm2jx0pu6e0jfy5a39jz0qwvxt",
              "shares": "60",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr19pkt3aeaprl8dt9x5c77u9unkrndfd3ze40tj9",
              "validator_addr": "cosmosaccaddr19pkt3aeaprl8dt9x5c77u9unkrndfd3ze40tj9",
              "shares": "5",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr19rlhdmplhnmvy3jmcd835awask4l46nvp3pt4g",
              "validator_addr": "cosmosaccaddr19rlhdmplhnmvy3jmcd835awask4l46nvp3pt4g",
              "shares": "30",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1990ceh7wh50j2z6wq08ekzez7yr6hpwtpap9ku",
              "validator_addr": "cosmosaccaddr12ceualfg92x7du73mtcv0zya4nxvq3tl2m52uz",
              "shares": "15",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1990ceh7wh50j2z6wq08ekzez7yr6hpwtpap9ku",
              "validator_addr": "cosmosaccaddr126ayk3hse5zvk9gxfmpsjr9565ef72pv9g20yx",
              "shares": "90100/65691",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1990ceh7wh50j2z6wq08ekzez7yr6hpwtpap9ku",
              "validator_addr": "cosmosaccaddr16rcrnftjyl2mctz78825ng8tx5ss22jf6jcp9l",
              "shares": "68500023/61650023",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1990ceh7wh50j2z6wq08ekzez7yr6hpwtpap9ku",
              "validator_addr": "cosmosaccaddr1u6ddcsjueax884l3tfrs66497c7g86sk3vfmqj",
              "shares": "10/9",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr19dfkpwfj3mls96ny2ulgq7cf0dh5xs5a05pn6e",
              "validator_addr": "cosmosaccaddr19dfkpwfj3mls96ny2ulgq7cf0dh5xs5a05pn6e",
              "shares": "3",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr193vn0gk3nsmjkxwz78gce8e8mkmagmvulpg5jt",
              "validator_addr": "cosmosaccaddr193vn0gk3nsmjkxwz78gce8e8mkmagmvulpg5jt",
              "shares": "1020",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr19n4s8fx9hr0j2dttr4fg48wr7kpxxste7evjzk",
              "validator_addr": "cosmosaccaddr19n4s8fx9hr0j2dttr4fg48wr7kpxxste7evjzk",
              "shares": "230",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr19ncpp07tuqqgu2kjp4wq2zlt3s0q5f9ywvhwuw",
              "validator_addr": "cosmosaccaddr19ncpp07tuqqgu2kjp4wq2zlt3s0q5f9ywvhwuw",
              "shares": "110",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr195qlnqc2mxsxnfvcljqr5ssfjjymn0usk35v75",
              "validator_addr": "cosmosaccaddr195qlnqc2mxsxnfvcljqr5ssfjjymn0usk35v75",
              "shares": "5",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr194pux2wkhuusey6x4wwaqm0wh26cleh06u28lr",
              "validator_addr": "cosmosaccaddr194pux2wkhuusey6x4wwaqm0wh26cleh06u28lr",
              "shares": "25",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr19h5dc2c76h4ryvl9zl9373mk52ka77vy5g48gs",
              "validator_addr": "cosmosaccaddr19h5dc2c76h4ryvl9zl9373mk52ka77vy5g48gs",
              "shares": "1",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr19uhnhct0p45ek6qxp3cjjrjtz4pacwcsgzvpuj",
              "validator_addr": "cosmosaccaddr19uhnhct0p45ek6qxp3cjjrjtz4pacwcsgzvpuj",
              "shares": "956",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1xtkfzzty24l66a0f9guqt4w8k0p66fpat7rlf2",
              "validator_addr": "cosmosaccaddr1640dsytftqn2uxk4prlfxjvn46nnp8wxt5s0ye",
              "shares": "10",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1xvzw5nz4mp5s5cat7rancc248v5wpecrlqtw6n",
              "validator_addr": "cosmosaccaddr1xvzw5nz4mp5s5cat7rancc248v5wpecrlqtw6n",
              "shares": "121",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1xvt4e7xd0j9dwv2w83g50tpcltsl90h5dfnz6h",
              "validator_addr": "cosmosaccaddr1xvt4e7xd0j9dwv2w83g50tpcltsl90h5dfnz6h",
              "shares": "100",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1xupfqk73y7rmc6qdgv7rtjy8wsngvt2g2t59t3",
              "validator_addr": "cosmosaccaddr1xupfqk73y7rmc6qdgv7rtjy8wsngvt2g2t59t3",
              "shares": "150",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1xuwcmdvef5xudnyqa9gh5vqlq6ff0meu87zjxh",
              "validator_addr": "cosmosaccaddr1xuwcmdvef5xudnyqa9gh5vqlq6ff0meu87zjxh",
              "shares": "171",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1xa8z3sza70wrt06yfhlxq0ujg6fa7ld8h7m6xz",
              "validator_addr": "cosmosaccaddr1xa8z3sza70wrt06yfhlxq0ujg6fa7ld8h7m6xz",
              "shares": "505",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1xlkah6d2pcr8rchunmn6qdmdptr02844eee6le",
              "validator_addr": "cosmosaccaddr1xlkah6d2pcr8rchunmn6qdmdptr02844eee6le",
              "shares": "1",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr18zwvgsrl3gc8y0fhzdev378u9gn37d9024njy0",
              "validator_addr": "cosmosaccaddr18zwvgsrl3gc8y0fhzdev378u9gn37d9024njy0",
              "shares": "8000/9",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr18z3pnjgdtt337z8g5drfts7r3fm6n9a0896h0r",
              "validator_addr": "cosmosaccaddr18z3pnjgdtt337z8g5drfts7r3fm6n9a0896h0r",
              "shares": "380",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr18raps2wmm63n2gldsrqc5x5z023p559l7r5vv6",
              "validator_addr": "cosmosaccaddr18raps2wmm63n2gldsrqc5x5z023p559l7r5vv6",
              "shares": "50",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr182rhfxwz8ycec8sjq5frdg8qpwcnwrclu2gnya",
              "validator_addr": "cosmosaccaddr1640dsytftqn2uxk4prlfxjvn46nnp8wxt5s0ye",
              "shares": "10",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1825jrcgu3keurjlnpsujx9y8lrzgddjfsmax0n",
              "validator_addr": "cosmosaccaddr1825jrcgu3keurjlnpsujx9y8lrzgddjfsmax0n",
              "shares": "2099/9",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr182ujqw3r8p5fffjqkf0rxzj29pg5q96nxd2khq",
              "validator_addr": "cosmosaccaddr1zpx36km7tk5cksyzxgvcp6g552p3uhwx84r53q",
              "shares": "100",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr182ujqw3r8p5fffjqkf0rxzj29pg5q96nxd2khq",
              "validator_addr": "cosmosaccaddr1yfseqtj5sjhzz2q2ym09jym4h4nc4yevae0jp2",
              "shares": "100",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr182ujqw3r8p5fffjqkf0rxzj29pg5q96nxd2khq",
              "validator_addr": "cosmosaccaddr182ujqw3r8p5fffjqkf0rxzj29pg5q96nxd2khq",
              "shares": "1750",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr182ujqw3r8p5fffjqkf0rxzj29pg5q96nxd2khq",
              "validator_addr": "cosmosaccaddr12ceualfg92x7du73mtcv0zya4nxvq3tl2m52uz",
              "shares": "100",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr182ujqw3r8p5fffjqkf0rxzj29pg5q96nxd2khq",
              "validator_addr": "cosmosaccaddr14e774gfzt5l9ka766ehfgu6n5hgy9f3sehzyv8",
              "shares": "100",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr182ujqw3r8p5fffjqkf0rxzj29pg5q96nxd2khq",
              "validator_addr": "cosmosaccaddr1hfygmryre7r8m9pfqmc85y7kw7ejphmp59k2x8",
              "shares": "100",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr18tkxwzhyn4a5ss63vttwpespvm7kxm895m9ppf",
              "validator_addr": "cosmosaccaddr1l0qw5znfd6e8pshpjvyghjjzyr4l6ymla080lt",
              "shares": "40",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr18thamkhnj9wz8pa4nhnp9rldprgant57ryzag7",
              "validator_addr": "cosmosaccaddr182ujqw3r8p5fffjqkf0rxzj29pg5q96nxd2khq",
              "shares": "15",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr18thamkhnj9wz8pa4nhnp9rldprgant57ryzag7",
              "validator_addr": "cosmosaccaddr18thamkhnj9wz8pa4nhnp9rldprgant57ryzag7",
              "shares": "6632/3",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr18thamkhnj9wz8pa4nhnp9rldprgant57ryzag7",
              "validator_addr": "cosmosaccaddr120skmenn2a0ra8y6zassrxvnfc5rlme8rqarvs",
              "shares": "40",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr18sqscwa5g4q6quc97eggwqk4wyzjyvrmvtq7f2",
              "validator_addr": "cosmosaccaddr19uhnhct0p45ek6qxp3cjjrjtz4pacwcsgzvpuj",
              "shares": "540",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr183tl5h22xfz8e8n5dp3fphu55kltfxxtr54vyw",
              "validator_addr": "cosmosaccaddr1uffeh2wq0lrvwm7z77c595v040q4kk4g3v8vte",
              "shares": "50",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr18ja9lhun5le78j36s8rs2je6z4pzc6far7xfuz",
              "validator_addr": "cosmosaccaddr163emaq8ahwp25pmfqrjaz3c7wm9vtytcwmlphg",
              "shares": "10",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr185n9jln2g4w79w39m0g85x2tns6uekqg99tq7u",
              "validator_addr": "cosmosaccaddr185n9jln2g4w79w39m0g85x2tns6uekqg99tq7u",
              "shares": "170",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr18cswjuy9qzwcjnkmvp32ryyht0rlaueazynegn",
              "validator_addr": "cosmosaccaddr1emaa7mwgpnpmc7yptm728ytp9quamsvu9rk4hp",
              "shares": "20",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr18cut539f6w9xewnz0fcxa95784pg4mk5ghk3fz",
              "validator_addr": "cosmosaccaddr1640dsytftqn2uxk4prlfxjvn46nnp8wxt5s0ye",
              "shares": "10",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr18m09d56pg5p2660de4sjfezpd8ud6jfghndfnt",
              "validator_addr": "cosmosaccaddr18m09d56pg5p2660de4sjfezpd8ud6jfghndfnt",
              "shares": "150",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr18up28pmkspg2r8uv879fdlrxg0gpxnnj5nn0gf",
              "validator_addr": "cosmosaccaddr1hm5f52equcd4f268uxf4xhwljpnljvdyqal3dv",
              "shares": "100/9",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1gq0qecxs8xdaarrqxxazwavwxm7qz5jzs5anvt",
              "validator_addr": "cosmosaccaddr1gq0qecxs8xdaarrqxxazwavwxm7qz5jzs5anvt",
              "shares": "1100",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1g8xwmm3tp7pgdrmjgt4fm7d26anrrnle4yavuw",
              "validator_addr": "cosmosaccaddr1g8xwmm3tp7pgdrmjgt4fm7d26anrrnle4yavuw",
              "shares": "100",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1g6sc5t9t68vmcj3alk7dfqr54tvastpxac28k6",
              "validator_addr": "cosmosaccaddr1g6sc5t9t68vmcj3alk7dfqr54tvastpxac28k6",
              "shares": "100",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1fzqwytcm25gsxg67l2z3rtmskpx75u6l5a2mq7",
              "validator_addr": "cosmosaccaddr16rcrnftjyl2mctz78825ng8tx5ss22jf6jcp9l",
              "shares": "5",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1frdjhnkgvlq5s0v5unf9njzlshdtsruv4z3sey",
              "validator_addr": "cosmosaccaddr18thamkhnj9wz8pa4nhnp9rldprgant57ryzag7",
              "shares": "2017900/163449",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1fyevhdmxpkz8stjxggv9wejf9caed02vyp8y5p",
              "validator_addr": "cosmosaccaddr1fyevhdmxpkz8stjxggv9wejf9caed02vyp8y5p",
              "shares": "10",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1ftrggumlvkesqny90par3gc9lp70yqhtlqucsk",
              "validator_addr": "cosmosaccaddr1ftrggumlvkesqny90par3gc9lp70yqhtlqucsk",
              "shares": "2196",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1fskmzyt2hr8usr3s65dq3rfur3fy6g2hjp23tm",
              "validator_addr": "cosmosaccaddr1fskmzyt2hr8usr3s65dq3rfur3fy6g2hjp23tm",
              "shares": "1020",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1f4sah2y8r995wsazt2dw909akhdjvr4f9zv28r",
              "validator_addr": "cosmosaccaddr1uffeh2wq0lrvwm7z77c595v040q4kk4g3v8vte",
              "shares": "5",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1fhfcs5npydv8s96wrk9p5ychctslu92t4n5qd4",
              "validator_addr": "cosmosaccaddr1fhfcs5npydv8s96wrk9p5ychctslu92t4n5qd4",
              "shares": "150",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr12qprqyxtxxlu7unncd9rrhmer5emg2pmrw27qk",
              "validator_addr": "cosmosaccaddr12qprqyxtxxlu7unncd9rrhmer5emg2pmrw27qk",
              "shares": "30",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr12pn4p6r5wjpsep9kn655ng7fh59yez7t0rahru",
              "validator_addr": "cosmosaccaddr12pn4p6r5wjpsep9kn655ng7fh59yez7t0rahru",
              "shares": "100",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr128ty3kzhcepjacu4q0xjgq60qa3zz8na3jl793",
              "validator_addr": "cosmosaccaddr128ty3kzhcepjacu4q0xjgq60qa3zz8na3jl793",
              "shares": "152",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr12fm2n7pkmtqk7untyhs9my749jyzef3vhne95q",
              "validator_addr": "cosmosaccaddr12fm2n7pkmtqk7untyhs9my749jyzef3vhne95q",
              "shares": "1998",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr120skmenn2a0ra8y6zassrxvnfc5rlme8rqarvs",
              "validator_addr": "cosmosaccaddr120skmenn2a0ra8y6zassrxvnfc5rlme8rqarvs",
              "shares": "15",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr12ceualfg92x7du73mtcv0zya4nxvq3tl2m52uz",
              "validator_addr": "cosmosaccaddr12ceualfg92x7du73mtcv0zya4nxvq3tl2m52uz",
              "shares": "1160",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr126ayk3hse5zvk9gxfmpsjr9565ef72pv9g20yx",
              "validator_addr": "cosmosaccaddr126ayk3hse5zvk9gxfmpsjr9565ef72pv9g20yx",
              "shares": "301889030/65691",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1t9rqvvys4zh300x93a6fr3h064gp9y8al7m4en",
              "validator_addr": "cosmosaccaddr1t9rqvvys4zh300x93a6fr3h064gp9y8al7m4en",
              "shares": "606",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1t9nxlu834au3u7mltyf8udzw7kzsf0wntx7e4w",
              "validator_addr": "cosmosaccaddr1t9nxlu834au3u7mltyf8udzw7kzsf0wntx7e4w",
              "shares": "1",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1t9l4ztfwk0mp3za3x38yfqkernuclxwgdqzme6",
              "validator_addr": "cosmosaccaddr15rrxdflhq4744qt9m5ysstq3zpykhacf908eez",
              "shares": "2196",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1t8lqjkvy0w0fzs3kv4s00s9cnlqqwepmnux9g5",
              "validator_addr": "cosmosaccaddr1fskmzyt2hr8usr3s65dq3rfur3fy6g2hjp23tm",
              "shares": "17",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1t8lqjkvy0w0fzs3kv4s00s9cnlqqwepmnux9g5",
              "validator_addr": "cosmosaccaddr164jntjfk9zs8x29mc27qansfwvjqs60gj6ermu",
              "shares": "13",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1tsysfw53f36ap440e5ljf6hnc3ueuwal9vmcx4",
              "validator_addr": "cosmosaccaddr1tsysfw53f36ap440e5ljf6hnc3ueuwal9vmcx4",
              "shares": "2",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1t3afuz2kt99jz4pe5k4vjvkdmldn2x0lqzv83w",
              "validator_addr": "cosmosaccaddr1t3afuz2kt99jz4pe5k4vjvkdmldn2x0lqzv83w",
              "shares": "295",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1t3afuz2kt99jz4pe5k4vjvkdmldn2x0lqzv83w",
              "validator_addr": "cosmosaccaddr10505nl7yftsme9jk2glhjhta7w0475uva87paj",
              "shares": "80/9",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1t3afuz2kt99jz4pe5k4vjvkdmldn2x0lqzv83w",
              "validator_addr": "cosmosaccaddr157mg9hnhchfrqvk3enrvmvj29yhmlwf759xrgw",
              "shares": "3",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1t3afuz2kt99jz4pe5k4vjvkdmldn2x0lqzv83w",
              "validator_addr": "cosmosaccaddr1umaajfgap5ef6yxvk5706kwk8j08l7wh6h9fp2",
              "shares": "5",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1t3lvwnp0rmfyc96c0qh6qa67nplpdhytr5dd59",
              "validator_addr": "cosmosaccaddr1t3lvwnp0rmfyc96c0qh6qa67nplpdhytr5dd59",
              "shares": "250",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1vqhlc253yxx3457lr4grawaqhu8xuphm45eld5",
              "validator_addr": "cosmosaccaddr1825jrcgu3keurjlnpsujx9y8lrzgddjfsmax0n",
              "shares": "79",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1vrc7zpg5teawwuzkfh6t7c6sy353sukhlarxpa",
              "validator_addr": "cosmosaccaddr1vrc7zpg5teawwuzkfh6t7c6sy353sukhlarxpa",
              "shares": "1250",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1v4d9c3wy6n7h4udj4lkumc0l99x95cml44uyv8",
              "validator_addr": "cosmosaccaddr1v4d9c3wy6n7h4udj4lkumc0l99x95cml44uyv8",
              "shares": "1005",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1vk5dwhn8m3sht9xrartj9kuv098t00spkqphcd",
              "validator_addr": "cosmosaccaddr193vn0gk3nsmjkxwz78gce8e8mkmagmvulpg5jt",
              "shares": "10/9",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1vk5dwhn8m3sht9xrartj9kuv098t00spkqphcd",
              "validator_addr": "cosmosaccaddr1vk5dwhn8m3sht9xrartj9kuv098t00spkqphcd",
              "shares": "18",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1vlmpl003lzm4lq0vr94227l2u02n2y490gsc89",
              "validator_addr": "cosmosaccaddr1aaumwh9rc9ksh7dwuxux43jd93e7t4u88g5kg4",
              "shares": "3",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1dq7jt3pn7mrce2twac907jue7hjd0p0rgt3qnq",
              "validator_addr": "cosmosaccaddr1dq7jt3pn7mrce2twac907jue7hjd0p0rgt3qnq",
              "shares": "1",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1dkm8z36m96k8aukg5rxnfpw46dt5mnmvsfuze2",
              "validator_addr": "cosmosaccaddr1dkm8z36m96k8aukg5rxnfpw46dt5mnmvsfuze2",
              "shares": "5",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1dcgqsyam6d0am5z2m7dufp7s9k2px2kcqaswxt",
              "validator_addr": "cosmosaccaddr126ayk3hse5zvk9gxfmpsjr9565ef72pv9g20yx",
              "shares": "450500/65691",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1dcgqsyam6d0am5z2m7dufp7s9k2px2kcqaswxt",
              "validator_addr": "cosmosaccaddr16rcrnftjyl2mctz78825ng8tx5ss22jf6jcp9l",
              "shares": "342500115/61650023",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1dl9mu7uya6mfsd80wyw2rgjmwjtk5ajjf8mxcz",
              "validator_addr": "cosmosaccaddr1dl9mu7uya6mfsd80wyw2rgjmwjtk5ajjf8mxcz",
              "shares": "30",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1dl0hk7064k5wp43zkllmd0p9pgu2mnra898gyh",
              "validator_addr": "cosmosaccaddr1dl0hk7064k5wp43zkllmd0p9pgu2mnra898gyh",
              "shares": "50",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1wf3sncgk7s2ykamrhy4etf09y94rrrg4k73wwr",
              "validator_addr": "cosmosaccaddr1wf3sncgk7s2ykamrhy4etf09y94rrrg4k73wwr",
              "shares": "2045/9",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1w2j930e3ueg8f365ff2t92m5lnp856fe47f3hy",
              "validator_addr": "cosmosaccaddr1w2j930e3ueg8f365ff2t92m5lnp856fe47f3hy",
              "shares": "10",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1wd965wxzyl2krp67sdwddy4ep48xhsrdrq63rd",
              "validator_addr": "cosmosaccaddr1640dsytftqn2uxk4prlfxjvn46nnp8wxt5s0ye",
              "shares": "10",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1wnwr086pvkv2pqa5khznunm93ac52u677zgqe6",
              "validator_addr": "cosmosaccaddr1wnwr086pvkv2pqa5khznunm93ac52u677zgqe6",
              "shares": "5",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1wk0t6na230vxhf6ddresw2c40k5y3ayrww0s2m",
              "validator_addr": "cosmosaccaddr1wk0t6na230vxhf6ddresw2c40k5y3ayrww0s2m",
              "shares": "150",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1whqnxfj5fakx2tmjsja8y9ev7fqdzdldaesmx2",
              "validator_addr": "cosmosaccaddr1640dsytftqn2uxk4prlfxjvn46nnp8wxt5s0ye",
              "shares": "10",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1whxd48da3r56n8eecym8zg0c6xmf35fn2myart",
              "validator_addr": "cosmosaccaddr1whxd48da3r56n8eecym8zg0c6xmf35fn2myart",
              "shares": "1100",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1whd2u62u3zk7w09yuq4cz97fueq8e96qp69xz0",
              "validator_addr": "cosmosaccaddr1whd2u62u3zk7w09yuq4cz97fueq8e96qp69xz0",
              "shares": "20",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1wesup293xcfjze34u98j05ffu04cssy32gjv3u",
              "validator_addr": "cosmosaccaddr1wesup293xcfjze34u98j05ffu04cssy32gjv3u",
              "shares": "11",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr10yx3zn39va2rdvcn76cljlzn7zmlyecc7wdtvq",
              "validator_addr": "cosmosaccaddr16rcrnftjyl2mctz78825ng8tx5ss22jf6jcp9l",
              "shares": "5",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr100ztxxmwkqr44jts3a543xhme7vjychzyagz3g",
              "validator_addr": "cosmosaccaddr100ztxxmwkqr44jts3a543xhme7vjychzyagz3g",
              "shares": "14",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr10s62wd8qca4pnhdk602xqnvpy0ck8uzll4hrsf",
              "validator_addr": "cosmosaccaddr18thamkhnj9wz8pa4nhnp9rldprgant57ryzag7",
              "shares": "20/9",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr10s62wd8qca4pnhdk602xqnvpy0ck8uzll4hrsf",
              "validator_addr": "cosmosaccaddr16rcrnftjyl2mctz78825ng8tx5ss22jf6jcp9l",
              "shares": "14",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr10505nl7yftsme9jk2glhjhta7w0475uva87paj",
              "validator_addr": "cosmosaccaddr1pj6vfgetplkcuq0k9m98dpej6zew3pdest842v",
              "shares": "25",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr10505nl7yftsme9jk2glhjhta7w0475uva87paj",
              "validator_addr": "cosmosaccaddr1p75wzde5z2ataefw65j3mtrm25kha779yfrvt6",
              "shares": "25",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr10505nl7yftsme9jk2glhjhta7w0475uva87paj",
              "validator_addr": "cosmosaccaddr1xlkah6d2pcr8rchunmn6qdmdptr02844eee6le",
              "shares": "25",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr10505nl7yftsme9jk2glhjhta7w0475uva87paj",
              "validator_addr": "cosmosaccaddr1t3afuz2kt99jz4pe5k4vjvkdmldn2x0lqzv83w",
              "shares": "5",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr10505nl7yftsme9jk2glhjhta7w0475uva87paj",
              "validator_addr": "cosmosaccaddr1vrc7zpg5teawwuzkfh6t7c6sy353sukhlarxpa",
              "shares": "5",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr10505nl7yftsme9jk2glhjhta7w0475uva87paj",
              "validator_addr": "cosmosaccaddr1wnwr086pvkv2pqa5khznunm93ac52u677zgqe6",
              "shares": "25",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr10505nl7yftsme9jk2glhjhta7w0475uva87paj",
              "validator_addr": "cosmosaccaddr10505nl7yftsme9jk2glhjhta7w0475uva87paj",
              "shares": "390",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr10505nl7yftsme9jk2glhjhta7w0475uva87paj",
              "validator_addr": "cosmosaccaddr1jt77t00gucffqhrq7lh6vwmwzrwgw2fpvnhwk7",
              "shares": "25",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr10505nl7yftsme9jk2glhjhta7w0475uva87paj",
              "validator_addr": "cosmosaccaddr157eu0k3fvnk4d3rpgtyz2jujf8shlk8mxeqjkk",
              "shares": "25",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr10505nl7yftsme9jk2glhjhta7w0475uva87paj",
              "validator_addr": "cosmosaccaddr157mg9hnhchfrqvk3enrvmvj29yhmlwf759xrgw",
              "shares": "5",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr10505nl7yftsme9jk2glhjhta7w0475uva87paj",
              "validator_addr": "cosmosaccaddr1h4q0xkdg30cl9vw0u8ejm0rs337dszy98gnd4a",
              "shares": "5",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr10505nl7yftsme9jk2glhjhta7w0475uva87paj",
              "validator_addr": "cosmosaccaddr1ejhdq2gpj20schz5n2r7lyxt6pwp2usahaf3wz",
              "shares": "25",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr10505nl7yftsme9jk2glhjhta7w0475uva87paj",
              "validator_addr": "cosmosaccaddr16hvayg2lv7rmqz7rjsd3fnq3y34j6r97dxrecf",
              "shares": "50",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr10505nl7yftsme9jk2glhjhta7w0475uva87paj",
              "validator_addr": "cosmosaccaddr1ul3ef4trvjfmgaptu70ed4fjlky2aa49t4pfle",
              "shares": "25",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr10505nl7yftsme9jk2glhjhta7w0475uva87paj",
              "validator_addr": "cosmosaccaddr1lmazkxl9dtetn5lxzl45ut3ywwn94xvs3l4z8f",
              "shares": "25",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr104e3vqacz2kncw979dczttp5jwts3sacv3ksz7",
              "validator_addr": "cosmosaccaddr1l0qw5znfd6e8pshpjvyghjjzyr4l6ymla080lt",
              "shares": "40",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr106ccejx2h7hwzj9xxaz3hju22yyq5mfy4rydav",
              "validator_addr": "cosmosaccaddr106ccejx2h7hwzj9xxaz3hju22yyq5mfy4rydav",
              "shares": "110",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr10mfxvxjga3rxp5w93rx8jdxtt68j5qk0wv02hn",
              "validator_addr": "cosmosaccaddr10mfxvxjga3rxp5w93rx8jdxtt68j5qk0wv02hn",
              "shares": "10",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr10atd0mjtdt7uzk0jgr32c32semnh3xvu4nx77p",
              "validator_addr": "cosmosaccaddr10atd0mjtdt7uzk0jgr32c32semnh3xvu4nx77p",
              "shares": "5",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1syhzpjwx6vv3erv2myq7txrjhrp95hrhgcl242",
              "validator_addr": "cosmosaccaddr1syhzpjwx6vv3erv2myq7txrjhrp95hrhgcl242",
              "shares": "90",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1st4262sfqk7jqdwp2s7gk2umm5yhexhjrew53z",
              "validator_addr": "cosmosaccaddr1st4262sfqk7jqdwp2s7gk2umm5yhexhjrew53z",
              "shares": "10",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1s4w7w56ymm5fnm25nhdvae5yuc5792zxfynq9h",
              "validator_addr": "cosmosaccaddr1qkc3tghc3fms7eye7vtu0g0370epr4jkje2ne7",
              "shares": "3",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1s4w7w56ymm5fnm25nhdvae5yuc5792zxfynq9h",
              "validator_addr": "cosmosaccaddr1rvm0em6w3qkzcwnzf9hkqvksujl895dfww4ecn",
              "shares": "10/3",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1s4w7w56ymm5fnm25nhdvae5yuc5792zxfynq9h",
              "validator_addr": "cosmosaccaddr18thamkhnj9wz8pa4nhnp9rldprgant57ryzag7",
              "shares": "201790/54483",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1s4w7w56ymm5fnm25nhdvae5yuc5792zxfynq9h",
              "validator_addr": "cosmosaccaddr1ftrggumlvkesqny90par3gc9lp70yqhtlqucsk",
              "shares": "3",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1s4w7w56ymm5fnm25nhdvae5yuc5792zxfynq9h",
              "validator_addr": "cosmosaccaddr126ayk3hse5zvk9gxfmpsjr9565ef72pv9g20yx",
              "shares": "100/27",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1s4w7w56ymm5fnm25nhdvae5yuc5792zxfynq9h",
              "validator_addr": "cosmosaccaddr1hfygmryre7r8m9pfqmc85y7kw7ejphmp59k2x8",
              "shares": "10/3",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1s4w7w56ymm5fnm25nhdvae5yuc5792zxfynq9h",
              "validator_addr": "cosmosaccaddr1u6ddcsjueax884l3tfrs66497c7g86sk3vfmqj",
              "shares": "10/3",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1s4w7w56ymm5fnm25nhdvae5yuc5792zxfynq9h",
              "validator_addr": "cosmosaccaddr1ul3ef4trvjfmgaptu70ed4fjlky2aa49t4pfle",
              "shares": "10/3",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1s4w7w56ymm5fnm25nhdvae5yuc5792zxfynq9h",
              "validator_addr": "cosmosaccaddr1l3agyg568dsk2r9v7f35wge3yr85ampza7cdfj",
              "shares": "10/3",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1shuqhpl273t96yg6nnqvyfeewj3ew3mdcwvcnu",
              "validator_addr": "cosmosaccaddr1shuqhpl273t96yg6nnqvyfeewj3ew3mdcwvcnu",
              "shares": "1130",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr13q0fl9qakvt5zkh6qurlrur4xq7gullkv9f5zq",
              "validator_addr": "cosmosaccaddr13q0fl9qakvt5zkh6qurlrur4xq7gullkv9f5zq",
              "shares": "10",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr13psmeywshfyfcg3tr5l4qm968smtkcu8um0nqn",
              "validator_addr": "cosmosaccaddr13psmeywshfyfcg3tr5l4qm968smtkcu8um0nqn",
              "shares": "10",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr13rt2suxfgefcquy9pveuc5uxfayy5rscr4vld7",
              "validator_addr": "cosmosaccaddr13rt2suxfgefcquy9pveuc5uxfayy5rscr4vld7",
              "shares": "359",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr13fw4ff9u2qftsjj39mda54ufrpjs77q0ke6m49",
              "validator_addr": "cosmosaccaddr13fw4ff9u2qftsjj39mda54ufrpjs77q0ke6m49",
              "shares": "5",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr13fet8nre9ra2rs7gtw035evk2254pgw9mpn79z",
              "validator_addr": "cosmosaccaddr1fskmzyt2hr8usr3s65dq3rfur3fy6g2hjp23tm",
              "shares": "100",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr13dre54wd7md0wg77x5fj5scj5u20e6stvdr829",
              "validator_addr": "cosmosaccaddr182ujqw3r8p5fffjqkf0rxzj29pg5q96nxd2khq",
              "shares": "8",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr13dre54wd7md0wg77x5fj5scj5u20e6stvdr829",
              "validator_addr": "cosmosaccaddr18thamkhnj9wz8pa4nhnp9rldprgant57ryzag7",
              "shares": "7",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr13dre54wd7md0wg77x5fj5scj5u20e6stvdr829",
              "validator_addr": "cosmosaccaddr12ceualfg92x7du73mtcv0zya4nxvq3tl2m52uz",
              "shares": "7",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr13lcax0d224sch59zz76m7n60nw0qfzty957l3s",
              "validator_addr": "cosmosaccaddr14sqp5mqtmfjdk9khae6lkfytwj9crtunpzr8f7",
              "shares": "200/9",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1j9w65cmt5s5kw9430jsv87yd4dgmmq6w60f3tl",
              "validator_addr": "cosmosaccaddr1j9w65cmt5s5kw9430jsv87yd4dgmmq6w60f3tl",
              "shares": "1",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1jtt6al0acr8h0uvq489rt9zx9lnau7rlcu30pt",
              "validator_addr": "cosmosaccaddr1jtt6al0acr8h0uvq489rt9zx9lnau7rlcu30pt",
              "shares": "150",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1jt77t00gucffqhrq7lh6vwmwzrwgw2fpvnhwk7",
              "validator_addr": "cosmosaccaddr1jt77t00gucffqhrq7lh6vwmwzrwgw2fpvnhwk7",
              "shares": "20",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1jdrxh34d00rm8ydclhxgkzmgp55y67m54umags",
              "validator_addr": "cosmosaccaddr1jdrxh34d00rm8ydclhxgkzmgp55y67m54umags",
              "shares": "10",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1jdehygp2p8z4f2wyf9t5q6f6mfqex66mfhfwdg",
              "validator_addr": "cosmosaccaddr1xvzw5nz4mp5s5cat7rancc248v5wpecrlqtw6n",
              "shares": "30",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1jw554408yw2h438q200jyuqgln76xh4ax0q4s0",
              "validator_addr": "cosmosaccaddr1jw554408yw2h438q200jyuqgln76xh4ax0q4s0",
              "shares": "100",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1jsch8k385shvse6j5dfx20483qym5uhq76xpjf",
              "validator_addr": "cosmosaccaddr1jsch8k385shvse6j5dfx20483qym5uhq76xpjf",
              "shares": "100",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1jkmn6asju47zuuzrf8rjt7sllaj5cx4kueyv8p",
              "validator_addr": "cosmosaccaddr1jkmn6asju47zuuzrf8rjt7sllaj5cx4kueyv8p",
              "shares": "150",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1jh3grawl62juvx5p8fz5xsy9hpw8w3mngqafe4",
              "validator_addr": "cosmosaccaddr1jh3grawl62juvx5p8fz5xsy9hpw8w3mngqafe4",
              "shares": "501",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1jc0nw5s49pxdjg2nh2ew7rm9r7c566mh5z4a93",
              "validator_addr": "cosmosaccaddr1ftrggumlvkesqny90par3gc9lp70yqhtlqucsk",
              "shares": "141",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1jck6gp4mqy33sk6a0fr5c8maq53hf4245v3mgg",
              "validator_addr": "cosmosaccaddr1jck6gp4mqy33sk6a0fr5c8maq53hf4245v3mgg",
              "shares": "150",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1je98ejgpvv2uzku3t6h5cydyr9hzva2cqhf3v5",
              "validator_addr": "cosmosaccaddr1je98ejgpvv2uzku3t6h5cydyr9hzva2cqhf3v5",
              "shares": "5",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1ja98tyxchsde6v3l02thp4fc35eymg88kts070",
              "validator_addr": "cosmosaccaddr1ja98tyxchsde6v3l02thp4fc35eymg88kts070",
              "shares": "8",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1n8yt3faem889lpr8x7f0pe0wth8003aclldlr3",
              "validator_addr": "cosmosaccaddr1xvt4e7xd0j9dwv2w83g50tpcltsl90h5dfnz6h",
              "shares": "1",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1n0l69ax723z2n060eeh357c0vsz2kagf06vscz",
              "validator_addr": "cosmosaccaddr1mvmdrtuxjttv90wssvh6xreeguvq8tksm789m5",
              "shares": "500/3",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1nnyel6v0kx8zmfh9edmre3ua4dt9306cfxsxgd",
              "validator_addr": "cosmosaccaddr1nnyel6v0kx8zmfh9edmre3ua4dt9306cfxsxgd",
              "shares": "100",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1na3wr7ujdp3qdej6l5y0k4znzrkaz77t2yjaqf",
              "validator_addr": "cosmosaccaddr1na3wr7ujdp3qdej6l5y0k4znzrkaz77t2yjaqf",
              "shares": "151",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr15pzn94awwmcff4m9cmju84t59t0uj5m65fwtl6",
              "validator_addr": "cosmosaccaddr18thamkhnj9wz8pa4nhnp9rldprgant57ryzag7",
              "shares": "10/3",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr15pzn94awwmcff4m9cmju84t59t0uj5m65fwtl6",
              "validator_addr": "cosmosaccaddr16rcrnftjyl2mctz78825ng8tx5ss22jf6jcp9l",
              "shares": "6",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr15rrxdflhq4744qt9m5ysstq3zpykhacf908eez",
              "validator_addr": "cosmosaccaddr15rrxdflhq4744qt9m5ysstq3zpykhacf908eez",
              "shares": "313",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr15yqm96jm0mhlluerr2ze0u369vqjdhfye2f9e9",
              "validator_addr": "cosmosaccaddr1640dsytftqn2uxk4prlfxjvn46nnp8wxt5s0ye",
              "shares": "10",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr15w2rengajq9js8hu57kjw88dly5vy7gsqedn0n",
              "validator_addr": "cosmosaccaddr15w2rengajq9js8hu57kjw88dly5vy7gsqedn0n",
              "shares": "150",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr15wkl5rq6uk8qdnpehdkp7clng255e06z2p9lfk",
              "validator_addr": "cosmosaccaddr15wkl5rq6uk8qdnpehdkp7clng255e06z2p9lfk",
              "shares": "20",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr15nhcasamcyav7gs34nezq0r0g9zyswrzh6gt3f",
              "validator_addr": "cosmosaccaddr1mvmdrtuxjttv90wssvh6xreeguvq8tksm789m5",
              "shares": "30",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr15mzck4rm5w55sfgr9kgexgtna22v608907t8y6",
              "validator_addr": "cosmosaccaddr12ceualfg92x7du73mtcv0zya4nxvq3tl2m52uz",
              "shares": "160",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr15mzck4rm5w55sfgr9kgexgtna22v608907t8y6",
              "validator_addr": "cosmosaccaddr10505nl7yftsme9jk2glhjhta7w0475uva87paj",
              "shares": "50",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr15mzck4rm5w55sfgr9kgexgtna22v608907t8y6",
              "validator_addr": "cosmosaccaddr15mzck4rm5w55sfgr9kgexgtna22v608907t8y6",
              "shares": "65",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr15u9ve7fz8fqaf7r2u3p4f9ru4ze47pau5cxgcg",
              "validator_addr": "cosmosaccaddr15u9ve7fz8fqaf7r2u3p4f9ru4ze47pau5cxgcg",
              "shares": "200",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr157eu0k3fvnk4d3rpgtyz2jujf8shlk8mxeqjkk",
              "validator_addr": "cosmosaccaddr157eu0k3fvnk4d3rpgtyz2jujf8shlk8mxeqjkk",
              "shares": "5",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr157mg9hnhchfrqvk3enrvmvj29yhmlwf759xrgw",
              "validator_addr": "cosmosaccaddr157mg9hnhchfrqvk3enrvmvj29yhmlwf759xrgw",
              "shares": "1150",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr14zfph0h8rexsca6gg6jkwqup3sgl6mwj6eu4e6",
              "validator_addr": "cosmosaccaddr14zfph0h8rexsca6gg6jkwqup3sgl6mwj6eu4e6",
              "shares": "100",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr14sqp5mqtmfjdk9khae6lkfytwj9crtunpzr8f7",
              "validator_addr": "cosmosaccaddr14sqp5mqtmfjdk9khae6lkfytwj9crtunpzr8f7",
              "shares": "2045/9",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1k9pqxd8fxqlk52uwfxnlsexqj6xnmw5swhss45",
              "validator_addr": "cosmosaccaddr1k9pqxd8fxqlk52uwfxnlsexqj6xnmw5swhss45",
              "shares": "1600/9",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1kwftgznylzerft2ksgkzvvfn5rfpy4nk2ye8na",
              "validator_addr": "cosmosaccaddr1kwftgznylzerft2ksgkzvvfn5rfpy4nk2ye8na",
              "shares": "149",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1k5ah4mpy8ykzp2mpcvapnyv6gemp35zuyhw8sy",
              "validator_addr": "cosmosaccaddr1xuwcmdvef5xudnyqa9gh5vqlq6ff0meu87zjxh",
              "shares": "10",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1k5ah4mpy8ykzp2mpcvapnyv6gemp35zuyhw8sy",
              "validator_addr": "cosmosaccaddr1k5ah4mpy8ykzp2mpcvapnyv6gemp35zuyhw8sy",
              "shares": "40",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1hqclw9xcymkhtrazzq9x576v9khhe6j9ja3jp0",
              "validator_addr": "cosmosaccaddr1h4q0xkdg30cl9vw0u8ejm0rs337dszy98gnd4a",
              "shares": "2",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1hydk58dtxr87ml34xqc07ph2h66yt2f0rrhyt8",
              "validator_addr": "cosmosaccaddr1hydk58dtxr87ml34xqc07ph2h66yt2f0rrhyt8",
              "shares": "50",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1hfygmryre7r8m9pfqmc85y7kw7ejphmp59k2x8",
              "validator_addr": "cosmosaccaddr1hfygmryre7r8m9pfqmc85y7kw7ejphmp59k2x8",
              "shares": "150",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1h4q0xkdg30cl9vw0u8ejm0rs337dszy98gnd4a",
              "validator_addr": "cosmosaccaddr1h4q0xkdg30cl9vw0u8ejm0rs337dszy98gnd4a",
              "shares": "140",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1h6el6v9rn0unkw2fh9s86juvg6fsvcmy8wy72s",
              "validator_addr": "cosmosaccaddr1qkllv8f6qakkw3hk9dqvytys090lk6twsyv8vf",
              "shares": "5",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1h6el6v9rn0unkw2fh9s86juvg6fsvcmy8wy72s",
              "validator_addr": "cosmosaccaddr1h6el6v9rn0unkw2fh9s86juvg6fsvcmy8wy72s",
              "shares": "19",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1hm5f52equcd4f268uxf4xhwljpnljvdyqal3dv",
              "validator_addr": "cosmosaccaddr1hm5f52equcd4f268uxf4xhwljpnljvdyqal3dv",
              "shares": "49/3",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1hmkg328mkqen9kc7z3j07hhu5ep9qh64q8kkfh",
              "validator_addr": "cosmosaccaddr1hmkg328mkqen9kc7z3j07hhu5ep9qh64q8kkfh",
              "shares": "11",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1h7yr9pm3ffhq5djx75ycmumvpwlxhulss3e2yv",
              "validator_addr": "cosmosaccaddr1h7yr9pm3ffhq5djx75ycmumvpwlxhulss3e2yv",
              "shares": "1340",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1cze0acxhmpx26awlc98h2v33ay95k232yenkqx",
              "validator_addr": "cosmosaccaddr1cze0acxhmpx26awlc98h2v33ay95k232yenkqx",
              "shares": "12",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1cypmdvszcd9kd3jenmqxd03cpceql8rtuvxftp",
              "validator_addr": "cosmosaccaddr1cypmdvszcd9kd3jenmqxd03cpceql8rtuvxftp",
              "shares": "100",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1cy5t4e62z4h7r74cn7vkylx77n5s4hds2xeud6",
              "validator_addr": "cosmosaccaddr1cy5t4e62z4h7r74cn7vkylx77n5s4hds2xeud6",
              "shares": "20",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1cyayv35tv47t829mel6jm0vqzmrdhf3jq87tqg",
              "validator_addr": "cosmosaccaddr1cyayv35tv47t829mel6jm0vqzmrdhf3jq87tqg",
              "shares": "100",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1ckr3s08lm6mddnueaffaplnn4eet3c4n6h8zzh",
              "validator_addr": "cosmosaccaddr1h6el6v9rn0unkw2fh9s86juvg6fsvcmy8wy72s",
              "shares": "10",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1cut9le04243e3fjq6jc8jysup8lpl3ngwa4ntx",
              "validator_addr": "cosmosaccaddr1vrc7zpg5teawwuzkfh6t7c6sy353sukhlarxpa",
              "shares": "1",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1c7r8w934puc05af2yssgwsdu3w4vgp6a5p2xaz",
              "validator_addr": "cosmosaccaddr1640dsytftqn2uxk4prlfxjvn46nnp8wxt5s0ye",
              "shares": "10",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1ezlvjcfq5s6fzrzp0rgrfudkte280dy528v93m",
              "validator_addr": "cosmosaccaddr1640dsytftqn2uxk4prlfxjvn46nnp8wxt5s0ye",
              "shares": "10",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1eyd6g0h5h96heegs42djrptlq9zflx3q2cndmp",
              "validator_addr": "cosmosaccaddr1y2z20pwqu5qpclque3pqkguruvheum2djtzjw3",
              "shares": "30",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1eyd6g0h5h96heegs42djrptlq9zflx3q2cndmp",
              "validator_addr": "cosmosaccaddr12ceualfg92x7du73mtcv0zya4nxvq3tl2m52uz",
              "shares": "10",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1eg88jsn38mrael6cqu7d2u8j6dynya7fv2p2tl",
              "validator_addr": "cosmosaccaddr1eg88jsn38mrael6cqu7d2u8j6dynya7fv2p2tl",
              "shares": "150",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1e078whe3xwhagy2hjjsht36tq3uehz3v2cs2n5",
              "validator_addr": "cosmosaccaddr1640dsytftqn2uxk4prlfxjvn46nnp8wxt5s0ye",
              "shares": "10",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1e3qm8jd9357zhdemwnaafmf0wy3f4yqmd307c2",
              "validator_addr": "cosmosaccaddr1e3qm8jd9357zhdemwnaafmf0wy3f4yqmd307c2",
              "shares": "392",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1e32c8x0dcgv0kz6qr7xceerne2gznt5s8e8tr7",
              "validator_addr": "cosmosaccaddr1e32c8x0dcgv0kz6qr7xceerne2gznt5s8e8tr7",
              "shares": "32",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1ejhdq2gpj20schz5n2r7lyxt6pwp2usahaf3wz",
              "validator_addr": "cosmosaccaddr15rrxdflhq4744qt9m5ysstq3zpykhacf908eez",
              "shares": "3",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1ehhkppuyz9mz9qeevqyvhrl2kyqrc0vuhz2436",
              "validator_addr": "cosmosaccaddr1ehhkppuyz9mz9qeevqyvhrl2kyqrc0vuhz2436",
              "shares": "20",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1emaa7mwgpnpmc7yptm728ytp9quamsvu9rk4hp",
              "validator_addr": "cosmosaccaddr1emaa7mwgpnpmc7yptm728ytp9quamsvu9rk4hp",
              "shares": "380",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1eu22vuernyw7yemd29lnkmd4x2l5grvx9kh6he",
              "validator_addr": "cosmosaccaddr1640dsytftqn2uxk4prlfxjvn46nnp8wxt5s0ye",
              "shares": "10",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr16rcrnftjyl2mctz78825ng8tx5ss22jf6jcp9l",
              "validator_addr": "cosmosaccaddr16rcrnftjyl2mctz78825ng8tx5ss22jf6jcp9l",
              "shares": "50500023/500000",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr16yxdmg8w5vc038cg0myh0w32s4h74ch59j67u6",
              "validator_addr": "cosmosaccaddr16yxdmg8w5vc038cg0myh0w32s4h74ch59j67u6",
              "shares": "3",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr16yk2su2ndeetfqvra5cwwcf3sk59yjagyz99a8",
              "validator_addr": "cosmosaccaddr1640dsytftqn2uxk4prlfxjvn46nnp8wxt5s0ye",
              "shares": "10",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr16xae9kqqllsjak3ntxmlgn4s3jmvvyeqcwzgen",
              "validator_addr": "cosmosaccaddr1yfseqtj5sjhzz2q2ym09jym4h4nc4yevae0jp2",
              "shares": "5",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr162k6wg0ufzy6pe079m3evt0tqrjfrrath5qw0d",
              "validator_addr": "cosmosaccaddr19uhnhct0p45ek6qxp3cjjrjtz4pacwcsgzvpuj",
              "shares": "5",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr16dfl4ruvt54qyhyf26fwmle62jnywp7mu9akt7",
              "validator_addr": "cosmosaccaddr16dfl4ruvt54qyhyf26fwmle62jnywp7mu9akt7",
              "shares": "20",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr160836cl5pqyy4mmk2k5u5v5xwcuztc8fksrk3k",
              "validator_addr": "cosmosaccaddr1jkmn6asju47zuuzrf8rjt7sllaj5cx4kueyv8p",
              "shares": "10",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr163emaq8ahwp25pmfqrjaz3c7wm9vtytcwmlphg",
              "validator_addr": "cosmosaccaddr163emaq8ahwp25pmfqrjaz3c7wm9vtytcwmlphg",
              "shares": "10",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1640dsytftqn2uxk4prlfxjvn46nnp8wxt5s0ye",
              "validator_addr": "cosmosaccaddr1640dsytftqn2uxk4prlfxjvn46nnp8wxt5s0ye",
              "shares": "10",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr164jntjfk9zs8x29mc27qansfwvjqs60gj6ermu",
              "validator_addr": "cosmosaccaddr164jntjfk9zs8x29mc27qansfwvjqs60gj6ermu",
              "shares": "980",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr16hvayg2lv7rmqz7rjsd3fnq3y34j6r97dxrecf",
              "validator_addr": "cosmosaccaddr16hvayg2lv7rmqz7rjsd3fnq3y34j6r97dxrecf",
              "shares": "5",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr16m99y8vsmyenxt4d93qft4gs0czzk6np80y775",
              "validator_addr": "cosmosaccaddr1hfygmryre7r8m9pfqmc85y7kw7ejphmp59k2x8",
              "shares": "1",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1mx36cmexzz6dxntf6lxpsdsf8u7kuzerh07m4s",
              "validator_addr": "cosmosaccaddr1fskmzyt2hr8usr3s65dq3rfur3fy6g2hjp23tm",
              "shares": "100",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1mtez787cl6phgkylneq9wjj3d6letn0247lhf5",
              "validator_addr": "cosmosaccaddr1mtez787cl6phgkylneq9wjj3d6letn0247lhf5",
              "shares": "1030",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1mvqtcm5a8pf8q5xv50ag6k445032fz7qaax5ge",
              "validator_addr": "cosmosaccaddr157mg9hnhchfrqvk3enrvmvj29yhmlwf759xrgw",
              "shares": "2",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1mvsaaudvvhkfxnkjev3r9g73nrjejvdfc36zz9",
              "validator_addr": "cosmosaccaddr1fskmzyt2hr8usr3s65dq3rfur3fy6g2hjp23tm",
              "shares": "15",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1mvsaaudvvhkfxnkjev3r9g73nrjejvdfc36zz9",
              "validator_addr": "cosmosaccaddr157eu0k3fvnk4d3rpgtyz2jujf8shlk8mxeqjkk",
              "shares": "20/9",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1mvsaaudvvhkfxnkjev3r9g73nrjejvdfc36zz9",
              "validator_addr": "cosmosaccaddr164jntjfk9zs8x29mc27qansfwvjqs60gj6ermu",
              "shares": "5",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1mvmdrtuxjttv90wssvh6xreeguvq8tksm789m5",
              "validator_addr": "cosmosaccaddr1mvmdrtuxjttv90wssvh6xreeguvq8tksm789m5",
              "shares": "20",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1mwmcpq5nhaen8zyy86mrrept2gmp0z5peapkhu",
              "validator_addr": "cosmosaccaddr1mwmcpq5nhaen8zyy86mrrept2gmp0z5peapkhu",
              "shares": "147",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1m0u78hks3hcf5c6szvlly7x40dkattyal0wy2w",
              "validator_addr": "cosmosaccaddr185n9jln2g4w79w39m0g85x2tns6uekqg99tq7u",
              "shares": "30",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1m7rtacq9ygng2zf2y4z3mer8hdtacg4hyspmrh",
              "validator_addr": "cosmosaccaddr1vrc7zpg5teawwuzkfh6t7c6sy353sukhlarxpa",
              "shares": "1",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1uffeh2wq0lrvwm7z77c595v040q4kk4g3v8vte",
              "validator_addr": "cosmosaccaddr1uffeh2wq0lrvwm7z77c595v040q4kk4g3v8vte",
              "shares": "100",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1u374sn4c4arc8nsnpqdaqm2s5n70gpg6xvw4ny",
              "validator_addr": "cosmosaccaddr1640dsytftqn2uxk4prlfxjvn46nnp8wxt5s0ye",
              "shares": "10",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1u40w3tq6w9cycgdpy4s692e92k03wnt8824zhp",
              "validator_addr": "cosmosaccaddr12ceualfg92x7du73mtcv0zya4nxvq3tl2m52uz",
              "shares": "3",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1uknhjkw775zp5xkztc4clxr240ut59ek5seerx",
              "validator_addr": "cosmosaccaddr1uknhjkw775zp5xkztc4clxr240ut59ek5seerx",
              "shares": "50",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1u6ddcsjueax884l3tfrs66497c7g86sk3vfmqj",
              "validator_addr": "cosmosaccaddr1u6ddcsjueax884l3tfrs66497c7g86sk3vfmqj",
              "shares": "2725",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1umaajfgap5ef6yxvk5706kwk8j08l7wh6h9fp2",
              "validator_addr": "cosmosaccaddr1pmntq5en2rgtr5rzr4e304efrve4lr43z32y5s",
              "shares": "500",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1umaajfgap5ef6yxvk5706kwk8j08l7wh6h9fp2",
              "validator_addr": "cosmosaccaddr1zpx36km7tk5cksyzxgvcp6g552p3uhwx84r53q",
              "shares": "1000",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1umaajfgap5ef6yxvk5706kwk8j08l7wh6h9fp2",
              "validator_addr": "cosmosaccaddr1y2z20pwqu5qpclque3pqkguruvheum2djtzjw3",
              "shares": "1000",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1umaajfgap5ef6yxvk5706kwk8j08l7wh6h9fp2",
              "validator_addr": "cosmosaccaddr1xuwcmdvef5xudnyqa9gh5vqlq6ff0meu87zjxh",
              "shares": "1500",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1umaajfgap5ef6yxvk5706kwk8j08l7wh6h9fp2",
              "validator_addr": "cosmosaccaddr185n9jln2g4w79w39m0g85x2tns6uekqg99tq7u",
              "shares": "1000",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1umaajfgap5ef6yxvk5706kwk8j08l7wh6h9fp2",
              "validator_addr": "cosmosaccaddr1jh3grawl62juvx5p8fz5xsy9hpw8w3mngqafe4",
              "shares": "500",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1umaajfgap5ef6yxvk5706kwk8j08l7wh6h9fp2",
              "validator_addr": "cosmosaccaddr1na3wr7ujdp3qdej6l5y0k4znzrkaz77t2yjaqf",
              "shares": "1000",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1umaajfgap5ef6yxvk5706kwk8j08l7wh6h9fp2",
              "validator_addr": "cosmosaccaddr1kwftgznylzerft2ksgkzvvfn5rfpy4nk2ye8na",
              "shares": "1000",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1umaajfgap5ef6yxvk5706kwk8j08l7wh6h9fp2",
              "validator_addr": "cosmosaccaddr1hfygmryre7r8m9pfqmc85y7kw7ejphmp59k2x8",
              "shares": "1090",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1umaajfgap5ef6yxvk5706kwk8j08l7wh6h9fp2",
              "validator_addr": "cosmosaccaddr1cypmdvszcd9kd3jenmqxd03cpceql8rtuvxftp",
              "shares": "1000",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1umaajfgap5ef6yxvk5706kwk8j08l7wh6h9fp2",
              "validator_addr": "cosmosaccaddr1umaajfgap5ef6yxvk5706kwk8j08l7wh6h9fp2",
              "shares": "1150",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1ul3ef4trvjfmgaptu70ed4fjlky2aa49t4pfle",
              "validator_addr": "cosmosaccaddr1ul3ef4trvjfmgaptu70ed4fjlky2aa49t4pfle",
              "shares": "9487/9",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1uluae4kaxq3wxd698t256a8luv2tf0ngy40qp3",
              "validator_addr": "cosmosaccaddr1uluae4kaxq3wxd698t256a8luv2tf0ngy40qp3",
              "shares": "1",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1artqm53zt7j7znm8cpgqzgp9jhlx0xrtle2pjr",
              "validator_addr": "cosmosaccaddr1artqm53zt7j7znm8cpgqzgp9jhlx0xrtle2pjr",
              "shares": "10",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1a8ppz2vh0zjfh8vaqckp67ct0ajm9crv0t3v0g",
              "validator_addr": "cosmosaccaddr164jntjfk9zs8x29mc27qansfwvjqs60gj6ermu",
              "shares": "10",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1agfqwpcxpl8s8pl3uhxhw7h0r05an60dmlgukx",
              "validator_addr": "cosmosaccaddr1agfqwpcxpl8s8pl3uhxhw7h0r05an60dmlgukx",
              "shares": "1",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1aaumwh9rc9ksh7dwuxux43jd93e7t4u88g5kg4",
              "validator_addr": "cosmosaccaddr1aaumwh9rc9ksh7dwuxux43jd93e7t4u88g5kg4",
              "shares": "104",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr17rqsh3fw6rehnwzarz90jkqtt5fupmh50gy556",
              "validator_addr": "cosmosaccaddr17rqsh3fw6rehnwzarz90jkqtt5fupmh50gy556",
              "shares": "156",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr17gvlvfpsfl6hffn5u2hahk22un4ynpykc44tat",
              "validator_addr": "cosmosaccaddr17gvlvfpsfl6hffn5u2hahk22un4ynpykc44tat",
              "shares": "700",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr170yylzwpzx9c6er2mn086xtdy80w2dmwp9wa3c",
              "validator_addr": "cosmosaccaddr1t9nxlu834au3u7mltyf8udzw7kzsf0wntx7e4w",
              "shares": "1",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr170yylzwpzx9c6er2mn086xtdy80w2dmwp9wa3c",
              "validator_addr": "cosmosaccaddr1dq7jt3pn7mrce2twac907jue7hjd0p0rgt3qnq",
              "shares": "2",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr17n0y36k4430006yc6m2r709t9uxk73yptwsrsg",
              "validator_addr": "cosmosaccaddr1hm5f52equcd4f268uxf4xhwljpnljvdyqal3dv",
              "shares": "100/9",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr17nlusdvrk34fq65jemy3umfjfwaxfzv4asyl60",
              "validator_addr": "cosmosaccaddr17nlusdvrk34fq65jemy3umfjfwaxfzv4asyl60",
              "shares": "150",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1lzy0ztj9cz7lemlev99u0upulqjd5cu6nk3t2n",
              "validator_addr": "cosmosaccaddr1lzy0ztj9cz7lemlev99u0upulqjd5cu6nk3t2n",
              "shares": "95/9",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1l9jt8xgejkcm4sulhwj8c83ftprz5p9lyq4605",
              "validator_addr": "cosmosaccaddr1jkmn6asju47zuuzrf8rjt7sllaj5cx4kueyv8p",
              "shares": "50",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1l9jt8xgejkcm4sulhwj8c83ftprz5p9lyq4605",
              "validator_addr": "cosmosaccaddr1l9jt8xgejkcm4sulhwj8c83ftprz5p9lyq4605",
              "shares": "100",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1lvmxcdhua2kt3qd86wnnej5m6gm4max4um5sx2",
              "validator_addr": "cosmosaccaddr1lvmxcdhua2kt3qd86wnnej5m6gm4max4um5sx2",
              "shares": "3",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1lw2d0qljstqzp6fqhsd5jazackgad9y2e3u2n8",
              "validator_addr": "cosmosaccaddr1lw2d0qljstqzp6fqhsd5jazackgad9y2e3u2n8",
              "shares": "5",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1l0qw5znfd6e8pshpjvyghjjzyr4l6ymla080lt",
              "validator_addr": "cosmosaccaddr1xvt4e7xd0j9dwv2w83g50tpcltsl90h5dfnz6h",
              "shares": "24",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1l0qw5znfd6e8pshpjvyghjjzyr4l6ymla080lt",
              "validator_addr": "cosmosaccaddr1l0qw5znfd6e8pshpjvyghjjzyr4l6ymla080lt",
              "shares": "1150",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1l3agyg568dsk2r9v7f35wge3yr85ampza7cdfj",
              "validator_addr": "cosmosaccaddr1l3agyg568dsk2r9v7f35wge3yr85ampza7cdfj",
              "shares": "5026/9",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1lhjta6nt0lewj05m8444tuyhalkrffgpm7njpp",
              "validator_addr": "cosmosaccaddr1lhjta6nt0lewj05m8444tuyhalkrffgpm7njpp",
              "shares": "100",
              "height": "0"
            },
            {
              "delegator_addr": "cosmosaccaddr1lmazkxl9dtetn5lxzl45ut3ywwn94xvs3l4z8f",
              "validator_addr": "cosmosaccaddr1lmazkxl9dtetn5lxzl45ut3ywwn94xvs3l4z8f",
              "shares": "90",
              "height": "0"
            }
          ]
        }
      }
    }
  }
}`,
	"health.json": `
{
  "jsonrpc": "2.0",
  "id": "",
  "result": {}
}`,
	"net_info.json": `
{
  "jsonrpc": "2.0",
  "id": "",
  "result": {
    "listening": true,
    "listeners": [
      "Listener(@172.31.10.143:26656)"
    ],
    "n_peers": "10",
    "peers": [
      {
        "node_info": {
          "id": "8e44e63e3b536dd780feff44fc83596c30affb7d",
          "listen_addr": "133.130.90.146:26656",
          "network": "gaia-8001",
          "version": "0.23.1",
          "channels": "4020212223303800",
          "moniker": "kamon-moniker",
          "other": [
            "amino_version=0.12.0",
            "p2p_version=0.5.0",
            "consensus_version=v1/0.2.2",
            "rpc_version=0.7.0/3",
            "tx_index=on",
            "rpc_addr=tcp://0.0.0.0:26657"
          ]
        },
        "is_outbound": true,
        "connection_status": {
          "Duration": "9223372036854775807",
          "SendMonitor": {
            "Active": true,
            "Start": "2018-10-23T10:12:44.46Z",
            "Duration": "236407020000000",
            "Idle": "0",
            "Bytes": "2569860566",
            "Samples": "1161271",
            "InstRate": "9480",
            "CurRate": "12721",
            "AvgRate": "10870",
            "PeakRate": "296950",
            "BytesRem": "0",
            "TimeRem": "0",
            "Progress": 0
          },
          "RecvMonitor": {
            "Active": true,
            "Start": "2018-10-23T10:12:44.46Z",
            "Duration": "236407000000000",
            "Idle": "0",
            "Bytes": "2427249978",
            "Samples": "943873",
            "InstRate": "64688",
            "CurRate": "10020",
            "AvgRate": "10267",
            "PeakRate": "402230",
            "BytesRem": "0",
            "TimeRem": "0",
            "Progress": 0
          },
          "Channels": [
            {
              "ID": 48,
              "SendQueueCapacity": "1",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "0"
            },
            {
              "ID": 64,
              "SendQueueCapacity": "1000",
              "SendQueueSize": "0",
              "Priority": "10",
              "RecentlySent": "0"
            },
            {
              "ID": 32,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "6828"
            },
            {
              "ID": 33,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "10",
              "RecentlySent": "32005"
            },
            {
              "ID": 34,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "55416"
            },
            {
              "ID": 35,
              "SendQueueCapacity": "2",
              "SendQueueSize": "0",
              "Priority": "1",
              "RecentlySent": "25"
            },
            {
              "ID": 56,
              "SendQueueCapacity": "1",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "0"
            },
            {
              "ID": 0,
              "SendQueueCapacity": "10",
              "SendQueueSize": "0",
              "Priority": "1",
              "RecentlySent": "0"
            }
          ]
        }
      },
      {
        "node_info": {
          "id": "7043e77f36609425116882066e6b0d1cf6a78185",
          "listen_addr": "209.250.244.51:26656",
          "network": "gaia-8001",
          "version": "0.23.1",
          "channels": "40202122233038",
          "moniker": "shensi-cosmos-validator",
          "other": [
            "amino_version=0.12.0",
            "p2p_version=0.5.0",
            "consensus_version=v1/0.2.2",
            "rpc_version=0.7.0/3",
            "tx_index=on",
            "rpc_addr=tcp://0.0.0.0:26657"
          ]
        },
        "is_outbound": true,
        "connection_status": {
          "Duration": "9223372036854775807",
          "SendMonitor": {
            "Active": true,
            "Start": "2018-10-24T14:16:44.26Z",
            "Duration": "135367220000000",
            "Idle": "0",
            "Bytes": "1139095391",
            "Samples": "632038",
            "InstRate": "7650",
            "CurRate": "12332",
            "AvgRate": "8415",
            "PeakRate": "265870",
            "BytesRem": "0",
            "TimeRem": "0",
            "Progress": 0
          },
          "RecvMonitor": {
            "Active": true,
            "Start": "2018-10-24T14:16:44.26Z",
            "Duration": "135367200000000",
            "Idle": "40000000",
            "Bytes": "1484561914",
            "Samples": "568522",
            "InstRate": "141260",
            "CurRate": "13918",
            "AvgRate": "10967",
            "PeakRate": "357060",
            "BytesRem": "0",
            "TimeRem": "0",
            "Progress": 0
          },
          "Channels": [
            {
              "ID": 48,
              "SendQueueCapacity": "1",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "0"
            },
            {
              "ID": 64,
              "SendQueueCapacity": "1000",
              "SendQueueSize": "0",
              "Priority": "10",
              "RecentlySent": "0"
            },
            {
              "ID": 32,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "7032"
            },
            {
              "ID": 33,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "10",
              "RecentlySent": "22241"
            },
            {
              "ID": 34,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "40435"
            },
            {
              "ID": 35,
              "SendQueueCapacity": "2",
              "SendQueueSize": "0",
              "Priority": "1",
              "RecentlySent": "12"
            },
            {
              "ID": 56,
              "SendQueueCapacity": "1",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "0"
            },
            {
              "ID": 0,
              "SendQueueCapacity": "10",
              "SendQueueSize": "0",
              "Priority": "1",
              "RecentlySent": "0"
            }
          ]
        }
      },
      {
        "node_info": {
          "id": "9c5c946df2cab55d111d1fa7e6b03c50666160c4",
          "listen_addr": "142.93.31.31:26656",
          "network": "gaia-8001",
          "version": "0.23.1",
          "channels": "4020212223303800",
          "moniker": "4llfr33d0m_Sentry_A",
          "other": [
            "amino_version=0.12.0",
            "p2p_version=0.5.0",
            "consensus_version=v1/0.2.2",
            "rpc_version=0.7.0/3",
            "tx_index=on",
            "rpc_addr=tcp://0.0.0.0:26657"
          ]
        },
        "is_outbound": true,
        "connection_status": {
          "Duration": "9223372036854775807",
          "SendMonitor": {
            "Active": true,
            "Start": "2018-10-24T18:20:44.22Z",
            "Duration": "120727260000000",
            "Idle": "0",
            "Bytes": "912567815",
            "Samples": "560179",
            "InstRate": "7650",
            "CurRate": "12173",
            "AvgRate": "7559",
            "PeakRate": "355080",
            "BytesRem": "0",
            "TimeRem": "0",
            "Progress": 0
          },
          "RecvMonitor": {
            "Active": true,
            "Start": "2018-10-24T18:20:44.22Z",
            "Duration": "120727280000000",
            "Idle": "0",
            "Bytes": "1307121526",
            "Samples": "485700",
            "InstRate": "4610",
            "CurRate": "11855",
            "AvgRate": "10827",
            "PeakRate": "321740",
            "BytesRem": "0",
            "TimeRem": "0",
            "Progress": 0
          },
          "Channels": [
            {
              "ID": 48,
              "SendQueueCapacity": "1",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "0"
            },
            {
              "ID": 64,
              "SendQueueCapacity": "1000",
              "SendQueueSize": "0",
              "Priority": "10",
              "RecentlySent": "0"
            },
            {
              "ID": 32,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "7091"
            },
            {
              "ID": 33,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "10",
              "RecentlySent": "26668"
            },
            {
              "ID": 34,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "39041"
            },
            {
              "ID": 35,
              "SendQueueCapacity": "2",
              "SendQueueSize": "0",
              "Priority": "1",
              "RecentlySent": "0"
            },
            {
              "ID": 56,
              "SendQueueCapacity": "1",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "0"
            },
            {
              "ID": 0,
              "SendQueueCapacity": "10",
              "SendQueueSize": "0",
              "Priority": "1",
              "RecentlySent": "0"
            }
          ]
        }
      },
      {
        "node_info": {
          "id": "d346857fed996ffa1a034f4c983e06eaf9c00664",
          "listen_addr": "35.180.139.254:26656",
          "network": "gaia-8001",
          "version": "0.23.1",
          "channels": "4020212223303800",
          "moniker": "delegate-tomorrow",
          "other": [
            "amino_version=0.12.0",
            "p2p_version=0.5.0",
            "consensus_version=v1/0.2.2",
            "rpc_version=0.7.0/3",
            "tx_index=on",
            "rpc_addr=tcp://0.0.0.0:26657"
          ]
        },
        "is_outbound": true,
        "connection_status": {
          "Duration": "9223372036854775807",
          "SendMonitor": {
            "Active": true,
            "Start": "2018-10-25T07:28:14.26Z",
            "Duration": "73477220000000",
            "Idle": "0",
            "Bytes": "593804845",
            "Samples": "342826",
            "InstRate": "7650",
            "CurRate": "12241",
            "AvgRate": "8081",
            "PeakRate": "244670",
            "BytesRem": "0",
            "TimeRem": "0",
            "Progress": 0
          },
          "RecvMonitor": {
            "Active": true,
            "Start": "2018-10-25T07:28:14.26Z",
            "Duration": "73477200000000",
            "Idle": "40000000",
            "Bytes": "822625465",
            "Samples": "335076",
            "InstRate": "14800",
            "CurRate": "12636",
            "AvgRate": "11196",
            "PeakRate": "366120",
            "BytesRem": "0",
            "TimeRem": "0",
            "Progress": 0
          },
          "Channels": [
            {
              "ID": 48,
              "SendQueueCapacity": "1",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "0"
            },
            {
              "ID": 64,
              "SendQueueCapacity": "1000",
              "SendQueueSize": "0",
              "Priority": "10",
              "RecentlySent": "0"
            },
            {
              "ID": 32,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "7039"
            },
            {
              "ID": 33,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "10",
              "RecentlySent": "26666"
            },
            {
              "ID": 34,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "37382"
            },
            {
              "ID": 35,
              "SendQueueCapacity": "2",
              "SendQueueSize": "0",
              "Priority": "1",
              "RecentlySent": "0"
            },
            {
              "ID": 56,
              "SendQueueCapacity": "1",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "0"
            },
            {
              "ID": 0,
              "SendQueueCapacity": "10",
              "SendQueueSize": "0",
              "Priority": "1",
              "RecentlySent": "0"
            }
          ]
        }
      },
      {
        "node_info": {
          "id": "4a53d0a92deb68a0085aa0b7f100fc475af42b94",
          "listen_addr": "50.17.110.27:26656",
          "network": "gaia-8001",
          "version": "0.23.1",
          "channels": "4020212223303800",
          "moniker": "zemya",
          "other": [
            "amino_version=0.12.0",
            "p2p_version=0.5.0",
            "consensus_version=v1/0.2.2",
            "rpc_version=0.7.0/3",
            "tx_index=on",
            "rpc_addr=tcp://0.0.0.0:26657"
          ]
        },
        "is_outbound": true,
        "connection_status": {
          "Duration": "9223372036854775807",
          "SendMonitor": {
            "Active": true,
            "Start": "2018-10-23T18:33:44.02Z",
            "Duration": "206347460000000",
            "Idle": "0",
            "Bytes": "1240156209",
            "Samples": "938412",
            "InstRate": "9480",
            "CurRate": "12117",
            "AvgRate": "6010",
            "PeakRate": "255360",
            "BytesRem": "0",
            "TimeRem": "0",
            "Progress": 0
          },
          "RecvMonitor": {
            "Active": true,
            "Start": "2018-10-23T18:33:44.02Z",
            "Duration": "206347460000000",
            "Idle": "20000000",
            "Bytes": "2158464491",
            "Samples": "892283",
            "InstRate": "60732",
            "CurRate": "13203",
            "AvgRate": "10460",
            "PeakRate": "338780",
            "BytesRem": "0",
            "TimeRem": "0",
            "Progress": 0
          },
          "Channels": [
            {
              "ID": 48,
              "SendQueueCapacity": "1",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "0"
            },
            {
              "ID": 64,
              "SendQueueCapacity": "1000",
              "SendQueueSize": "0",
              "Priority": "10",
              "RecentlySent": "0"
            },
            {
              "ID": 32,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "7201"
            },
            {
              "ID": 33,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "10",
              "RecentlySent": "24793"
            },
            {
              "ID": 34,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "29877"
            },
            {
              "ID": 35,
              "SendQueueCapacity": "2",
              "SendQueueSize": "0",
              "Priority": "1",
              "RecentlySent": "37"
            },
            {
              "ID": 56,
              "SendQueueCapacity": "1",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "0"
            },
            {
              "ID": 0,
              "SendQueueCapacity": "10",
              "SendQueueSize": "0",
              "Priority": "1",
              "RecentlySent": "0"
            }
          ]
        }
      },
      {
        "node_info": {
          "id": "292c1eb49a6e3d88c48adb8e42a9c1e61488cb78",
          "listen_addr": "54.215.221.213:26656",
          "network": "gaia-8001",
          "version": "0.23.1",
          "channels": "4020212223303800",
          "moniker": "bianjie-sentry",
          "other": [
            "amino_version=0.12.0",
            "p2p_version=0.5.0",
            "consensus_version=v1/0.2.2",
            "rpc_version=0.7.0/3",
            "tx_index=on",
            "rpc_addr=tcp://0.0.0.0:26657"
          ]
        },
        "is_outbound": true,
        "connection_status": {
          "Duration": "9223372036854775807",
          "SendMonitor": {
            "Active": true,
            "Start": "2018-10-23T10:24:14.14Z",
            "Duration": "235717340000000",
            "Idle": "0",
            "Bytes": "1718699289",
            "Samples": "1077945",
            "InstRate": "7650",
            "CurRate": "1511",
            "AvgRate": "7291",
            "PeakRate": "305560",
            "BytesRem": "0",
            "TimeRem": "0",
            "Progress": 0
          },
          "RecvMonitor": {
            "Active": true,
            "Start": "2018-10-23T10:24:14.14Z",
            "Duration": "235717300000000",
            "Idle": "60000000",
            "Bytes": "2597831392",
            "Samples": "976319",
            "InstRate": "1877",
            "CurRate": "11208",
            "AvgRate": "11021",
            "PeakRate": "451020",
            "BytesRem": "0",
            "TimeRem": "0",
            "Progress": 0
          },
          "Channels": [
            {
              "ID": 48,
              "SendQueueCapacity": "1",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "0"
            },
            {
              "ID": 64,
              "SendQueueCapacity": "1000",
              "SendQueueSize": "0",
              "Priority": "10",
              "RecentlySent": "0"
            },
            {
              "ID": 32,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "7122"
            },
            {
              "ID": 33,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "10",
              "RecentlySent": "18584"
            },
            {
              "ID": 34,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "36128"
            },
            {
              "ID": 35,
              "SendQueueCapacity": "2",
              "SendQueueSize": "0",
              "Priority": "1",
              "RecentlySent": "3"
            },
            {
              "ID": 56,
              "SendQueueCapacity": "1",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "0"
            },
            {
              "ID": 0,
              "SendQueueCapacity": "10",
              "SendQueueSize": "0",
              "Priority": "1",
              "RecentlySent": "0"
            }
          ]
        }
      },
      {
        "node_info": {
          "id": "e81a57e3275340174d007636d78a28a238cbba46",
          "listen_addr": "192.241.154.136:26656",
          "network": "gaia-8001",
          "version": "0.23.1",
          "channels": "4020212223303800",
          "moniker": "gogoc2",
          "other": [
            "amino_version=0.12.0",
            "p2p_version=0.5.0",
            "consensus_version=v1/0.2.2",
            "rpc_version=0.7.0/3",
            "tx_index=on",
            "rpc_addr=tcp://0.0.0.0:26657"
          ]
        },
        "is_outbound": true,
        "connection_status": {
          "Duration": "9223372036854775807",
          "SendMonitor": {
            "Active": true,
            "Start": "2018-10-23T10:23:14.14Z",
            "Duration": "235777340000000",
            "Idle": "0",
            "Bytes": "1699478070",
            "Samples": "1095268",
            "InstRate": "7650",
            "CurRate": "11950",
            "AvgRate": "7208",
            "PeakRate": "360360",
            "BytesRem": "0",
            "TimeRem": "0",
            "Progress": 0
          },
          "RecvMonitor": {
            "Active": true,
            "Start": "2018-10-23T10:23:14.14Z",
            "Duration": "235777340000000",
            "Idle": "20000000",
            "Bytes": "2390987333",
            "Samples": "962567",
            "InstRate": "34470",
            "CurRate": "14575",
            "AvgRate": "10141",
            "PeakRate": "401500",
            "BytesRem": "0",
            "TimeRem": "0",
            "Progress": 0
          },
          "Channels": [
            {
              "ID": 48,
              "SendQueueCapacity": "1",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "0"
            },
            {
              "ID": 64,
              "SendQueueCapacity": "1000",
              "SendQueueSize": "0",
              "Priority": "10",
              "RecentlySent": "0"
            },
            {
              "ID": 32,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "7129"
            },
            {
              "ID": 33,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "10",
              "RecentlySent": "30410"
            },
            {
              "ID": 34,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "33180"
            },
            {
              "ID": 35,
              "SendQueueCapacity": "2",
              "SendQueueSize": "0",
              "Priority": "1",
              "RecentlySent": "37"
            },
            {
              "ID": 56,
              "SendQueueCapacity": "1",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "0"
            },
            {
              "ID": 0,
              "SendQueueCapacity": "10",
              "SendQueueSize": "0",
              "Priority": "1",
              "RecentlySent": "0"
            }
          ]
        }
      },
      {
        "node_info": {
          "id": "303481d4f855f227fb576c765ed88a15b67150b9",
          "listen_addr": "188.1.89.170:26656",
          "network": "gaia-8001",
          "version": "0.23.1",
          "channels": "4020212223303800",
          "moniker": "i-1",
          "other": [
            "amino_version=0.12.0",
            "p2p_version=0.5.0",
            "consensus_version=v1/0.2.2",
            "rpc_version=0.7.0/3",
            "tx_index=on",
            "rpc_addr=tcp://0.0.0.0:26657"
          ]
        },
        "is_outbound": true,
        "connection_status": {
          "Duration": "9223372036854775807",
          "SendMonitor": {
            "Active": true,
            "Start": "2018-10-23T10:27:14.52Z",
            "Duration": "235536960000000",
            "Idle": "0",
            "Bytes": "2343079299",
            "Samples": "1116088",
            "InstRate": "9480",
            "CurRate": "1693",
            "AvgRate": "9948",
            "PeakRate": "283840",
            "BytesRem": "0",
            "TimeRem": "0",
            "Progress": 0
          },
          "RecvMonitor": {
            "Active": true,
            "Start": "2018-10-23T10:27:14.52Z",
            "Duration": "235536980000000",
            "Idle": "0",
            "Bytes": "2584387876",
            "Samples": "1019882",
            "InstRate": "5142",
            "CurRate": "10779",
            "AvgRate": "10972",
            "PeakRate": "329583",
            "BytesRem": "0",
            "TimeRem": "0",
            "Progress": 0
          },
          "Channels": [
            {
              "ID": 48,
              "SendQueueCapacity": "1",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "0"
            },
            {
              "ID": 64,
              "SendQueueCapacity": "1000",
              "SendQueueSize": "0",
              "Priority": "10",
              "RecentlySent": "0"
            },
            {
              "ID": 32,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "6836"
            },
            {
              "ID": 33,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "10",
              "RecentlySent": "18132"
            },
            {
              "ID": 34,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "50083"
            },
            {
              "ID": 35,
              "SendQueueCapacity": "2",
              "SendQueueSize": "0",
              "Priority": "1",
              "RecentlySent": "24"
            },
            {
              "ID": 56,
              "SendQueueCapacity": "1",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "0"
            },
            {
              "ID": 0,
              "SendQueueCapacity": "10",
              "SendQueueSize": "0",
              "Priority": "1",
              "RecentlySent": "0"
            }
          ]
        }
      },
      {
        "node_info": {
          "id": "e179345924be168def105e64daff369c27e7c47e",
          "listen_addr": "104.237.6.136:26656",
          "network": "gaia-8001",
          "version": "0.23.1",
          "channels": "4020212223303800",
          "moniker": "nuevax",
          "other": [
            "amino_version=0.12.0",
            "p2p_version=0.5.0",
            "consensus_version=v1/0.2.2",
            "rpc_version=0.7.0/3",
            "tx_index=on",
            "rpc_addr=tcp://0.0.0.0:26657"
          ]
        },
        "is_outbound": true,
        "connection_status": {
          "Duration": "9223372036854775807",
          "SendMonitor": {
            "Active": true,
            "Start": "2018-10-23T11:12:44.08Z",
            "Duration": "232807400000000",
            "Idle": "0",
            "Bytes": "1820642856",
            "Samples": "1084061",
            "InstRate": "9480",
            "CurRate": "12268",
            "AvgRate": "7820",
            "PeakRate": "405160",
            "BytesRem": "0",
            "TimeRem": "0",
            "Progress": 0
          },
          "RecvMonitor": {
            "Active": true,
            "Start": "2018-10-23T11:12:44.08Z",
            "Duration": "232807400000000",
            "Idle": "20000000",
            "Bytes": "2436977974",
            "Samples": "1019913",
            "InstRate": "65780",
            "CurRate": "13157",
            "AvgRate": "10468",
            "PeakRate": "388340",
            "BytesRem": "0",
            "TimeRem": "0",
            "Progress": 0
          },
          "Channels": [
            {
              "ID": 48,
              "SendQueueCapacity": "1",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "0"
            },
            {
              "ID": 64,
              "SendQueueCapacity": "1000",
              "SendQueueSize": "0",
              "Priority": "10",
              "RecentlySent": "0"
            },
            {
              "ID": 32,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "7179"
            },
            {
              "ID": 33,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "10",
              "RecentlySent": "32305"
            },
            {
              "ID": 34,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "40514"
            },
            {
              "ID": 35,
              "SendQueueCapacity": "2",
              "SendQueueSize": "0",
              "Priority": "1",
              "RecentlySent": "0"
            },
            {
              "ID": 56,
              "SendQueueCapacity": "1",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "0"
            },
            {
              "ID": 0,
              "SendQueueCapacity": "10",
              "SendQueueSize": "0",
              "Priority": "1",
              "RecentlySent": "0"
            }
          ]
        }
      },
      {
        "node_info": {
          "id": "1d20090454966d7c5b74f9442b1d39e9e8511ca9",
          "listen_addr": "91.90.148.194:26656",
          "network": "gaia-8001",
          "version": "0.23.0",
          "channels": "4020212223303800",
          "moniker": "Raichu_Numba_One",
          "other": [
            "amino_version=0.12.0",
            "p2p_version=0.5.0",
            "consensus_version=v1/0.2.2",
            "rpc_version=0.7.0/3",
            "tx_index=on",
            "rpc_addr=tcp://0.0.0.0:26657"
          ]
        },
        "is_outbound": true,
        "connection_status": {
          "Duration": "9223372036854775807",
          "SendMonitor": {
            "Active": true,
            "Start": "2018-10-23T20:07:14.32Z",
            "Duration": "200737160000000",
            "Idle": "0",
            "Bytes": "1745775996",
            "Samples": "941940",
            "InstRate": "5820",
            "CurRate": "12286",
            "AvgRate": "8697",
            "PeakRate": "254880",
            "BytesRem": "0",
            "TimeRem": "0",
            "Progress": 0
          },
          "RecvMonitor": {
            "Active": true,
            "Start": "2018-10-23T20:07:14.32Z",
            "Duration": "200737180000000",
            "Idle": "60000000",
            "Bytes": "2239663505",
            "Samples": "911929",
            "InstRate": "24157",
            "CurRate": "13707",
            "AvgRate": "11157",
            "PeakRate": "285540",
            "BytesRem": "0",
            "TimeRem": "0",
            "Progress": 0
          },
          "Channels": [
            {
              "ID": 48,
              "SendQueueCapacity": "1",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "0"
            },
            {
              "ID": 64,
              "SendQueueCapacity": "1000",
              "SendQueueSize": "0",
              "Priority": "10",
              "RecentlySent": "0"
            },
            {
              "ID": 32,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "7020"
            },
            {
              "ID": 33,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "10",
              "RecentlySent": "30502"
            },
            {
              "ID": 34,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "37969"
            },
            {
              "ID": 35,
              "SendQueueCapacity": "2",
              "SendQueueSize": "0",
              "Priority": "1",
              "RecentlySent": "25"
            },
            {
              "ID": 56,
              "SendQueueCapacity": "1",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "0"
            },
            {
              "ID": 0,
              "SendQueueCapacity": "10",
              "SendQueueSize": "0",
              "Priority": "1",
              "RecentlySent": "0"
            }
          ]
        }
      }
    ]
  }
}`,
	"num_unconfirmed_txs.json": `
{
  "jsonrpc": "2.0",
  "id": "",
  "result": {
    "n_txs": "0",
    "txs": null
  }
}`,
	"status.json": `
{
  "jsonrpc": "2.0",
  "id": "",
  "result": {
    "node_info": {
      "id": "55acd26ae3304d676a14cc6e819fec18fa49a2e6",
      "listen_addr": "172.31.10.143:26656",
      "network": "gaia-8001",
      "version": "0.23.1",
      "channels": "4020212223303800",
      "moniker": "q-validator",
      "other": [
        "amino_version=0.12.0",
        "p2p_version=0.5.0",
        "consensus_version=v1/0.2.2",
        "rpc_version=0.7.0/3",
        "tx_index=on",
        "rpc_addr=tcp://0.0.0.0:26657"
      ]
    },
    "sync_info": {
      "latest_block_hash": "5AE4E524C5291B555B1D07748BB53831B0E5152B",
      "latest_app_hash": "9979166E98DE9BA1A401E7EF4CDCF0F6A6D030B9",
      "latest_block_height": "870167",
      "latest_block_time": "2018-10-26T03:52:43.640425408Z",
      "catching_up": false
    },
    "validator_info": {
      "address": "E56E00A6B6AC1BE546CE9A9F34C7C1BB7585A8C1",
      "pub_key": {
        "type": "tendermint/PubKeyEd25519",
        "value": "TBsp7W9VwXn2YhLL5Dn9Ihpgy4dr+5bgT3uy5g+fjtQ="
      },
      "voting_power": "74"
    }
  }
}`,
	"subscribe.json": `
{
  "jsonrpc": "2.0",
  "id": "",
  "error": {
    "code": -32601,
    "message": "Method not found"
  }
}`,
	"tx.json": `
{
  "jsonrpc": "2.0",
  "id": "",
  "error": {
    "code": -32603,
    "message": "Internal error",
    "data": "Transaction hash cannot be empty"
  }
}`,
	"tx_search.json": `
{
  "jsonrpc": "2.0",
  "id": "",
  "error": {
    "code": -32603,
    "message": "Internal error",
    "data": "\nparse error near Unknown (line 1 symbol 1 - line 1 symbol 1):\n\"\"\n"
  }
}`,
	"unconfirmed_txs.json": `
{
  "jsonrpc": "2.0",
  "id": "",
  "result": {
    "n_txs": "0",
    "txs": []
  }
}`,
	"unsubscribe.json": `
{
  "jsonrpc": "2.0",
  "id": "",
  "error": {
    "code": -32601,
    "message": "Method not found"
  }
}`,
	"unsubscribe_all.json": `
{
  "jsonrpc": "2.0",
  "id": "",
  "error": {
    "code": -32601,
    "message": "Method not found"
  }
}`,
	"validators.json": `
{
  "jsonrpc": "2.0",
  "id": "",
  "result": {
    "block_height": "870169",
    "validators": [
      {
        "address": "0077F68359FE12C678980AC3442460CB1CC1ABC4",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "zf+aS3anGNm9CrpliDTqpUlHCBpi0qZknUo+aYOXC+s="
        },
        "voting_power": "169",
        "accum": "4184"
      },
      {
        "address": "0081A939ABBB733E8DC36E406EC51E4309B9AE3A",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "Sk7VmvaRgatGyVrbkbQ4QBCeda/DIUs6E3NzstWQV6g="
        },
        "voting_power": "154",
        "accum": "20471"
      },
      {
        "address": "01F78669F9515FD83DF9250F5C0EE143D3DAD65C",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "1K9/Z+oiK380Obf+LYVN6P7ydFIWCFvNQpT+ToRUURg="
        },
        "voting_power": "627",
        "accum": "-19075"
      },
      {
        "address": "0548261C50222FD710EB5EBDF03A0E6567B21D20",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "8K3clCjVU33BTIpUhdahGmu++WxHj4NUE9krCRkk++s="
        },
        "voting_power": "1181",
        "accum": "-8386"
      },
      {
        "address": "066C70AF46E5E1B473BD125FD8937EF90E555641",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "To+MjLVGM4gNuE7JahfDwIc5Q+31H+Q60Cy4AoU2fjE="
        },
        "voting_power": "5",
        "accum": "10458"
      },
      {
        "address": "0A692DBFFE9E5DD28E7DECC33CED1AEB4C0D014E",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "N3K5kDdfcKJurfaa6s2zfKgtYvz1Pagz7VWi9ZfX8yM="
        },
        "voting_power": "439",
        "accum": "4317"
      },
      {
        "address": "0DDF97111D73DB825138EC15EC27DE5FE8536901",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "H0SIA/BU6Xj8oT5bQkvLpEITN3CqFLbMeBcQ72NZrAE="
        },
        "voting_power": "18",
        "accum": "42"
      },
      {
        "address": "1ECD8C5FE1341D74CDC633A245DD56858DE2F9A2",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "lUMCRAeu47BsOhNvCQTQJQeB68z0/VaElC9j5gDt9y8="
        },
        "voting_power": "578",
        "accum": "-7594"
      },
      {
        "address": "20FBB777AB2676E19A31E2EB35B2C7B426C99C13",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "/DtRehBSochyxchfAn/Tb4tbwC9u6Ryq++VeBE72FuA="
        },
        "voting_power": "39",
        "accum": "-18309"
      },
      {
        "address": "296869AB66F7003A5AADA1A8DEEEA2913486668F",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "4cI0avx/jOL9eWe7PyRm4Mhv1QSSrhxpyr30f7tJ7i0="
        },
        "voting_power": "10",
        "accum": "13800"
      },
      {
        "address": "29B93E4EECD8C591AB76C2D52FD7C43CBEEEC50B",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "GTwvM3OOWuXdwH7suTJIWdYdtwUU1hLAHIXmpdzGfl0="
        },
        "voting_power": "7",
        "accum": "4772"
      },
      {
        "address": "2E6AF0D5B1A85E173F5EBEDA93D7E7D97A88D06C",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "G3YJxcM2oMCW8R5s/HSqmnudsQjcj+e/vpZBtMyob1w="
        },
        "voting_power": "121",
        "accum": "-13055"
      },
      {
        "address": "33412C9A69C60CB0FAA24A5C5B3A906DE24B47C2",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "4DEDoU/RsHMPS54GzgwkWnW5zPQfMt9aInFFc3GyfA8="
        },
        "voting_power": "40",
        "accum": "-10800"
      },
      {
        "address": "3363E8F97B02ECC00289E72173D827543047ACDA",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "mPnu910hOOa1tAQ7pbOLFDxvllbQUmrbtGjqQrYg1nM="
        },
        "voting_power": "89",
        "accum": "-21933"
      },
      {
        "address": "36E9B7FAEC964F97D84C028ED62493E373AD0B38",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "ViJzrw9Unk1XyDg4f0UQLFE4zfUaFh23wZZY0f31Qg0="
        },
        "voting_power": "139",
        "accum": "-11676"
      },
      {
        "address": "398610C4CF11C84C89AC6975470972EE75DA17E4",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "5GAxIHzu06l0n3MU8wNQrCcrrnpZR9EKe5qZMsM2h/E="
        },
        "voting_power": "490",
        "accum": "5961"
      },
      {
        "address": "3AC99B708404452366E0953A4BE849226DC86CBC",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "ISIM341M+EUYOMpwhn9gyCvRUZ06xomp/+lOwa50meQ="
        },
        "voting_power": "480",
        "accum": "12142"
      },
      {
        "address": "3E899BA36A93C7370347BD5FE85909F2ECD93F8C",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "M2Iu/A8RQcmoiIEar/iV4wBjVj0I0gcd/nNnvwmnT2g="
        },
        "voting_power": "39",
        "accum": "15732"
      },
      {
        "address": "411C430038318FE4DB32AFC508A6C12782161294",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "LmJN98a+ce+shthBz95qKsdwUwWRs1QdBLO4HGNQDpM="
        },
        "voting_power": "107",
        "accum": "-20586"
      },
      {
        "address": "4944F8BBBB5303E17BC0607DF332320EDBC750EA",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "A6GzeXUM3vsXaDAEYMSDgSKkqn9AoUYjs8empH46MGY="
        },
        "voting_power": "1288",
        "accum": "12222"
      },
      {
        "address": "4B9C6FF8B36EE6134AA8D822772C0162C471186D",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "g0f/B/FxRDU0VbS7vMXs8cJwSUITpbtn3NjEQ0NJ5Fg="
        },
        "voting_power": "20",
        "accum": "-25208"
      },
      {
        "address": "50C18305F49BE332E4C6BCDE0DEECD4A22B1D6A4",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "PX8v6nM+dk8fnqQWuXaRpRrdlhL+sMUHYo75CJbnY9U="
        },
        "voting_power": "109",
        "accum": "20676"
      },
      {
        "address": "5104F488EA6B1E2C33CBDA3C18389C88104DBC72",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "JREG9eji9gfJR2b61qMu8jNZ835ezPY2A4o+zjYMfDg="
        },
        "voting_power": "299",
        "accum": "-7998"
      },
      {
        "address": "570B6A755ED339EAC56BB3390A48069F6AC9449B",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "5XCMLGNTGkMh9/JsUzi/8Xl7Y0cdDe38stX+zvBLXuE="
        },
        "voting_power": "44",
        "accum": "-20574"
      },
      {
        "address": "5B9628695CDFC6025857578114E3D3126687EAE1",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "NCwjL8K9R2CLOMYub3MhAkkB0fktT0mZC75TEtcXyQ4="
        },
        "voting_power": "130",
        "accum": "1290"
      },
      {
        "address": "5BDF8EA41B1835FBAD865B181A2A4B23F511B5E4",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "RBuRRNEzA9RA1Wrdi9PPFQJ29/n/bqN9O2tQv9Gq248="
        },
        "voting_power": "6040",
        "accum": "-3200"
      },
      {
        "address": "5DE9674DA2D535D5B1547D90990E9087EAE15E39",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "yRk1EO5Ta0iCI68BjluGEyZIlkBTGvUzbl+s/Rujhco="
        },
        "voting_power": "5",
        "accum": "10458"
      },
      {
        "address": "5E06FB65FD0E0EBB11D0D74A1E480B4A8D7D0FE7",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "iPPJziJLaL5QA6Mr6uaUGwoNrGts45W6LVIKbo11GfA="
        },
        "voting_power": "41",
        "accum": "8302"
      },
      {
        "address": "5E8673673E37450F01B64138FBF4B172BDB52D68",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "CuPin/hLM2tK6KvcUlOi/Brm5xNi3zq6UPgpLFPSzt8="
        },
        "voting_power": "776",
        "accum": "-10667"
      },
      {
        "address": "651995D8471D07A22C149265F5237A89149795C1",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "WX9FldKAX82ALt3u3v80v/1zLUNJQO689C9vwXooDuQ="
        },
        "voting_power": "142",
        "accum": "-6569"
      },
      {
        "address": "668309BE1A6F5317197B0BC6E47153D249DD5DE9",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "DKXu53utgN9Q/4dd7uRC0c8j1bNHTGSVDzhlEcrRFl8="
        },
        "voting_power": "103",
        "accum": "-11063"
      },
      {
        "address": "69D99B2C66043ACBEAA8447525C356AFC6408E0C",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "8pfpbIxBBiu88hpxS3CeRpv7kClEjl8SwVgckDNBGlE="
        },
        "voting_power": "541",
        "accum": "16467"
      },
      {
        "address": "69EDADFA4A803E9DBF548C900D4A465E7D154262",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "LRK6flgMsHrBYkIIahegEkzlwC7gQ3dAyw0jcAd1u/k="
        },
        "voting_power": "107",
        "accum": "-20637"
      },
      {
        "address": "6B275FFA0571EEFC4482CD7A6CAE836F5785CB4E",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "pggq0WUTlsiA8cBiptRgxlw4WUG2GmXfLQxkbMbCBPk="
        },
        "voting_power": "202",
        "accum": "-11850"
      },
      {
        "address": "6D3AD380E398D8592B03AA124BAFED5F125DF863",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "ENAVynNXVpj/IdYx9kCPKaPs4bWSxRIHNlmS9QiDuZQ="
        },
        "voting_power": "54",
        "accum": "-1118"
      },
      {
        "address": "6F6C93B8BF3495063FEC9B5C771158607F4AE9D6",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "S8s6fdAQNQ3bN9SNVAsHB/j8uv1CM1roxeLesL+fh4g="
        },
        "voting_power": "141",
        "accum": "-16876"
      },
      {
        "address": "71BC71D28E4599BCA22214BBB2CE0DFEE90A9EE4",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "MWID5oHOb367w/js26+cbQdyGqw2KEpe967D44z/eSQ="
        },
        "voting_power": "1",
        "accum": "-24920"
      },
      {
        "address": "732CEEF54C374DDC6ADECBFD707AEFD07FEDC143",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "HjSC7VkhKih6xMhudlqfaFE8ZZnP8RKJPv4iqR7RhcE="
        },
        "voting_power": "357",
        "accum": "22740"
      },
      {
        "address": "7A20421F99FC4362B13C2FD43548F3E75B057017",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "HWtL6kXsdXL6Lxoq4XXGG3FkpPAmDBFwrP1ndYCvsnA="
        },
        "voting_power": "20",
        "accum": "-5112"
      },
      {
        "address": "7A32D912F7A790EEE5610653DC4DDBC0DB4F2497",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "d5o/ErYmwkS/Cw06SH1l3vrBDXRUJLtjj8kcz9CEiFQ="
        },
        "voting_power": "79",
        "accum": "-17571"
      },
      {
        "address": "7EB7593F519F1180C7E872D0494611464EDF03AE",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "VakMQSPBEuSC9Nwuv8WWhrZVUmH31bUR4+G6pJhkgE8="
        },
        "voting_power": "1807",
        "accum": "-6004"
      },
      {
        "address": "807CF675BA8479FC4FEC1047215DF07970040F1F",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "2p8s/pRBZPjYWKKMlR7AOXypDzDmPo762iXlKpCwtco="
        },
        "voting_power": "14",
        "accum": "10513"
      },
      {
        "address": "8340F82D32C4A221C57FC33FCBA44309C4DF1BBD",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "FSj9jMpCOdWfX0rCQ0TpnEa8/uOroQl4V4Bk3Msck3w="
        },
        "voting_power": "65",
        "accum": "-4064"
      },
      {
        "address": "83AB321E012C57746689FED3A6064DF331F81007",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "aUViYC2znC55sleHfmsIN9hZ45SbYPbDcYA0gVzglsc="
        },
        "voting_power": "946",
        "accum": "7357"
      },
      {
        "address": "850D387A41ADB21F7D67EAC5076D594DF3D861C7",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "Sowu0HYQbGFJ4dA3urnWwXqMc1+6kzQHre+9/iVlpjg="
        },
        "voting_power": "10",
        "accum": "-7944"
      },
      {
        "address": "85335B5077FAE8E2926AD338634079124EE1C4A1",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "lGnYgu+rv9HZj3Y+fS6Ad/GszZgtwwyZcM9mAc/tuu4="
        },
        "voting_power": "475",
        "accum": "1600"
      },
      {
        "address": "86DDBEEF35D35786A68834B34869174AC314BB74",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "ilDFePqWkNRNWUEoWKq42jZC5CD4rglfs7+gCuq6Qc8="
        },
        "voting_power": "1867",
        "accum": "23972"
      },
      {
        "address": "89B357643A945BD433CB5B2B9CBD9B90CB9CCD5E",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "UVRp+iVQUv9erXJwqfI6ClWIXk/8fSs0CRudMMU49Zo="
        },
        "voting_power": "36",
        "accum": "-631"
      },
      {
        "address": "8DF81F7913F1EBF856618B1475C12EEEC1214BB0",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "0WFwqm+9/VGNlZDcAlkqCdQm/8M+W19LWnzXIDg/py0="
        },
        "voting_power": "221",
        "accum": "9198"
      },
      {
        "address": "8EA62D00960B7557560E0D78278B6B36165AD0A8",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "43IDu79dbTK5jDg1x2xR09pzWgWdAGHoIrz4Hk9H//Q="
        },
        "voting_power": "10",
        "accum": "-10014"
      },
      {
        "address": "8F06FE1FD042683C426137FBC4BBAA4288924217",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "lB+CgSbTzpHXJHUU15fJIPaPl6XF0nSfdNxBnlKJqTk="
        },
        "voting_power": "2466",
        "accum": "5809"
      },
      {
        "address": "8F302C08318444DDC43EABDE96E869D0B1F5A421",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "+yk80I86yZ+wqwf/9iDEfewwZGDaXCKosDuWZnlSf4M="
        },
        "voting_power": "110",
        "accum": "110"
      },
      {
        "address": "927CC759310F4DF5C8936EF6D666ED3135D609B7",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "tkdOxR2CIOctyc5e9IMGwnMwB420OTN3Rfb6fUMVhr8="
        },
        "voting_power": "11",
        "accum": "-5829"
      },
      {
        "address": "9524C2BBF5C9BCC7AFEE121AF03E31D8480DD5E5",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "nWft/3AE/5/Iv5ePNUbkFHeG8HpNnnIFA3hQKR6qRmQ="
        },
        "voting_power": "5",
        "accum": "10458"
      },
      {
        "address": "96BDE45EC52606CF5CB59853097871AD60680549",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "4JoJuRfaANhdM1x3AWRo1/Cj9DH3VA+fi1SynzknV+w="
        },
        "voting_power": "19",
        "accum": "-25888"
      },
      {
        "address": "989E892A603E6BA7F325BC52555DFB47A8751E66",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "zkQOeG4/QylfQCqxV8WPzC0bUvIkRb467gbL80m4/KA="
        },
        "voting_power": "10",
        "accum": "-25489"
      },
      {
        "address": "9D93AD02FE1BD7E76B9E58D7EB461088EAD011E5",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "jj0Y/Fy8JSJR3g+PHU6Ce0ecYwHGUVJ4bVyR7WwcyLI="
        },
        "voting_power": "614",
        "accum": "-10125"
      },
      {
        "address": "A0ED8A257D7862094FE8015BCE2135D111722DF6",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "dvLH3HPQ3YH0oDefmGTPvpwn/gTFhvBAQiYNV4DVs5k="
        },
        "voting_power": "195",
        "accum": "-16478"
      },
      {
        "address": "A5868971300F155020634D7BB49C7634F7AACCD3",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "cCoFsZzKZ9SQZbHe4NueVObIezP6ts0tRTZ/aN96dig="
        },
        "voting_power": "18",
        "accum": "40"
      },
      {
        "address": "AC96444A9083890B7C7B430DB17BF1D3EA5EED04",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "QzxYamaSVNx+Idlb9bCcYsPI4Oa/zQ+rx4MN+XA2xLI="
        },
        "voting_power": "649",
        "accum": "16704"
      },
      {
        "address": "AFCFB9B8DF23559CB5EB97C7CF1D77B411B2F47F",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "hWd0eUfhxpAqDZrKTgXcZJZXv5R92mQ4JJvZcAab0Ic="
        },
        "voting_power": "34",
        "accum": "4781"
      },
      {
        "address": "B0155252D73B7EEB74D2A8CC814397E66970A839",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "w3rKE+tQoLK8G+XPmjn+NszCk07iQ0sWaBbN5hQZcBY="
        },
        "voting_power": "124",
        "accum": "4513"
      },
      {
        "address": "C39A8A2B85198DCB8F23AD1F6C198E2BA7AB5D60",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "R/3f7VruxWpu+2hiHlVpplTwoOou5kfQI1k/6/9H/y8="
        },
        "voting_power": "181",
        "accum": "18459"
      },
      {
        "address": "C65E1E696C038A9A4A3565014ED168E97E7E63DF",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "pdxc69pGnaiK9oF+a49PjCnNww2PLYRBWlPFW/X6HNY="
        },
        "voting_power": "10",
        "accum": "-25489"
      },
      {
        "address": "D31B1198AE055BC52C600C19751011D95A64298A",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "NErIb5Hsp4ai02anNfdYui0Lm38uno+YzVJXO3xWS2g="
        },
        "voting_power": "7571",
        "accum": "-21258"
      },
      {
        "address": "D6285C3422C8886445D53757CBE12D2B91B5A9D0",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "NQX4yKpOztKrmgBhGIC5WOALOLOq3LTpbzsN4ZLXGec="
        },
        "voting_power": "50",
        "accum": "9297"
      },
      {
        "address": "D77FB84A729972CD3E757C9BEB61561855832B92",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "UjTvuOew2EaooduJBiYmBWeF5ai0yFJG8uio5YXpJgg="
        },
        "voting_power": "68",
        "accum": "10833"
      },
      {
        "address": "D981CB6DDF336E311A33C585C23C10F9F033A305",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "0ytofYCtFJT6zOYl0PMNqHan4cUwZg6ZnXHksICYfJ0="
        },
        "voting_power": "161",
        "accum": "-21253"
      },
      {
        "address": "DA38013D7BCC3EC46109F8CDDF14F8D080CC4A9D",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "q+y8mjbjNKBC3HSWgVC8HVU34QC/DQDxtuSmx9eL4qI="
        },
        "voting_power": "14",
        "accum": "10513"
      },
      {
        "address": "DB71F2DC8C967F0BABE9EF797760F17EF3D99ACE",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "u7qa9ST9HyzuL+xj3GyEoC+uBD52lVogDZRka26XxRI="
        },
        "voting_power": "52",
        "accum": "4684"
      },
      {
        "address": "DBA70FA7E9D55E035AD87B41C4DC0C38511FD09A",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "3O/S1+hdkUi5BxAohHHkxytw7S+tukjby46dRr6VhPE="
        },
        "voting_power": "3648",
        "accum": "-3397"
      },
      {
        "address": "DDDC63E5B9C7D1509805601E2EF6934BAEAC9115",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "pOYVpP34JYpnOMUbDzfbFHvLA2vfgJ8n/eknvqPmBS4="
        },
        "voting_power": "28",
        "accum": "20119"
      },
      {
        "address": "E39B5778E0D4297A46F45CF6940B0A163A430FF8",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "0HqB2x6x5HzeozpHatePECw07x1UcDdSz8kQGNznnA8="
        },
        "voting_power": "1768",
        "accum": "1826"
      },
      {
        "address": "E3B19738F7132F950668504D326C7660F159D170",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "tOEqjO2t51PEgO9Tv0B7qM0yPmy1n5tMa3Beg0tp3ns="
        },
        "voting_power": "33",
        "accum": "-16036"
      },
      {
        "address": "E42F8AEABB685219F7027F54D2B0DBC7846F633A",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "eeImG09hOPo1W7j7lKepN/Lx6I9GGHqVBVEKmznxACc="
        },
        "voting_power": "240",
        "accum": "-17760"
      },
      {
        "address": "E4E6C2E81269DFFA196C17F98240E491479A6008",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "ydjx2ea+PVuChrny6X2dluJwyXta+BsNQRsgHXp8fXw="
        },
        "voting_power": "19",
        "accum": "-25888"
      },
      {
        "address": "E50ADECD5FD27FA2CD610C07F8AED36A2FC4A6A6",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "BaaCxmYHKJ6obIzTCdRtjw1cc8d2mUJcMbLWCjf1aLo="
        },
        "voting_power": "7171",
        "accum": "-16099"
      },
      {
        "address": "E56E00A6B6AC1BE546CE9A9F34C7C1BB7585A8C1",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "TBsp7W9VwXn2YhLL5Dn9Ihpgy4dr+5bgT3uy5g+fjtQ="
        },
        "voting_power": "74",
        "accum": "20644"
      },
      {
        "address": "EE67A5CA3269434AFAA41CE43E17B467488B5210",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "CEQGeQBsQNFXg9XpExD3cfD8ucyiAB5vGw6o/Ux76q8="
        },
        "voting_power": "179",
        "accum": "-22345"
      },
      {
        "address": "F23FF36BD5B90C33CE3A03ED72DBDCF5EC07D6AF",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "2JoNf1gavJ1d6XFIumO1Mki5GVMOcg58AioHksU3maE="
        },
        "voting_power": "199",
        "accum": "19927"
      },
      {
        "address": "F4268597DE93E3E2CED38845010F8AA34E2A879E",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "Epd2FDKZwDybzT38Z7WB0y2jiLn9/2OLzmY3Zu18l6I="
        },
        "voting_power": "325",
        "accum": "18792"
      },
      {
        "address": "F6738260186D33D9C14FC6E7017AFE6BB952A63D",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "2qtEBT+Tc+SD2wJsdrVMHXrBKfvesxtmtSKDK5fXwA0="
        },
        "voting_power": "25",
        "accum": "4041"
      }
    ]
  }
}`,
}
