# BTPN SYARIAH - RAKAMIN Final Assignments

## Generate Keys
Using https://travistidwell.com/jsencrypt/demo/ and encoded with base64

## What to be improved 
- Clean architecture
  - https://github.com/bxcodec/go-clean-arch
  - https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html
  - https://manakuro.medium.com/clean-architecture-with-go-bce409427d31
  - https://eltonminetto.dev/en/post/2020-07-06-clean-architecture-2years-later/
- Using UUID instead of id
  - if you don't have extremely strict memory and performance requirements, I'd recommend using UUIDs instead of integer IDs. It reduces the chances of future ID conflicts and makes building a distributed architecture much easier.
- Table driven testing
  - https://dev.to/boncheff/table-driven-unit-tests-in-go-407b
  - https://yourbasic.org/golang/table-driven-unit-test/