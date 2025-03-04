#!/bin/bash

max=200
for (( i=0; i < max; i++ )); do echo "line $i"; done

printf "\`abjkdsljfkds\` has been deprecated, please migrate to \`abjkdsljfkds\`\n at abjkdsljfkds(/src/node_modules/somepath:38392:35)\n at Something.SomethingElse (/src/node_modules/someotherpath.js:123:24)\n"
cat output


