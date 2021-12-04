const fs = require('fs')

let data
try {
    data = fs.readFileSync('./input.txt', 'utf8')
} catch (err) {
    console.error(err)
}

let lines = data.split("\n")
let callNums = lines.shift()
callNums = callNums.split(",").map(el => parseInt(el))

lines.shift()
let boards = []
let currentBoard = []

lines.forEach((line, index) => {
    if (line === '') {
       boards.push(currentBoard)
        currentBoard = []
        return
    }

    currentBoard.push(line.split(" ").filter(el => el !== '').map(el => [parseInt(el), 0]))

    if (index === lines.length - 1) {
        boards.push(currentBoard)
        currentBoard = []
        return
    }
})

function markNum(boards, calledNum) {
    boards.forEach(board => {
        board.forEach(row => {
            row.forEach(num => {
                if (num[0] === calledNum) num[1] = 1
            })
        })
    })

    return boards
}

function checkForBingo(board, num) {
    let hasBingo = false
    board.forEach(row => {
        if (hasBingo) return
        let compactRow = row.map(el => el[1]).join("")
        if (num === 24) console.log(compactRow, !compactRow.includes("0"))
        if (!compactRow.includes("0")) hasBingo = true
    })

    if (hasBingo) return true

    for (let i = 0; i < board[0].length; i++) {
        if (hasBingo) return
        let col = ''
        board.forEach(row => {
            col = col + row[i][1]
        })
        if (!col.includes(0)) hasBingo = true
    }

    return hasBingo
}

let foundBingo = false
callNums.forEach((num, numIndex) => {
    if (foundBingo) return
    boards = markNum(boards, num)
    boards.forEach((board, index) => {
        if (foundBingo) return
        let bingo = checkForBingo(board, num)
        if (bingo) {
            foundBingo = true
            let bingoBoard = boards[index]
            let unmarked = 0
            bingoBoard.forEach(row => {
                row.forEach(el => {
                    if (el[1] === 0) unmarked += el[0]
                })
            })

            console.log(`unmarked: ${unmarked}, num: ${num}, result: ${unmarked * num}`)
        }
    })
})
