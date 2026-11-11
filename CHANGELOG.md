# Changelog

## [10.0.0](https://github.com/substrait-io/substrait-go/compare/v9.0.1...v10.0.0) (2026-11-11)


### ⚠ BREAKING CHANGES

* **codec:** proto conversion now lives in the codec/v9 module, not core.
* Callers must use wire conversion functions instead of the domain packages' protobuf conversion APIs.
* **plan:** plan.ParquetReadOptions, plan.ArrowReadOptions, plan.OrcReadOptions, and plan.DwrfReadOptions are now substrait-go structs rather than aliases of the protobuf messages.
* **plan:** FetchRel Count() now returns expr.Expression
* **expr:** types.UserDefinedLiteral is removed; it is now expr.UserDefinedLiteral, with its value modeled as expr.UserDefinedLiteralValue (UserDefinedValueAny / UserDefinedValueStruct).
* **plan:** plan.Hint, plan.Stats, and plan.RuntimeConstraint are now substrait-go structs rather than aliases of the protobuf messages.
* **extensions:** extensions.AdvancedExtension is now a substrait-go struct rather than an alias of the protobuf message.
* **plan:** `SetComparisonReductionOp` and `SetComparisonComparisonOp` are now substrait-go types rather than aliases of the protobuf enums.
* **type:** `AggregationPhase` and `AggregationInvocation` are now substrait-go types rather than aliases of the protobuf enums.
* **expr:** `expr.MaskListElement` and `expr.MaskListSlice` are no longer defined over the protobuf messages; they are plain structs. Code relying on the old proto-typed layout must use the struct fields.
* **type:** `types.FunctionOption` is now a substrait-go struct, not an alias of `proto.FunctionOption`.
* **type:** make PrecisionTime, PrecisionTimestamp, PrecisionTimestampTz domain structs ([#322](https://github.com/substrait-io/substrait-go/issues/322))
* **plan:** `plan.SetPredicateOp` is no longer an alias for the protobuf enum; its constants are now typed values. Code that passes one where the protobuf enum is expected needs a cast.
* **type:** `types.IntervalDayToSecond` is no longer an alias for `proto.Expression_Literal_IntervalDayToSecond`; it is a plain struct whose sub-second precision is carried by the `PrecisionMode` oneof interface. Code that used the protobuf message directly must construct the struct.
* **plan:** `plan.WriteOp` and `plan.OutputMode` are no longer aliases for the protobuf enums; their constants are now typed values. Code that passes one where the other is expected needs a cast.
* **type:** `types.IntervalYearToMonth` is no longer an alias for `proto.Expression_Literal_IntervalYearToMonth`; it is now a plain struct. Code that passed the protobuf message where the alias was expected needs to construct the struct instead.
* **plan:** `plan.SetOp` is no longer an alias for `proto.SetRel_SetOp`; its constants are now typed `SetOp` values. Code that passes one where the other is expected needs a cast.
* **type:** `types.BoundsType` is no longer an alias for `proto.Expression_WindowFunction_BoundsType`; its constants are now typed `BoundsType` values. Code that passes one where the other is expected needs a cast.
* **plan:** `plan.JoinType` is no longer an alias for `proto.JoinRel_JoinType`; its constants are now typed `JoinType` values. Code that passes one where the other is expected needs a cast.
* **type:** `types.SortDirection`'s underlying type changes from `proto.SortField_SortDirection` to `int32`. Code that passes one where the other is expected needs a cast.
* **type:** `types.Decimal` is no longer an alias for `proto.Expression_Literal_Decimal`.
* **type:** `types.VarChar` is no longer an alias for `proto.Expression_Literal_VarChar`.
* **plan:** `plan.SimpleComparisonType` is no longer an alias for `proto.ComparisonJoinKey_SimpleComparisonType`. Code that passes one where the other is expected needs a cast.
* **type:** make CastFailBehavior a domain enum ([#305](https://github.com/substrait-io/substrait-go/issues/305))
* **type:** types.Version is no longer an alias for substraitpb.Version
* import path changes to substrait-go/v9
* *Remap functions have been removed from Builder interface
* **type:** `types.Nullability` is no longer an alias for `proto.Type_Nullability`. Code that passes one where the other is expected needs a cast.
* **literal:** `NewList` now rejects elements with differing Substrait types, including nullability, parameters, variations, and nested types.
* **extensions:** use distinct extension ID types ([#258](https://github.com/substrait-io/substrait-go/issues/258))
* **testcases:** model non-value results explicitly ([#234](https://github.com/substrait-io/substrait-go/issues/234))
* **extensions:** drop loader URI tracking ([#250](https://github.com/substrait-io/substrait-go/issues/250))

### Features

* **codec:** add a new codec module and wire it into ci ([#283](https://github.com/substrait-io/substrait-go/issues/283)) ([43bae29](https://github.com/substrait-io/substrait-go/commit/43bae29153db57f322c2c833489fe5372386b15b))
* **codec:** release the codec module and drop proto from core ([#365](https://github.com/substrait-io/substrait-go/issues/365)) ([c1c7302](https://github.com/substrait-io/substrait-go/commit/c1c73027eea8c54319c3c1e640a8f7faf0cf5201))
* **expr:** make MaskListElement and MaskListSlice domain structs ([#319](https://github.com/substrait-io/substrait-go/issues/319)) ([390c806](https://github.com/substrait-io/substrait-go/commit/390c8061925c428b07830e3bdb0d775f8959c840))
* **expr:** make UserDefinedLiteral a domain type ([#339](https://github.com/substrait-io/substrait-go/issues/339)) ([c48676c](https://github.com/substrait-io/substrait-go/commit/c48676c3b4695e5423f854e541968ee13f4b8df9))
* **extensions:** drop loader URI tracking ([#250](https://github.com/substrait-io/substrait-go/issues/250)) ([16e339a](https://github.com/substrait-io/substrait-go/commit/16e339a2ed34b4aeaf6d4bcd03f72b501dfcbfc9))
* **extensions:** expose function variant return types ([#306](https://github.com/substrait-io/substrait-go/issues/306)) ([278db13](https://github.com/substrait-io/substrait-go/commit/278db13915b198a833042e649ee5f016549cbb85))
* **extensions:** make AdvancedExtension a domain type ([#340](https://github.com/substrait-io/substrait-go/issues/340)) ([df007ba](https://github.com/substrait-io/substrait-go/commit/df007ba63b7d66395c4d4a35c371e047d2f9c47b))
* **extensions:** validate local UDT references ([#252](https://github.com/substrait-io/substrait-go/issues/252)) ([5fbe06d](https://github.com/substrait-io/substrait-go/commit/5fbe06de0170fb1652cdbb5159bfb7f89bf113e1))
* **extensions:** validate simple extension dependencies ([#247](https://github.com/substrait-io/substrait-go/issues/247)) ([45c593e](https://github.com/substrait-io/substrait-go/commit/45c593ef3b6f796a48e9ff678f3b984f09b17735))
* **literal:** add NewMap, NewEmptyMap, NewEmptyList constructors ([#260](https://github.com/substrait-io/substrait-go/issues/260)) ([642f207](https://github.com/substrait-io/substrait-go/commit/642f207fd4a3e3e10575c04428f26802eb8e54f3))
* **plan:** make Hint, Stats, and RuntimeConstraint domain types ([#341](https://github.com/substrait-io/substrait-go/issues/341)) ([5fac87f](https://github.com/substrait-io/substrait-go/commit/5fac87feb7725fb5c6e436970b65509f4d741a35))
* **plan:** make JoinType a domain enum ([#312](https://github.com/substrait-io/substrait-go/issues/312)) ([48809b4](https://github.com/substrait-io/substrait-go/commit/48809b4325e0091449487bd52f834c1bbec119c4))
* **plan:** make SetComparisonReductionOp and SetComparisonComparisonOp domain enums ([#338](https://github.com/substrait-io/substrait-go/issues/338)) ([a09cc3c](https://github.com/substrait-io/substrait-go/commit/a09cc3c10be8a9136bc408ba6b8e58399d4aa9fa))
* **plan:** make SetOp a domain enum ([#315](https://github.com/substrait-io/substrait-go/issues/315)) ([11f865a](https://github.com/substrait-io/substrait-go/commit/11f865ac48f710d9b5afb944341e2f73e024bc6b))
* **plan:** make SetPredicateOp a domain enum ([#335](https://github.com/substrait-io/substrait-go/issues/335)) ([f7e81a1](https://github.com/substrait-io/substrait-go/commit/f7e81a150c37986831731eb5c0bfc299ea223616))
* **plan:** make SimpleComparisonType a domain enum ([#307](https://github.com/substrait-io/substrait-go/issues/307)) ([27995e3](https://github.com/substrait-io/substrait-go/commit/27995e325bae641d65e6aebdd37424f9d9e06d18))
* **plan:** make WriteOp and OutputMode domain enums ([#318](https://github.com/substrait-io/substrait-go/issues/318)) ([cb2d6e6](https://github.com/substrait-io/substrait-go/commit/cb2d6e648bc0ee1549e354dbfe705d9f8a1ce2b1))
* **plan:** support expression encoding for FetchRel offset and count ([#295](https://github.com/substrait-io/substrait-go/issues/295)) ([fd888bd](https://github.com/substrait-io/substrait-go/commit/fd888bd5e2f8207387d9c049051b57c8fdb14066))
* support new ComparisonJoinKey field in HashJoinRel/MergeJoinRel ([#256](https://github.com/substrait-io/substrait-go/issues/256)) ([92b4f71](https://github.com/substrait-io/substrait-go/commit/92b4f717d0c0d5b8d1668958e86de31fbb15af9d))
* **testcases:** model non-value results explicitly ([#234](https://github.com/substrait-io/substrait-go/issues/234)) ([16abb54](https://github.com/substrait-io/substrait-go/commit/16abb54bf79ba49412caf592601abe5266272890))
* **type:** define Nullability instead of aliasing the protobuf enum ([#282](https://github.com/substrait-io/substrait-go/issues/282)) ([b922a54](https://github.com/substrait-io/substrait-go/commit/b922a54264fa6a54ca8c5d50022ff30ea4b06858))
* **type:** make AggregationPhase and AggregationInvocation domain enums ([#337](https://github.com/substrait-io/substrait-go/issues/337)) ([a903eda](https://github.com/substrait-io/substrait-go/commit/a903edab54285fd55bc863bb4f5a86d5b231f5f5))
* **type:** make BoundsType a domain enum ([#313](https://github.com/substrait-io/substrait-go/issues/313)) ([a499fd8](https://github.com/substrait-io/substrait-go/commit/a499fd88527209c5d048e344e831f4d9bbceffb0))
* **type:** make CastFailBehavior a domain enum ([#305](https://github.com/substrait-io/substrait-go/issues/305)) ([77182a2](https://github.com/substrait-io/substrait-go/commit/77182a2830f052c22996d53b81b029fa86608449))
* **type:** make Decimal a domain struct ([#309](https://github.com/substrait-io/substrait-go/issues/309)) ([4d06c2b](https://github.com/substrait-io/substrait-go/commit/4d06c2b963347d74ee1c9fc1c72147dfa2037862))
* **type:** make FunctionOption a domain struct ([#321](https://github.com/substrait-io/substrait-go/issues/321)) ([3d056d3](https://github.com/substrait-io/substrait-go/commit/3d056d3133797437170ec9b439e326f0600125c5))
* **type:** make IntervalDayToSecond a domain type ([#317](https://github.com/substrait-io/substrait-go/issues/317)) ([f1a5f20](https://github.com/substrait-io/substrait-go/commit/f1a5f204478605652309bd64533d98ef0a7d10f1))
* **type:** make IntervalYearToMonth a domain struct ([#314](https://github.com/substrait-io/substrait-go/issues/314)) ([51aba74](https://github.com/substrait-io/substrait-go/commit/51aba7474498964693f9eeac98b8b40db5fd4c81))
* **type:** make PrecisionTime, PrecisionTimestamp, PrecisionTimestampTz domain structs ([#322](https://github.com/substrait-io/substrait-go/issues/322)) ([402d44e](https://github.com/substrait-io/substrait-go/commit/402d44ea819565be788a2ec7425d269c6370a356))
* **type:** make SortDirection a domain enum ([#311](https://github.com/substrait-io/substrait-go/issues/311)) ([ef09d37](https://github.com/substrait-io/substrait-go/commit/ef09d378c5679d2cd84b3002dd2c3bfcda99e799))
* **type:** make types.Version a Go struct ([#300](https://github.com/substrait-io/substrait-go/issues/300)) ([45b9246](https://github.com/substrait-io/substrait-go/commit/45b92468381136213cafd407c36a93dc7a5065c4))
* **type:** make VarChar a domain struct ([#308](https://github.com/substrait-io/substrait-go/issues/308)) ([757fb9b](https://github.com/substrait-io/substrait-go/commit/757fb9b4bcc1b46969c21cd3031224186120b690))


### Bug Fixes

* avoid StructFieldRef.GetType panics on invalid field indices ([#296](https://github.com/substrait-io/substrait-go/issues/296)) ([47adfb6](https://github.com/substrait-io/substrait-go/commit/47adfb6b5a744400cb572363ceb9beb220a0eae9))
* decode schemas in extension relations ([#276](https://github.com/substrait-io/substrait-go/issues/276)) ([985d911](https://github.com/substrait-io/substrait-go/commit/985d911da1f72c4e66ea0cd2e1470105f897bc83))
* **expr:** accept pre-epoch precision timestamps ([#346](https://github.com/substrait-io/substrait-go/issues/346)) ([c2b82be](https://github.com/substrait-io/substrait-go/commit/c2b82be366bd8114c21a46101a668542f5899d53))
* **expr:** initialize MapLiteral KeyValue entries before use ([#235](https://github.com/substrait-io/substrait-go/issues/235)) ([c9615c8](https://github.com/substrait-io/substrait-go/commit/c9615c8cc24caf069bf492199f8cb8a5d0cce2c7))
* **extensions:** return errors for malformed function argument YAML ([#327](https://github.com/substrait-io/substrait-go/issues/327)) ([652df2c](https://github.com/substrait-io/substrait-go/commit/652df2c79a600405c97863d022c06012d16e8375)), closes [#287](https://github.com/substrait-io/substrait-go/issues/287)
* persist extension registry in SubqueryConverters for FromProto ([#273](https://github.com/substrait-io/substrait-go/issues/273)) ([4a512e6](https://github.com/substrait-io/substrait-go/commit/4a512e6f8750ec91fe5d8e8afd4d67cf9f58a8f2))
* **plan:** correct physical join schemas and post-filter binding ([8a6e2b1](https://github.com/substrait-io/substrait-go/commit/8a6e2b104d62385064177aa2a0da1f7b4de6154e)), closes [#330](https://github.com/substrait-io/substrait-go/issues/330)
* **plan:** derive right join record types ([#271](https://github.com/substrait-io/substrait-go/issues/271)) ([c5dd519](https://github.com/substrait-io/substrait-go/commit/c5dd519663a3b68d8b27758f77ded6aa37a4800c))
* **plan:** handle absent RelCommon without panicking ([#344](https://github.com/substrait-io/substrait-go/issues/344)) ([ea72bcd](https://github.com/substrait-io/substrait-go/commit/ea72bcdac09e50ce71e10546f329bee545aedaa8))
* return error when function outputType is nil in ExprFromProto ([#259](https://github.com/substrait-io/substrait-go/issues/259)) ([4a63222](https://github.com/substrait-io/substrait-go/commit/4a63222436f6ddd36071a92636c733f841404a14))
* **type:** keep the type variation reference in WithNullability ([#290](https://github.com/substrait-io/substrait-go/issues/290)) ([334a587](https://github.com/substrait-io/substrait-go/commit/334a587dfcb3f06fd7361df81065d30c27d8fc5a))
* **types:** avoid mutating interval day nullability ([#349](https://github.com/substrait-io/substrait-go/issues/349)) ([f8e5ffe](https://github.com/substrait-io/substrait-go/commit/f8e5ffec6826b5327c7b5cee5efd73f8b08eadb8))


### Code Refactoring

* **extensions:** use distinct extension ID types ([#258](https://github.com/substrait-io/substrait-go/issues/258)) ([ab8b65e](https://github.com/substrait-io/substrait-go/commit/ab8b65e9eb6d581e0bd3cab642385249289c5ab6))
* move Substrait protobuf conversions into wire ([#352](https://github.com/substrait-io/substrait-go/issues/352)) ([81fd963](https://github.com/substrait-io/substrait-go/commit/81fd9630530358e2274ee46aff7a8d468549999e))
* **plan:** make FileOrFiles read-options domain types ([#350](https://github.com/substrait-io/substrait-go/issues/350)) ([94416fd](https://github.com/substrait-io/substrait-go/commit/94416fd07edcfb903bfb3c8affc5b8d2f215540c))
* remove deprecated *Remap builder functions ([#294](https://github.com/substrait-io/substrait-go/issues/294)) ([0f3a086](https://github.com/substrait-io/substrait-go/commit/0f3a0868a6cd76c92f4b836c05385af5d3fef26c))


### Build System

* bump go.mod to v9 ([#301](https://github.com/substrait-io/substrait-go/issues/301)) ([6e5b684](https://github.com/substrait-io/substrait-go/commit/6e5b68434a5b6f55da44462228777d98bdda6c08))

## [9.0.1](https://github.com/substrait-io/substrait-go/compare/v9.0.0...v9.0.1) (2026-11-03)


### Bug Fixes

* **extensions:** return errors for malformed function argument YAML ([#327](https://github.com/substrait-io/substrait-go/issues/327)) ([652df2c](https://github.com/substrait-io/substrait-go/commit/652df2c79a600405c97863d022c06012d16e8375)), closes [#287](https://github.com/substrait-io/substrait-go/issues/287)

## [9.0.0](https://github.com/substrait-io/substrait-go/compare/v8.1.1...v9.0.0) (2026-11-03)


### ⚠ BREAKING CHANGES

* **codec:** proto conversion now lives in the codec/v9 module, not core.
* Callers must use wire conversion functions instead of the domain packages' protobuf conversion APIs.
* **plan:** plan.ParquetReadOptions, plan.ArrowReadOptions, plan.OrcReadOptions, and plan.DwrfReadOptions are now substrait-go structs rather than aliases of the protobuf messages.
* **plan:** FetchRel Count() now returns expr.Expression
* **expr:** types.UserDefinedLiteral is removed; it is now expr.UserDefinedLiteral, with its value modeled as expr.UserDefinedLiteralValue (UserDefinedValueAny / UserDefinedValueStruct).
* **plan:** plan.Hint, plan.Stats, and plan.RuntimeConstraint are now substrait-go structs rather than aliases of the protobuf messages.
* **extensions:** extensions.AdvancedExtension is now a substrait-go struct rather than an alias of the protobuf message.
* **plan:** `SetComparisonReductionOp` and `SetComparisonComparisonOp` are now substrait-go types rather than aliases of the protobuf enums.
* **type:** `AggregationPhase` and `AggregationInvocation` are now substrait-go types rather than aliases of the protobuf enums.
* **expr:** `expr.MaskListElement` and `expr.MaskListSlice` are no longer defined over the protobuf messages; they are plain structs. Code relying on the old proto-typed layout must use the struct fields.
* **type:** `types.FunctionOption` is now a substrait-go struct, not an alias of `proto.FunctionOption`.
* **type:** make PrecisionTime, PrecisionTimestamp, PrecisionTimestampTz domain structs ([#322](https://github.com/substrait-io/substrait-go/issues/322))
* **plan:** `plan.SetPredicateOp` is no longer an alias for the protobuf enum; its constants are now typed values. Code that passes one where the protobuf enum is expected needs a cast.
* **type:** `types.IntervalDayToSecond` is no longer an alias for `proto.Expression_Literal_IntervalDayToSecond`; it is a plain struct whose sub-second precision is carried by the `PrecisionMode` oneof interface. Code that used the protobuf message directly must construct the struct.
* **plan:** `plan.WriteOp` and `plan.OutputMode` are no longer aliases for the protobuf enums; their constants are now typed values. Code that passes one where the other is expected needs a cast.
* **type:** `types.IntervalYearToMonth` is no longer an alias for `proto.Expression_Literal_IntervalYearToMonth`; it is now a plain struct. Code that passed the protobuf message where the alias was expected needs to construct the struct instead.
* **plan:** `plan.SetOp` is no longer an alias for `proto.SetRel_SetOp`; its constants are now typed `SetOp` values. Code that passes one where the other is expected needs a cast.
* **type:** `types.BoundsType` is no longer an alias for `proto.Expression_WindowFunction_BoundsType`; its constants are now typed `BoundsType` values. Code that passes one where the other is expected needs a cast.
* **plan:** `plan.JoinType` is no longer an alias for `proto.JoinRel_JoinType`; its constants are now typed `JoinType` values. Code that passes one where the other is expected needs a cast.
* **type:** `types.SortDirection`'s underlying type changes from `proto.SortField_SortDirection` to `int32`. Code that passes one where the other is expected needs a cast.
* **type:** `types.Decimal` is no longer an alias for `proto.Expression_Literal_Decimal`.
* **type:** `types.VarChar` is no longer an alias for `proto.Expression_Literal_VarChar`.
* **plan:** `plan.SimpleComparisonType` is no longer an alias for `proto.ComparisonJoinKey_SimpleComparisonType`. Code that passes one where the other is expected needs a cast.
* **type:** make CastFailBehavior a domain enum ([#305](https://github.com/substrait-io/substrait-go/issues/305))
* **type:** types.Version is no longer an alias for substraitpb.Version
* import path changes to substrait-go/v9
* *Remap functions have been removed from Builder interface
* **type:** `types.Nullability` is no longer an alias for `proto.Type_Nullability`. Code that passes one where the other is expected needs a cast.
* **literal:** `NewList` now rejects elements with differing Substrait types, including nullability, parameters, variations, and nested types.
* **extensions:** use distinct extension ID types ([#258](https://github.com/substrait-io/substrait-go/issues/258))
* **testcases:** model non-value results explicitly ([#234](https://github.com/substrait-io/substrait-go/issues/234))
* **extensions:** drop loader URI tracking ([#250](https://github.com/substrait-io/substrait-go/issues/250))

### Features

* **codec:** add a new codec module and wire it into ci ([#283](https://github.com/substrait-io/substrait-go/issues/283)) ([43bae29](https://github.com/substrait-io/substrait-go/commit/43bae29153db57f322c2c833489fe5372386b15b))
* **codec:** release the codec module and drop proto from core ([#365](https://github.com/substrait-io/substrait-go/issues/365)) ([c1c7302](https://github.com/substrait-io/substrait-go/commit/c1c73027eea8c54319c3c1e640a8f7faf0cf5201))
* **expr:** make MaskListElement and MaskListSlice domain structs ([#319](https://github.com/substrait-io/substrait-go/issues/319)) ([390c806](https://github.com/substrait-io/substrait-go/commit/390c8061925c428b07830e3bdb0d775f8959c840))
* **expr:** make UserDefinedLiteral a domain type ([#339](https://github.com/substrait-io/substrait-go/issues/339)) ([c48676c](https://github.com/substrait-io/substrait-go/commit/c48676c3b4695e5423f854e541968ee13f4b8df9))
* **extensions:** drop loader URI tracking ([#250](https://github.com/substrait-io/substrait-go/issues/250)) ([16e339a](https://github.com/substrait-io/substrait-go/commit/16e339a2ed34b4aeaf6d4bcd03f72b501dfcbfc9))
* **extensions:** expose function variant return types ([#306](https://github.com/substrait-io/substrait-go/issues/306)) ([278db13](https://github.com/substrait-io/substrait-go/commit/278db13915b198a833042e649ee5f016549cbb85))
* **extensions:** make AdvancedExtension a domain type ([#340](https://github.com/substrait-io/substrait-go/issues/340)) ([df007ba](https://github.com/substrait-io/substrait-go/commit/df007ba63b7d66395c4d4a35c371e047d2f9c47b))
* **extensions:** validate local UDT references ([#252](https://github.com/substrait-io/substrait-go/issues/252)) ([5fbe06d](https://github.com/substrait-io/substrait-go/commit/5fbe06de0170fb1652cdbb5159bfb7f89bf113e1))
* **extensions:** validate simple extension dependencies ([#247](https://github.com/substrait-io/substrait-go/issues/247)) ([45c593e](https://github.com/substrait-io/substrait-go/commit/45c593ef3b6f796a48e9ff678f3b984f09b17735))
* **literal:** add NewMap, NewEmptyMap, NewEmptyList constructors ([#260](https://github.com/substrait-io/substrait-go/issues/260)) ([642f207](https://github.com/substrait-io/substrait-go/commit/642f207fd4a3e3e10575c04428f26802eb8e54f3))
* **plan:** make Hint, Stats, and RuntimeConstraint domain types ([#341](https://github.com/substrait-io/substrait-go/issues/341)) ([5fac87f](https://github.com/substrait-io/substrait-go/commit/5fac87feb7725fb5c6e436970b65509f4d741a35))
* **plan:** make JoinType a domain enum ([#312](https://github.com/substrait-io/substrait-go/issues/312)) ([48809b4](https://github.com/substrait-io/substrait-go/commit/48809b4325e0091449487bd52f834c1bbec119c4))
* **plan:** make SetComparisonReductionOp and SetComparisonComparisonOp domain enums ([#338](https://github.com/substrait-io/substrait-go/issues/338)) ([a09cc3c](https://github.com/substrait-io/substrait-go/commit/a09cc3c10be8a9136bc408ba6b8e58399d4aa9fa))
* **plan:** make SetOp a domain enum ([#315](https://github.com/substrait-io/substrait-go/issues/315)) ([11f865a](https://github.com/substrait-io/substrait-go/commit/11f865ac48f710d9b5afb944341e2f73e024bc6b))
* **plan:** make SetPredicateOp a domain enum ([#335](https://github.com/substrait-io/substrait-go/issues/335)) ([f7e81a1](https://github.com/substrait-io/substrait-go/commit/f7e81a150c37986831731eb5c0bfc299ea223616))
* **plan:** make SimpleComparisonType a domain enum ([#307](https://github.com/substrait-io/substrait-go/issues/307)) ([27995e3](https://github.com/substrait-io/substrait-go/commit/27995e325bae641d65e6aebdd37424f9d9e06d18))
* **plan:** make WriteOp and OutputMode domain enums ([#318](https://github.com/substrait-io/substrait-go/issues/318)) ([cb2d6e6](https://github.com/substrait-io/substrait-go/commit/cb2d6e648bc0ee1549e354dbfe705d9f8a1ce2b1))
* **plan:** support expression encoding for FetchRel offset and count ([#295](https://github.com/substrait-io/substrait-go/issues/295)) ([fd888bd](https://github.com/substrait-io/substrait-go/commit/fd888bd5e2f8207387d9c049051b57c8fdb14066))
* support new ComparisonJoinKey field in HashJoinRel/MergeJoinRel ([#256](https://github.com/substrait-io/substrait-go/issues/256)) ([92b4f71](https://github.com/substrait-io/substrait-go/commit/92b4f717d0c0d5b8d1668958e86de31fbb15af9d))
* **testcases:** model non-value results explicitly ([#234](https://github.com/substrait-io/substrait-go/issues/234)) ([16abb54](https://github.com/substrait-io/substrait-go/commit/16abb54bf79ba49412caf592601abe5266272890))
* **type:** define Nullability instead of aliasing the protobuf enum ([#282](https://github.com/substrait-io/substrait-go/issues/282)) ([b922a54](https://github.com/substrait-io/substrait-go/commit/b922a54264fa6a54ca8c5d50022ff30ea4b06858))
* **type:** make AggregationPhase and AggregationInvocation domain enums ([#337](https://github.com/substrait-io/substrait-go/issues/337)) ([a903eda](https://github.com/substrait-io/substrait-go/commit/a903edab54285fd55bc863bb4f5a86d5b231f5f5))
* **type:** make BoundsType a domain enum ([#313](https://github.com/substrait-io/substrait-go/issues/313)) ([a499fd8](https://github.com/substrait-io/substrait-go/commit/a499fd88527209c5d048e344e831f4d9bbceffb0))
* **type:** make CastFailBehavior a domain enum ([#305](https://github.com/substrait-io/substrait-go/issues/305)) ([77182a2](https://github.com/substrait-io/substrait-go/commit/77182a2830f052c22996d53b81b029fa86608449))
* **type:** make Decimal a domain struct ([#309](https://github.com/substrait-io/substrait-go/issues/309)) ([4d06c2b](https://github.com/substrait-io/substrait-go/commit/4d06c2b963347d74ee1c9fc1c72147dfa2037862))
* **type:** make FunctionOption a domain struct ([#321](https://github.com/substrait-io/substrait-go/issues/321)) ([3d056d3](https://github.com/substrait-io/substrait-go/commit/3d056d3133797437170ec9b439e326f0600125c5))
* **type:** make IntervalDayToSecond a domain type ([#317](https://github.com/substrait-io/substrait-go/issues/317)) ([f1a5f20](https://github.com/substrait-io/substrait-go/commit/f1a5f204478605652309bd64533d98ef0a7d10f1))
* **type:** make IntervalYearToMonth a domain struct ([#314](https://github.com/substrait-io/substrait-go/issues/314)) ([51aba74](https://github.com/substrait-io/substrait-go/commit/51aba7474498964693f9eeac98b8b40db5fd4c81))
* **type:** make PrecisionTime, PrecisionTimestamp, PrecisionTimestampTz domain structs ([#322](https://github.com/substrait-io/substrait-go/issues/322)) ([402d44e](https://github.com/substrait-io/substrait-go/commit/402d44ea819565be788a2ec7425d269c6370a356))
* **type:** make SortDirection a domain enum ([#311](https://github.com/substrait-io/substrait-go/issues/311)) ([ef09d37](https://github.com/substrait-io/substrait-go/commit/ef09d378c5679d2cd84b3002dd2c3bfcda99e799))
* **type:** make types.Version a Go struct ([#300](https://github.com/substrait-io/substrait-go/issues/300)) ([45b9246](https://github.com/substrait-io/substrait-go/commit/45b92468381136213cafd407c36a93dc7a5065c4))
* **type:** make VarChar a domain struct ([#308](https://github.com/substrait-io/substrait-go/issues/308)) ([757fb9b](https://github.com/substrait-io/substrait-go/commit/757fb9b4bcc1b46969c21cd3031224186120b690))


### Bug Fixes

* avoid StructFieldRef.GetType panics on invalid field indices ([#296](https://github.com/substrait-io/substrait-go/issues/296)) ([47adfb6](https://github.com/substrait-io/substrait-go/commit/47adfb6b5a744400cb572363ceb9beb220a0eae9))
* decode schemas in extension relations ([#276](https://github.com/substrait-io/substrait-go/issues/276)) ([985d911](https://github.com/substrait-io/substrait-go/commit/985d911da1f72c4e66ea0cd2e1470105f897bc83))
* **expr:** accept pre-epoch precision timestamps ([#346](https://github.com/substrait-io/substrait-go/issues/346)) ([c2b82be](https://github.com/substrait-io/substrait-go/commit/c2b82be366bd8114c21a46101a668542f5899d53))
* **expr:** initialize MapLiteral KeyValue entries before use ([#235](https://github.com/substrait-io/substrait-go/issues/235)) ([c9615c8](https://github.com/substrait-io/substrait-go/commit/c9615c8cc24caf069bf492199f8cb8a5d0cce2c7))
* persist extension registry in SubqueryConverters for FromProto ([#273](https://github.com/substrait-io/substrait-go/issues/273)) ([4a512e6](https://github.com/substrait-io/substrait-go/commit/4a512e6f8750ec91fe5d8e8afd4d67cf9f58a8f2))
* **plan:** correct physical join schemas and post-filter binding ([8a6e2b1](https://github.com/substrait-io/substrait-go/commit/8a6e2b104d62385064177aa2a0da1f7b4de6154e)), closes [#330](https://github.com/substrait-io/substrait-go/issues/330)
* **plan:** derive right join record types ([#271](https://github.com/substrait-io/substrait-go/issues/271)) ([c5dd519](https://github.com/substrait-io/substrait-go/commit/c5dd519663a3b68d8b27758f77ded6aa37a4800c))
* **plan:** handle absent RelCommon without panicking ([#344](https://github.com/substrait-io/substrait-go/issues/344)) ([ea72bcd](https://github.com/substrait-io/substrait-go/commit/ea72bcdac09e50ce71e10546f329bee545aedaa8))
* return error when function outputType is nil in ExprFromProto ([#259](https://github.com/substrait-io/substrait-go/issues/259)) ([4a63222](https://github.com/substrait-io/substrait-go/commit/4a63222436f6ddd36071a92636c733f841404a14))
* **type:** keep the type variation reference in WithNullability ([#290](https://github.com/substrait-io/substrait-go/issues/290)) ([334a587](https://github.com/substrait-io/substrait-go/commit/334a587dfcb3f06fd7361df81065d30c27d8fc5a))
* **types:** avoid mutating interval day nullability ([#349](https://github.com/substrait-io/substrait-go/issues/349)) ([f8e5ffe](https://github.com/substrait-io/substrait-go/commit/f8e5ffec6826b5327c7b5cee5efd73f8b08eadb8))


### Code Refactoring

* **extensions:** use distinct extension ID types ([#258](https://github.com/substrait-io/substrait-go/issues/258)) ([ab8b65e](https://github.com/substrait-io/substrait-go/commit/ab8b65e9eb6d581e0bd3cab642385249289c5ab6))
* move Substrait protobuf conversions into wire ([#352](https://github.com/substrait-io/substrait-go/issues/352)) ([81fd963](https://github.com/substrait-io/substrait-go/commit/81fd9630530358e2274ee46aff7a8d468549999e))
* **plan:** make FileOrFiles read-options domain types ([#350](https://github.com/substrait-io/substrait-go/issues/350)) ([94416fd](https://github.com/substrait-io/substrait-go/commit/94416fd07edcfb903bfb3c8affc5b8d2f215540c))
* remove deprecated *Remap builder functions ([#294](https://github.com/substrait-io/substrait-go/issues/294)) ([0f3a086](https://github.com/substrait-io/substrait-go/commit/0f3a0868a6cd76c92f4b836c05385af5d3fef26c))


### Build System

* bump go.mod to v9 ([#301](https://github.com/substrait-io/substrait-go/issues/301)) ([6e5b684](https://github.com/substrait-io/substrait-go/commit/6e5b68434a5b6f55da44462228777d98bdda6c08))
