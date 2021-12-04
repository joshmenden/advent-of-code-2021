const fs = require('fs')

let data
try {
    data = fs.readFileSync('./test.txt', 'utf8')
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
        // if (num === 24) console.log(compactRow, !compactRow.includes("0"))
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

let bingoBoardIndices = [...Array(boards.length).keys()]

function sumUnmarked(board) {
    let unmarked = 0
    board.forEach(row => {
        row.forEach(el => {
            if (el[1] === 0) unmarked += el[0]
        })
    })

    return unmarked
}

console.log(bingoBoardIndices)
callNums.forEach((num, numIndex) => {
    boards = markNum(boards, num)
    boards.forEach((board, index) => {
        if (!bingoBoardIndices.includes(index)) return
        let bingo = checkForBingo(board, num)
        if (bingo) {
            console.log(`found bingo board index ${index}`)
            let indexInMaster = bingoBoardIndices.indexOf(index);
            // console.log(`index in master is ${indexInMaster}`)
            // console.log(`removing ${bingoBoardIndices[bingoBoardIndex]}`)
            if (indexInMaster >= 0) bingoBoardIndices.splice(indexInMaster, 1)
            // console.log('new master', bingoBoardIndices)
        }


        if (bingoBoardIndices.length === 1) {
            console.log('HIII', bingoBoardIndices)
            let unmarkedSum = sumUnmarked(boards[bingoBoardIndices[0]])
            console.log(`unmarked: ${unmarkedSum}, num: ${num}, result: ${unmarkedSum * num}`)
        }
    })
})
