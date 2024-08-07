const fs = require('fs');


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
    /*
        Given two lists of objects with the following shape:
            {
                'id':number,
                'name':string
            }
        return a list of all the IDs that are common to both lists

        Warning: This version has a bug!!!
            1. What is the bug?
            2. How would you solve it?
            3. How would you test it?
    */

    common = []
    left.forEach((l) => {
        right.forEach((r) => {
            if ( JSON.stringify(l) == JSON.stringify(r) ) {
                common.push(l.id)
            }
        })
    })

    return common
}

function main(){
    if( ! process.argv[2] || ! process.argv[3] ){
        console.log("USAGE:", process.argv[1], "<left csv>", "<right csv>")
        return 
    }
    left = csvToJSON(process.argv[2])
    right = csvToJSON(process.argv[3])

    console.log(common_ids_with_bug(left, right))
}

main()
