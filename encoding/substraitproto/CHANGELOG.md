# Changelog

## [1.0.0](https://github.com/substrait-io/substrait-go/compare/encoding/substraitproto/v0.1.0...encoding/substraitproto/v1.0.0) (2026-11-20)


### ⚠ BREAKING CHANGES

* **plan:** SetProjection now returns an error. Callers should handle rejected updates and update matching interface or method-value signatures.

### Bug Fixes

* **encoding:** sort type variation and function declarations by anchor ([#385](https://github.com/substrait-io/substrait-go/issues/385)) ([16039c0](https://github.com/substrait-io/substrait-go/commit/16039c0123b9ea26b081d94185f8bb233f9d1100))
* **extensions:** preserve imported anchor assignments ([#323](https://github.com/substrait-io/substrait-go/issues/323)) ([e63a98f](https://github.com/substrait-io/substrait-go/commit/e63a98f76f0ce0f12e1e063f27682946cf97af6f))
* **plan:** apply read projections to output schemas ([#325](https://github.com/substrait-io/substrait-go/issues/325)) ([7f08714](https://github.com/substrait-io/substrait-go/commit/7f087148a813d196a6c30169f95762e562e8af84))
* **plan:** correct aggregate grouping schemas ([#324](https://github.com/substrait-io/substrait-go/issues/324)) ([e186fee](https://github.com/substrait-io/substrait-go/commit/e186feed5e86998e197de43b235e0067a40ff73d))

## 0.1.0 (2026-11-03)


### Bug Fixes

* **encoding:** always write both compound interval components ([#376](https://github.com/substrait-io/substrait-go/issues/376)) ([4ffbe38](https://github.com/substrait-io/substrait-go/commit/4ffbe38d34b91e758925913f3e75160a0f96ae5d))
