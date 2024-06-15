# attributes

start here:

https://play.max.com/movie/127b00c5-0131-4bac-b2d1-40762deefe09

then:

https://play.max.com/video/watch/b3b1410a-0c85-457b-bcc7-e13299bea2a8/1623fe4c-ef6e-4dd1-a10c-4a181f5f6579

then:

~~~
GET https://default.any-amer.prd.api.discomax.com/cms/routes/video/watch/b3b1410a-0c85-457b-bcc7-e13299bea2a8/1623fe4c-ef6e-4dd1-a10c-4a181f5f6579?include=default HTTP/2.0
authorization: Bearer eyJhbGciOiJSUzI1NiJ9.eyJqdGkiOiJ0b2tlbi0yNTIzYWEyNC1jNzU...
~~~

movie:

~~~json
{
  "data": {
    "attributes": {
      "canonical": true,
      "url": "/video/watch/b3b1410a-0c85-457b-bcc7-e13299bea2a8/1623fe4c-ef6e-4dd1-a10c-4a181f5f6579"
    },
    "id": "c65fc42b4caa4d290d40a9c3fcc192a4adf064b9870b677aa8aab2f8510072f1",
    "relationships": {
      "target": {
        "data": {
          "id": "215925083836394013750542481199525881753",
          "type": "page"
        }
      }
    },
    "type": "route"
  },
  "included": [
    {
      "attributes": {
        "airDate": "2011-04-01T00:01:00Z",
        "alternateId": "b3b1410a-0c85-457b-bcc7-e13299bea2a8",
        "description": "A soldier is transported into another man’s body aboard a commuter train, where he must foil a terrorist bombing plot within eight minutes.",
        "firstAvailableDate": "2024-04-01T07:01:00Z",
        "isFamilyContent": true,
        "isKidsContent": false,
        "kidsContent": false,
        "longDescription": "A soldier is transported into another man's body on a commuter train, where he must foil a terrorist plot within eight minutes. Learning more about the experimental military technology being deployed, he races against time to save the train passengers.",
        "materialType": "MOVIE",
        "name": "Source Code",
        "originalName": "Source Code",
        "secondaryTitle": "Source Code",
        "videoType": "MOVIE"
      },
      "id": "b3b1410a-0c85-457b-bcc7-e13299bea2a8",
      "relationships": {
        "contentPackages": {
          "data": [
            {
              "id": "VodEntertainment",
              "type": "package"
            },
            {
              "id": "Hbo",
              "type": "package"
            }
          ]
        },
        "creditGroups": {
          "data": [
            {
              "id": "b3b1410a-0c85-457b-bcc7-e13299bea2a8-credit-group-starring",
              "type": "creditGroup"
            },
            {
              "id": "b3b1410a-0c85-457b-bcc7-e13299bea2a8-credit-group-director",
              "type": "creditGroup"
            },
            {
              "id": "b3b1410a-0c85-457b-bcc7-e13299bea2a8-credit-group-writers",
              "type": "creditGroup"
            },
            {
              "id": "b3b1410a-0c85-457b-bcc7-e13299bea2a8-credit-group-producers",
              "type": "creditGroup"
            }
          ]
        },
        "edit": {
          "data": {
            "id": "1623fe4c-ef6e-4dd1-a10c-4a181f5f6579",
            "type": "edit"
          }
        },
        "images": {
          "data": [
            {
              "id": "i:b",
              "type": "image"
            },
            {
              "id": "i:c",
              "type": "image"
            },
            {
              "id": "i:d",
              "type": "image"
            }
          ]
        },
        "primaryChannel": {
          "data": {
            "id": "c0d1f27a-e2f8-4b3c-bf3c-ed0c4e258093",
            "type": "channel"
          }
        },
        "ratingDescriptors": {
          "data": [
            {
              "id": "4062c194-0c80-4ceb-9de5-d7f7e6e302b8",
              "type": "contentDescriptor"
            },
            {
              "id": "582e60b9-b1f7-4398-9182-a8758856209d",
              "type": "contentDescriptor"
            },
            {
              "id": "f8e36188-362f-4a7b-9934-c7d77cb86e64",
              "type": "contentDescriptor"
            }
          ]
        },
        "ratings": {
          "data": [
            {
              "id": "4d1bf7aa-ddc4-4641-b568-11f59abb4787",
              "type": "contentRating"
            }
          ]
        },
        "show": {
          "data": {
            "id": "127b00c5-0131-4bac-b2d1-40762deefe09",
            "type": "show"
          }
        },
        "txGenres": {
          "data": [
            {
              "id": "0a34f929-1e10-4e53-ba8d-d0e2dbedfdfd",
              "type": "taxonomyNode"
            },
            {
              "id": "9df43ec7-cc1a-4d8f-b7f1-d024810aed66",
              "type": "taxonomyNode"
            },
            {
              "id": "b676b596-7a71-44fb-8222-ecd28becc276",
              "type": "taxonomyNode"
            },
            {
              "id": "cdc65854-4013-4c85-b2b5-c48290ad3b5c",
              "type": "taxonomyNode"
            }
          ]
        }
      },
      "type": "video"
    }
  ],
  "meta": {
    "site": {
      "attributes": {
        "brandId": "beam",
        "mainTerritoryCode": "us",
        "theme": "beam",
        "websiteHostName": ""
      },
      "id": "beam_us",
      "type": "site"
    }
  }
}
~~~

