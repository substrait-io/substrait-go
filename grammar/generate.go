package grammar

//go:generate wget -nc https://www.antlr.org/download/antlr-4.13.2-complete.jar
//go:generate wget https://raw.githubusercontent.com/substrait-io/substrait/c18c0c1347376efa4dfab503bb4db9f820df3cf3/grammar/SubstraitLexer.g4
//go:generate wget https://raw.githubusercontent.com/substrait-io/substrait/c18c0c1347376efa4dfab503bb4db9f820df3cf3/grammar/SubstraitType.g4
//go:generate wget https://raw.githubusercontent.com/substrait-io/substrait/c18c0c1347376efa4dfab503bb4db9f820df3cf3/grammar/FuncTestCaseLexer.g4
//go:generate wget https://raw.githubusercontent.com/substrait-io/substrait/c18c0c1347376efa4dfab503bb4db9f820df3cf3/grammar/FuncTestCaseParser.g4
//go:generate -command antlr java -Xmx500M -cp "./antlr-4.13.2-complete.jar:$CLASSPATH" org.antlr.v4.Tool
//go:generate antlr -Dlanguage=Go -visitor -Dlanguage=Go -package baseparser -o "../types/parser/baseparser" SubstraitLexer.g4  SubstraitType.g4
//go:generate antlr -Dlanguage=Go -visitor -no-listener -Dlanguage=Go -package baseparser -o "../testcases/parser/baseparser" FuncTestCaseLexer.g4  FuncTestCaseParser.g4
