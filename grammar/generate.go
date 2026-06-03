// SPDX-License-Identifier: Apache-2.0

package grammar

// using substrait v0.87.0 with substrait-io/substrait#1088 applied to FuncTestCaseParser.g4

//go:generate wget -nc https://www.antlr.org/download/antlr-4.13.2-complete.jar
//go:generate wget https://raw.githubusercontent.com/substrait-io/substrait/v0.87.0/grammar/SubstraitLexer.g4
//go:generate wget https://raw.githubusercontent.com/substrait-io/substrait/v0.87.0/grammar/SubstraitType.g4
//go:generate wget https://raw.githubusercontent.com/substrait-io/substrait/v0.87.0/grammar/FuncTestCaseLexer.g4
//go:generate wget https://raw.githubusercontent.com/substrait-io/substrait/v0.87.0/grammar/FuncTestCaseParser.g4
//go:generate python3 -c "from pathlib import Path; p=Path('FuncTestCaseParser.g4'); s=p.read_text(); s=s.replace('    | lambdaArg\\n    | Identifier', '    | lambdaArg\\n    | funcCallArg\\n    | Identifier'); s=s.replace('lambdaArg\\n    : literalLambda DoubleColon funcType\\n    ;\\n\\nenumArg', 'lambdaArg\\n    : literalLambda DoubleColon funcType\\n    ;\\n\\nfuncCallArg\\n    : identifier OParen arguments CParen\\n    ;\\n\\nenumArg'); p.write_text(s)"
//go:generate -command antlr java -Xmx500M -cp "./antlr-4.13.2-complete.jar:$CLASSPATH" org.antlr.v4.Tool
//go:generate antlr -Dlanguage=Go -visitor -Dlanguage=Go -package baseparser -o "../types/parser/baseparser" SubstraitLexer.g4  SubstraitType.g4
//go:generate antlr -Dlanguage=Go -visitor -no-listener -Dlanguage=Go -package baseparser -o "../testcases/parser/baseparser" FuncTestCaseLexer.g4  FuncTestCaseParser.g4
