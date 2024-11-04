package grammar

//go:generate wget -nc https://www.antlr.org/download/antlr-4.13.2-complete.jar
//go:generate wget https://raw.githubusercontent.com/substrait-io/substrait/820085fc913692147d0c8fdfcbf289fb8b348835/grammar/SubstraitLexer.g4
//go:generate wget https://raw.githubusercontent.com/substrait-io/substrait/820085fc913692147d0c8fdfcbf289fb8b348835/grammar/SubstraitType.g4
//go:generate -command antlr java -Xmx500M -cp "./antlr-4.13.2-complete.jar:$CLASSPATH" org.antlr.v4.Tool
//go:generate antlr -Dlanguage=Go -visitor -Dlanguage=Go -package baseparser -o "../types/parser/baseparser" SubstraitLexer.g4  SubstraitType.g4
//go:generate antlr -Dlanguage=Go -visitor -Dlanguage=Go -package baseparser -o "../testcases/parser/baseparser" FuncTestCaseLexer.g4  FuncTestCaseParser.g4
