func totalWaviness(num1 int, num2 int) int {
    totalWaveAllRange := 0

    for i := num1 ; i<=num2 ; i++{
        totalWaveAllRange += totalWave(i)
    }
    return totalWaveAllRange
}

//4848
func totalWave(num int) int {
    totalWave := 0
    digits := []int{}
    for num > 0 {
        digit := num % 10
        digits = append(digits , digit)
        num /= 10
    }

    for i := 1; i < len(digits)-1 ; i++{
        currentMid := digits[i]

        if currentMid < digits[i-1] && currentMid < digits[i+1]{
            totalWave++
        }

        if currentMid > digits[i-1] && currentMid > digits[i+1]{
            totalWave++
        }
    }

    return totalWave
}