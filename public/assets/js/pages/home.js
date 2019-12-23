$(document).ready( function () {

  function onlyUnique(value, index, self) {
    return self.indexOf(value) === index;
  }

  $.ajax({
    method: "GET",
    url: "http://localhost:3000/select/stores"
  })
  .done(function( data ) {

    // console.log(data);

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
    url: "http://localhost:3000/select/items"
  })
  .done(function( data ) {

    // console.log(data);

    let items = data.Items
    let html;

    let categoryArray = []
    let divArray = []

    items.forEach(element => {

      categoryArray.push(element.Category.replace(/ /g, "_"))

    })

    divArray = categoryArray.filter(onlyUnique)

    divArray.forEach( element => {
      html = `
          <div class="col-md-6 itemsWrapper">
          <h3>`+element.replace(/_/g, " ")+`</h3>
            <div class="row" id=`+element+`></div>
          </div>
      `
      $("#itemsGroup").append(html)
    })

    items.forEach(element => {

      html = `
              <div class="col-md-6">
                <label>`+element.ItemName+`</label>
                <input class="form-control" name="`+element.ID+`" type="number">
              </div>
             `
      $("#"+element.Category.replace(/ /g, "_")).append(html)
    });

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
        obj['Amount'] = obj['value']; // Assign new key
        delete obj['name']; // Delete old key
        delete obj['value']; // Delete old key
        return obj;
    });

    // console.log(data);

    data.forEach(element => {
      console.log(element);

      element.UserUuid = user
      element.StoreId = store
    })

    console.log(data);

    $.ajax({
      method: "POST",
      url: "http://localhost:3000/updateOrders",
      contentType: "application/json",
      dataType: "text",
      data: JSON.stringify(data)
    })
    .done(function( data ) {
      console.log(data);

    })
    .fail(function( err ){

      console.log("error: ", err);

    })


  })


  $("#next").click(function(){
    
  })

})
