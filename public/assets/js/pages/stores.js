$(document).ready( function () {

  const currentStoresTable = $('#currentStores').DataTable();
  const deletedStoresTable = $('#deletedStores').DataTable();

  $("#storesForm").submit(function(e){
    e.preventDefault()

    let data = $("#storesForm").serialize()

    $.ajax({
      method: "POST",
      url: "http://localhost:3000/updateStores",
      data: data
    }).done(function( data ) {
        $("#responseMessage").empty()
        $("#responseMessage").append('<h5 class="animated fadeOut">'+data.Message+'</h5>');

        currentStoresTable.clear()

        data.Stores.forEach(function(value){

          let rmvButton = '<td><button id="'+value.ID+'" class="btn btn-danger remove">Remove</button></td>'

          let array = [value.ID, value.StoreName, value.Address, value.PhoneNumber, value.City, value.State, value.ZipCode, rmvButton]

          $('#currentStores').DataTable().row.add(array).draw()

            $("#storesForm")[0].reset()

        })



    })

  });//end of submit


  $(document).on("click", ".remove", function(){

    let store_id = $(this).attr('id')

    $("#responseMessage").empty();

    $.ajax({
      method: "GET",
      url: "http://localhost:3000/removeStore/"+ store_id,
    }).done(function( data ) {

        deletedStoresTable.clear()

        if (data.RemoveBy == store_id){
          $("#" + data.RemoveBy).parent().parent().remove()
          $("#responseMessage").append('<h5 class="animated fadeOut">'+data.Message+'</h5>');

          data.Stores.forEach(function(value){

            let restoreButton = '<td><button id="deleted_'+value.ID+'" class="btn btn-success restore">Restore</button></td>'

            let array = [value.ID, value.StoreName, value.Address, value.PhoneNumber, value.City, value.State, value.ZipCode, restoreButton]

            $('#deletedStores').DataTable().row.add(array).draw()

          })

        }else{
          $("#responseMessage").append('<h5 class="animated fadeOut">Error! Try Again!</h5>');
        }

        })
    })

    $(document).on("click", ".restore", function(){

      let store_id = $(this).attr('id').split("_")

      $("#responseMessage").empty();

      $.ajax({
        method: "GET",
        url: "http://localhost:3000/restoreStore/"+ store_id[1],
      }).done(function( data ) {
        currentStoresTable.clear()

        if( data.RemoveBy == store_id[1]){
          $("#deleted_" + data.RemoveBy).parent().parent().remove()
          $("#responseMessage").append('<h5 class="animated fadeOut">'+data.Message+'</h5>');

          data.Stores.forEach(function(value){

            let rmvButton = '<td><button id="'+value.ID+'" class="btn btn-danger remove">Remove</button></td>'

            let array = [value.ID, value.StoreName, value.Address, value.PhoneNumber, value.City, value.State, value.ZipCode, rmvButton]

            $('#currentStores').DataTable().row.add(array).draw()

          })
        }else{
          $("#responseMessage").append('<h5 class="animated fadeOut">Error! Try Again!</h5>');
        }



          })
      })



});//end of doc ready
