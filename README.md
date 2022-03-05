# BusinessNZ Go(lang) API Client

See [api.business.govt.nz](https://api.business.govt.nz/) for API details.

## Status


[![GitHub tag](https://img.shields.io/github/tag/ryankurte/go-businessnz.svg)](https://github.com/ryankurte/go-businessnz)
[![Documentation](https://img.shields.io/badge/docs-godoc-blue.svg)](https://godoc.org/github.com/ryankurte/go-businessnz/lib)


### Supported APIs

- [ ] [NZBN](https://api.business.govt.nz/api/apis/info?name=NZBN&version=v4&provider=mbiecreator)
  - [x] [`/entities`](https://api.business.govt.nz/api/apis/info?name=NZBN&version=v4&provider=mbiecreator#!/Entities/EntitiesGet)
  - [x] [`/entities/{nzbn}`](https://api.business.govt.nz/api/apis/info?name=NZBN&version=v4&provider=mbiecreator#!/Entities/EntitiesByNzbnGet)
- [ ] [Companies Office](https://api.business.govt.nz/api/apis/info?name=Companies&version=v1.2&provider=mbiecreator)
  - [ ] ...


At the moment these are artisanal and hand crafted... it might be worth investigating swagger generation for the rest of them.
