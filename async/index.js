// https://levelup.gitconnected.com/use-go-channels-as-promises-and-async-await-ee62d93078ec
const sleep = (time) => {
  return new Promise((resolve) => {
    setTimeout(() => resolve(), time)
  })
}

const longRunningTask = async () => {
  sleep(3000)
  return Math.floor(Math.random() * Math.floor(100))
}

const r = await longRunningTask()
console.log(r)
