$(document).ready( function () {

  function onlyUnique(value, index, self) {
    return self.indexOf(value) === index;
  }

function getLatestOrders(user_uuid){
  $.ajax({
    method: "GET",
    url: "/api/getLatestOrders/" + user_uuid
  })
  .done(function( data ) {

    data.LatestOrders.forEach(item => {
      // console.log(item);

      $("#" + item.ItemID).val(item.Amount)
    })

    if(data.LatestOrders.length >= 1){
      append = true
      $('#store').val(data.LatestOrders[0].StoreID)
      $("#store").prop('disabled', 'disabled')
    }



  })
  .fail(function( err ){

    console.log("error: ", err);

  })
}

  $.ajax({
    method: "GET",
    url: "/api/select/stores"
  })
  .done(function( data ) {
    let stores = data.Stores
    let html;
      stores.forEach(element => {
        html = '<option value="'+element.ID+'">'+element.StoreName+'</option>'
        $("#store").append(html)
      });
  })
  .fail(function( err ){
    console.log("error: ", err);
  })

  $.ajax({
    method: "GET",
    url: "/api/select/items"
  })
  .done(function( data ) {

    // console.log(data);

    let items = data.Items
    let html;
    let user= $("#user").val();

    let categoryArray = []
    let divArray = []

    items.forEach(element => {

      categoryArray.push(element.Category.replace(/ /g, "_"))

    })

    divArray = categoryArray.filter(onlyUnique)

    divArray.forEach( element => {
      html = `
          <div class="col-lg-6">
            <div class="itemsWrapper card">
              <h3 class="text-center category-header">`+element.replace(/_/g, " ")+`</h3>
              <div class="row" id=`+element+`></div>
            </div>
          </div>
      `
      $("#itemsGroup").append(html)
    })

    items.forEach(element => {

      html = `
              <div class="col-6 form-group">
                <label class="item-label" for="`+element.ID+`">`+element.ItemName+`</label>
                <input class="form-control item-input" name="`+element.ID+`" id="`+element.ID+`" type="number">
              </div>
             `
      $("#"+element.Category.replace(/ /g, "_")).append(html)
    });

    getLatestOrders(user)

  })
  .fail(function( err ){

    console.log("error: ", err);

  })

  $("#itemsForm").submit(function(e){

    e.preventDefault();

    let store = $("#store").val();
    let user= $("#user").val();

    // console.log(store);
    // console.log(user);

    let data = $(this).serializeArray()

    // console.log(data);

    data = data.map(function(obj) {
        obj['ItemId'] = obj['name']; // Assign new key
        obj['Amount'] = parseInt(obj['value']); // Assign new key
        delete obj['name']; // Delete old key
        delete obj['value']; // Delete old key
        return obj;
    });

    // console.log(data);

    data.forEach(element => {
      // console.log(element);

      element.UserUUID = user
      element.StoreId = store
    })

    let newData = data.filter(value => isNaN(value.Amount) === false)
    // console.log(newData);

    $("html, body").animate({ scrollTop: 0 }, "slow");

    $.ajax({
      method: "POST",
      url: "/api/updateOrders",
      contentType: "application/json",
      dataType: "text",
      data: JSON.stringify(newData)
    })
    .done(function( data ) {

      $('#responseModal').modal()
      $("#responseMessage").html(data)
      $('#store').prop('disabled', 'disabled')

    })
    .fail(function( err ){

      console.log("error: ", err);

    })


  })

})
