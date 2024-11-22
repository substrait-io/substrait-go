package grammar

//go:generate wget -nc https://www.antlr.org/download/antlr-4.13.2-complete.jar
//go:generate wget https://raw.githubusercontent.com/substrait-io/substrait/3d2ff77575a7177f82a4d5b53408a059e9818922/grammar/SubstraitLexer.g4
//go:generate wget https://raw.githubusercontent.com/substrait-io/substrait/3d2ff77575a7177f82a4d5b53408a059e9818922/grammar/SubstraitType.g4
//go:generate wget https://raw.githubusercontent.com/substrait-io/substrait/3d2ff77575a7177f82a4d5b53408a059e9818922/grammar/FuncTestCaseLexer.g4
//go:generate wget https://raw.githubusercontent.com/substrait-io/substrait/3d2ff77575a7177f82a4d5b53408a059e9818922/grammar/FuncTestCaseParser.g4
//go:generate -command antlr java -Xmx500M -cp "./antlr-4.13.2-complete.jar:$CLASSPATH" org.antlr.v4.Tool
//go:generate antlr -Dlanguage=Go -visitor -Dlanguage=Go -package baseparser -o "../types/parser/baseparser" SubstraitLexer.g4  SubstraitType.g4
//go:generate antlr -Dlanguage=Go -visitor -no-listener -Dlanguage=Go -package baseparser -o "../testcases/parser/baseparser" FuncTestCaseLexer.g4  FuncTestCaseParser.g4
