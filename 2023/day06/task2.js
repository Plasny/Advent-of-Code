/**
 * @returns {[number[], number[]]} Array with times and distances
 */
async function getData() {
    /**
     * An array of command line arguments
     * @type {string[]}
     */
    const args = process.argv
    args.splice(0, 2);

    /**
     * @type {BunFile}
     */
    let f

    /**
     * @type {string}
     */
    let ts, ds

    if (args.length == 1) {
        f = Bun.file(args[0])
    } else {
        f = Bun.stdin
    }


    try {
        [ts, ds] = (await f.text()).split('\n');
    } catch {
        console.error("Could not open file")
        process.exit(1)
    }

    ts = ts.replace(/Time:\s+/, "");
    ts = ts.replaceAll(/\s+/g, "");

    ds = ds.replace(/Distance:\s+/, "");
    ds = ds.replaceAll(/\s+/g, "");

    return [parseInt(ts), parseInt(ds)]
}

const [time, distance] = await getData()

let num = 0;
for (let j = 1; j < time; j++) {
    const ourDistance = j * (time - j)

    if (ourDistance > distance) {
        num++
    }
}

console.log(num)
