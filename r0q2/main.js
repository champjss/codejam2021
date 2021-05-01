const readline = require('readline')

function fillSymbolsForWildcard(symbols, wildcard) {
  const length = symbols.length

  const symbolBeforeWildcard = wildcard.startIndex > 0
    ? symbols[wildcard.startIndex - 1]
    : null
  const symbolAfterWildcard = wildcard.startIndex < length - 1
    ? symbols[wildcard.endIndex + 1]
    : null

  if (!symbolBeforeWildcard && !symbolAfterWildcard) {
    return symbols.fill(
      'C',
      wildcard.startIndex,
      wildcard.endIndex + 1
    )
  } else if (!symbolBeforeWildcard) {
    return symbols.fill(
      symbolAfterWildcard,
      wildcard.startIndex,
      wildcard.endIndex + 1
    )
  } else {
    return symbols.fill(
      symbolBeforeWildcard,
      wildcard.startIndex,
      wildcard.endIndex + 1
    )
  }
}

function fillSymbols(symbols) {
  const length = symbols.length

  let modifiedSymbols = symbols.split('')
  let latestWildcard = null

  for (let i = 0; i < length; i++) {
    const symbol = modifiedSymbols[i]
    if (symbol === '?') {
      latestWildcard = {
        startIndex: latestWildcard
          ? latestWildcard.startIndex
          : i,
        endIndex: i
      }
    } else if (latestWildcard) {
      modifiedSymbols = fillSymbolsForWildcard(modifiedSymbols, latestWildcard)
      latestWildcard = null
    }
  }

  if (latestWildcard) {
    modifiedSymbols = fillSymbolsForWildcard(modifiedSymbols, latestWildcard)
    latestWildcard = null
  }

  return modifiedSymbols.join('')
}

function calculateCost({ x, y, symbols }) {
  const length = symbols.length
  let cost = 0

  for (var i = 1; i < symbols.length; i++) {
    const symbolPair = `${symbols[i-1]}${symbols[i]}`
    if (symbolPair === 'CJ') {
      cost += x
    } else if (symbolPair === 'JC') {
      cost += y
    }
  }

  return cost
}

function solveTestCase({ x, y, symbols }) {
  const filledSymbols = fillSymbols(symbols)
  return calculateCost({ x, y, symbols: filledSymbols })
}

function solveProblem(problem) {
  return problem.testCases
    .map(testCase => solveTestCase(testCase))
    .map((result, i) => `Case #${i + 1}: ${result}`)
    .join('\n')
}

function main() {
  const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout,
    terminal: false
  })

  let problem = {
    totalTestCases: 0,
    testCases: [],
  }

  rl.on('line', function (line) {
    if (!line) {
      return
    }
    if (problem.totalTestCases === 0) {
      problem.totalTestCases = Number.parseInt(line)
    } else {
      const tokens = line.split(' ')
      problem.testCases.push({
        x: Number.parseInt(tokens[0]),
        y: Number.parseInt(tokens[1]),
        symbols: tokens[2]
      })
    }
  })
  .on('close', () => {
    console.log(solveProblem(problem))
    process.exit()
  })
}

main()