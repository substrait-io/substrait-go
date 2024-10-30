ANTLR_JAR=antlr-4.13.2-complete.jar
GRAMMARS=SubstraitLexer.g4 SubstraitType.g4
OUTPUT_DIR=../parser/types/baseparser

generate:
	java -jar $(ANTLR_JAR) -visitor -Dlanguage=Go -package baseparser -o $(OUTPUT_DIR) $(GRAMMARS)

clean:
	rm -rf $(OUTPUT_DIR)/*.go $(OUTPUT_DIR)/*.tokens $(OUTPUT_DIR)/*.interp