episode:

~~~json
{
  "data": {
    "attributes": {
      "canonical": true,
      "url": "/video/watch/fbdd33a2-1189-4b9a-8c10-13244fb21b7f/6cc15a42-130f-4531-807a-b2c147d8ac68"
    },
    "id": "83b758280df91c09cd2e3fd2e2a2013b90959854ca2ae772de1caa9942578e26",
    "relationships": {
      "target": {
        "data": {
          "id": "215925083836394013750542481199525881753",
          "type": "page"
        }
      }
    },
    "type": "route"
  },
  "included": [
    {
      "attributes": {
        "alternateId": "818c3d9d-1831-48a6-9583-0364a7f98453",
        "description": "James Gandolfini stars in this acclaimed series about a mob boss whose professional and private strains land him in therapy.",
        "firstAvailableDate": "2023-02-08T05:00:00Z",
        "isFamilyContent": false,
        "isFavorite": false,
        "isKidsContent": false,
        "kidsContent": false,
        "longDescription": "James Gandolfini stars in this acclaimed series about a mob boss whose professional and private strains land him in therapy.",
        "name": "The Sopranos",
        "numberOfNewEpisodes": 0,
        "originalName": "The Sopranos",
        "premiereDate": "1999-01-21T00:01:00Z",
        "showType": "SERIES"
      },
      "id": "818c3d9d-1831-48a6-9583-0364a7f98453",
      "relationships": {
        "primaryChannel": {
          "data": {
            "id": "c0d1f27a-e2f8-4b3c-bf3c-ed0c4e258093",
            "type": "channel"
          }
        },
        "ratings": {
          "data": [
            {
              "id": "fbdd00c3-6c8f-4e27-be17-ac5997ebdfd2",
              "type": "contentRating"
            }
          ]
        },
        "routes": {
          "data": [
            {
              "id": "ed01df35aae7ca31febef7d3b20dcc7539a52ea2814ab61e2430d95aabef922f",
              "type": "route"
            }
          ]
        },
        "txGenres": {
          "data": [
            {
              "id": "1ef95565-d5d5-4793-9656-31207932bfe1",
              "type": "taxonomyNode"
            },
            {
              "id": "3e183de4-3108-4fb5-801f-66b9ca405f94",
              "type": "taxonomyNode"
            },
            {
              "id": "435e27a6-3026-42fb-8549-b560f7616360",
              "type": "taxonomyNode"
            },
            {
              "id": "8557a482-557e-4bec-bf53-814094c40392",
              "type": "taxonomyNode"
            },
            {
              "id": "cdc65854-4013-4c85-b2b5-c48290ad3b5c",
              "type": "taxonomyNode"
            }
          ]
        }
      },
      "type": "show"
    },
    {
      "attributes": {
        "airDate": "2000-01-01T00:01:00Z",
        "alternateId": "fbdd33a2-1189-4b9a-8c10-13244fb21b7f",
        "description": "Tony travels to Italy to jump-start a car-importing \"business\" and recruits a new lieutenant named Furio.",
        "episodeNumber": 4,
        "firstAvailableDate": "2023-02-18T05:00:00Z",
        "isFamilyContent": false,
        "isFavorite": false,
        "isKidsContent": false,
        "kidsContent": false,
        "longDescription": "Tony travels to Italy to jump-start a car-importing 'business' and recruits a new lieutenant named Furio.",
        "materialType": "EPISODE",
        "name": "Commendatori",
        "originalName": "Commendatori",
        "seasonNumber": 2,
        "secondaryTitle": "Commendatori",
        "videoType": "EPISODE",
        "viewingHistory": {
          "completed": false,
          "lastReportedTimestamp": "2024-06-14T00:14:28.083Z",
          "position": 51,
          "viewed": true
        }
      },
      "id": "fbdd33a2-1189-4b9a-8c10-13244fb21b7f",
      "relationships": {
        "contentPackages": {
          "data": [
            {
              "id": "VodEntertainment",
              "type": "package"
            },
            {
              "id": "Hbo",
              "type": "package"
            }
          ]
        },
        "creditGroups": {
          "data": [
            {
              "id": "fbdd33a2-1189-4b9a-8c10-13244fb21b7f-credit-group-starring",
              "type": "creditGroup"
            },
            {
              "id": "fbdd33a2-1189-4b9a-8c10-13244fb21b7f-credit-group-director",
              "type": "creditGroup"
            },
            {
              "id": "fbdd33a2-1189-4b9a-8c10-13244fb21b7f-credit-group-writers",
              "type": "creditGroup"
            },
            {
              "id": "fbdd33a2-1189-4b9a-8c10-13244fb21b7f-credit-group-producers",
              "type": "creditGroup"
            },
            {
              "id": "fbdd33a2-1189-4b9a-8c10-13244fb21b7f-credit-group-creators",
              "type": "creditGroup"
            }
          ]
        },
        "edit": {
          "data": {
            "id": "6cc15a42-130f-4531-807a-b2c147d8ac68",
            "type": "edit"
          }
        },
        "images": {
          "data": [
            {
              "id": "i:d0",
              "type": "image"
            },
            {
              "id": "i:d1",
              "type": "image"
            },
            {
              "id": "i:d2",
              "type": "image"
            }
          ]
        },
        "primaryChannel": {
          "data": {
            "id": "c0d1f27a-e2f8-4b3c-bf3c-ed0c4e258093",
            "type": "channel"
          }
        },
        "ratingDescriptors": {
          "data": [
            {
              "id": "02a91282-d517-4f2a-a620-259ff2ba3408",
              "type": "contentDescriptor"
            },
            {
              "id": "582e60b9-b1f7-4398-9182-a8758856209d",
              "type": "contentDescriptor"
            },
            {
              "id": "b2c2a6f6-a992-4824-b43e-d0313b43c9ae",
              "type": "contentDescriptor"
            },
            {
              "id": "eb7063ee-e517-45d6-a816-16212f916850",
              "type": "contentDescriptor"
            }
          ]
        },
        "ratings": {
          "data": [
            {
              "id": "fbdd00c3-6c8f-4e27-be17-ac5997ebdfd2",
              "type": "contentRating"
            }
          ]
        },
        "season": {
          "data": {
            "id": "12bb061d-6469-44be-8903-fff07143f63a",
            "type": "season"
          }
        },
        "show": {
          "data": {
            "id": "818c3d9d-1831-48a6-9583-0364a7f98453",
            "type": "show"
          }
        },
        "txGenres": {
          "data": [
            {
              "id": "1ef95565-d5d5-4793-9656-31207932bfe1",
              "type": "taxonomyNode"
            },
            {
              "id": "3e183de4-3108-4fb5-801f-66b9ca405f94",
              "type": "taxonomyNode"
            },
            {
              "id": "435e27a6-3026-42fb-8549-b560f7616360",
              "type": "taxonomyNode"
            },
            {
              "id": "8557a482-557e-4bec-bf53-814094c40392",
              "type": "taxonomyNode"
            },
            {
              "id": "cdc65854-4013-4c85-b2b5-c48290ad3b5c",
              "type": "taxonomyNode"
            }
          ]
        }
      },
      "type": "video"
    }
  ],
  "meta": {
    "site": {
      "attributes": {
        "brandId": "beam",
        "mainTerritoryCode": "us",
        "theme": "beam",
        "websiteHostName": ""
      },
      "id": "beam_us",
      "type": "site"
    }
  }
}
~~~