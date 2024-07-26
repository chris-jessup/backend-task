const fs = require('fs');

console.log(process.argv[2], process.argv[3])

function csvToJSON(filename) {
    let csv = []
    let buf = fs.readFileSync(filename)
    csv = String(buf) // convert the buffer to a string
        .split('\n') // Split the string into an array where each item contains one line
        .filter(Boolean); // Remove any empty lines

      // Do the rest of the operations on the CSV data here
      const [headers, ...data] = csv.map(row => row.split(','));
      return data.map(row => {
        const rowObject = {};
        row.forEach((value, index) => {
          rowObject[headers[index]] = value;
        });
        return rowObject
      });
}



function common_ids_with_bug(left, right){
    // """ 

    // Given two lists of dictionaries with the following shape:
    //     {
    //         'id':number,
    //         'name':string
    //     }
    // return a tuple of all the IDs that are common to both lists

    // Warning: This version has a bug!!!
    //     1. What is the bug?
    //     2. How would you solve it?
    //     3. How would you test it?

    // """

    common = []
    left.forEach((l) => {
        right.forEach((r) => {
            console.log(l,r)
            if ( JSON.stringify(l) == JSON.stringify(r) ) {
                common.push(l)
            }
        })
    })

    return common
}

left = csvToJSON(process.argv[2])
right = csvToJSON(process.argv[3])

console.log(common_ids_with_bug(left, right))
