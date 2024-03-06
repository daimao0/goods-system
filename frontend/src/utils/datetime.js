// 时间范围
export const datetimeCuts = [
    {
        text: '1周内',
        value: () => {
            const end = new Date()
            const start = new Date()
            start.setTime(start.getTime() - 3600 * 1000 * 24 * 7)
            return [start, end]
        },
    },
    {
        text: '1个月内',
        value: () => {
            const end = new Date()
            const start = new Date()
            start.setTime(start.getTime() - 3600 * 1000 * 24 * 30)
            return [start, end]
        },
    },
    {
        text: '3个月内',
        value: () => {
            const end = new Date()
            const start = new Date()
            start.setTime(start.getTime() - 3600 * 1000 * 24 * 90)
            return [start, end]
        },
    },
]

// 为了格式化成 "2024-03-02 13:09:10" 这样的格式，可以稍微调整一下toLocaleString的参数：
export function dateTimeToStr(datetime) {
    const padZero = (num) => ('0' + num).slice(-2);
    // 获取年、月、日、时、分、秒，并进行格式化
    let year = datetime.getFullYear();
    let month = padZero(datetime.getMonth() + 1); // 注意月份是从0开始的，所以需要+1
    let day = padZero(datetime.getDate());
    let hours = padZero(datetime.getHours());
    let minutes = padZero(datetime.getMinutes());
    let seconds = padZero(datetime.getSeconds());
    // 拼接成所需格式的字符串
    return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
}

// 为了格式化成 "2024-03-02 13:09:10" 这样的格式，可以稍微调整一下toLocaleString的参数：
export function dateTimeToDateStr(datetime) {
    const padZero = (num) => ('0' + num).slice(-2);
    // 获取年、月、日、时、分、秒，并进行格式化
    let year = datetime.getFullYear();
    let month = padZero(datetime.getMonth() + 1); // 注意月份是从0开始的，所以需要+1
    let day = padZero(datetime.getDate());
    // 拼接成所需格式的字符串
    return `${year}-${month}-${day}`;
}