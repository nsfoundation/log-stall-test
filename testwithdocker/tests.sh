#!/bin/bash

max=30
for (( i=0; i < max; i++ ))
do

	echo "line $i"
	printf "\`abjkdsljfkds\` has been deprecated, please migrate to \`abjkdsljfkds\`\n at abjkdsljfkds(/src/node_modules/somepath:38392:35)\n at Something.SomethingElse (/src/node_modules/someotherpath.js:123:24)\n"
	printf "\`abjkdsljfkds\` has been deprecated, please migrate to \`abjkdsljfkds\` at abjkdsljfkds(/src/node_modules/somepath:38392:35) at Something.SomethingElse (/src/node_modules/someotherpath.js:123:24)"
	printf "\`abjkdsljfkds\` has been deprecated, please migrate to \`abjkdsljfkds\`\r at abjkdsljfkds(/src/node_modules/somepath:38392:35)\r at Something.SomethingElse (/src/node_modules/someotherpath.js:123:24)\r"

	cp output /dev/stderr

done



