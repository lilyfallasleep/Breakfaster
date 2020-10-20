var index = 1
var api_url = 'https://223d4f229adf.ap.ngrok.io'

window.onload = function () {
  document.getElementById('eid').addEventListener('input', async (e) => {
    const eid_or_cardNum = e.target.value
    if (authenticate(eid_or_cardNum)) {
      document.getElementById('eid').classList.remove('invalid')
      document.getElementById('eid').classList.add('valid')
      // call API
      const { food, eid, pick } = await fetchOrderByEidOrCardNum(eid_or_cardNum)
      document.querySelector('#show-area h3').textContent = eid || '找不到此人'
      document.querySelector('#order-name').textContent = pick
        ? '(已取餐)' + food
        : food || '未點餐'
      console.log({ food, eid, pick })
      if (eid && food && !pick) {
        appendToTable(eid, food)
      }
      document.getElementById('eid').value = ''
    } else if (eid_or_cardNum !== '') {
      document.getElementById('eid').classList.remove('valid')
      document.getElementById('eid').classList.add('invalid')
    } else {
      document.getElementById('eid').classList.remove('invalid')
      document.getElementById('eid').classList.add('valid')
    }
  })
}

function appendToTable(eid, order) {
  // Get a reference to the table
  let tableRef = document.getElementsByTagName('tbody')[0]

  // Insert a row at the end of the table
  let newRow = tableRef.insertRow(0)

  // Append a text node to the cell
  let idxNode = document.createElement('th')
  idxNode.innerHTML = index
  newRow.appendChild(idxNode)

  let eidNode = document.createTextNode(eid)
  let orderNode = document.createTextNode(order)
  let tzoffset = (new Date()).getTimezoneOffset() * 60000
  let timeNode = document.createTextNode((new Date(Date.now() - tzoffset)).toISOString().slice(11, 19))

  let newCell = newRow.insertCell(-1)
  newCell.appendChild(eidNode)

  newCell = newRow.insertCell(-1)
  newCell.appendChild(orderNode)

  newCell = newRow.insertCell(-1)
  newCell.appendChild(timeNode)

  index++
}

Date.prototype.yyyymmdd = function () {
  let mm = this.getMonth() + 1; // getMonth() is zero-based
  let dd = this.getDate();

  return [this.getFullYear(),
  (mm > 9 ? '' : '0') + mm,
  (dd > 9 ? '' : '0') + dd
  ].join('-');
};

function fetchOrderByEidOrCardNum(eid_or_cardNum) {
  const payload = eid_or_cardNum
  const type = payload.length === 7 ? 'eid' : 'card'
  let tzoffset = (new Date()).getTimezoneOffset() * 60000
  let curDate = (new Date(Date.now() - tzoffset)).yyyymmdd()

  return fetch(
    api_url +
    '/api/v1/order?type=' +
    type +
    '&payload=' +
    payload +
    '&date=' +
    curDate,
    {
      mode: 'cors',
    }
  )
    .then((res) => {
      return res.json()
    })
    .then((res) => {
      console.log(res)
      var get_res = res
      if (res.pick) {
        return { food: get_res.food_name, eid: get_res.emp_id, pick: true }
      } else {
        const put_data = {
          emp_id: get_res.emp_id,
          date: '2020-09-25',
        }
        console.log(put_data)
        return fetch(api_url + '/api/v1/order/pick', {
          method: 'PUT',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(put_data),
        })
          .then((res) => {
            return res.json()
          })
          .then((res) => {
            console.log('PUT: ' + res)
            return { food: get_res.food_name, eid: get_res.emp_id, pick: false }
          })
      }
    })

    .catch((err) => {
      console.log(err)
    })
}

function authenticate(str) {
  return (
    (str.length === 7 && str.startsWith('LW') && !isNaN(str.substring(2))) ||
    str.length === 10
  )
}
