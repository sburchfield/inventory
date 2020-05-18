$(document).ready( function () {

  const currentItemsTable = $('#currentItems').DataTable();
  const deletedItemsTable = $('#deletedItems').DataTable();

  $("#itemsForm").submit(function(e){
    e.preventDefault()

    let data = $("#itemsForm").serialize()

    $.ajax({
      method: "POST",
      url: "http://localhost:3000/api/updateItems",
      data: data
    }).done(function( data ) {
        $("#responseMessage").empty()
        $("#responseMessage").append('<h5 class="animated fadeOut">'+data.Message+'</h5>');

        currentItemsTable.clear()

        data.Items.forEach(function(value){

          let rmvButton = '<td><button id="'+value.ID+'" class="btn btn-danger remove">Remove</button></td>'

          let array = [value.ID, value.ItemName, value.ItemCost, value.ItemPrice, value.Category, rmvButton]

          $('#currentItems').DataTable().row.add(array).draw()

            $("#itemsForm")[0].reset()

        })

    })

  });//end of submit


  $(document).on("click", ".remove", function(){

    let item_id = $(this).attr('id')

    $("#responseMessage").empty();

    $.ajax({
      method: "GET",
      url: "http://localhost:3000/api/removeItem/"+ item_id,
    }).done(function( data ) {

        deletedItemsTable.clear()

        if (data.RemoveBy == item_id){
          $("#" + data.RemoveBy).parent().parent().remove()
          $("#responseMessage").append('<h5 class="animated fadeOut">'+data.Message+'</h5>');

          data.Items.forEach(function(value){

            let restoreButton = '<td><button id="deleted_'+value.ID+'" class="btn btn-success restore">Restore</button></td>'

            let array = [value.ID, value.ItemName, value.ItemCost, value.ItemPrice, value.Category, restoreButton]

            $('#deletedItems').DataTable().row.add(array).draw()

          })

        }else{
          $("#responseMessage").append('<h5 class="animated fadeOut">Error! Try Again!</h5>');
        }

        })
    })

    $(document).on("click", ".restore", function(){

      let item_id = $(this).attr('id').split("_")

      $("#responseMessage").empty();

      $.ajax({
        method: "GET",
        url: "http://localhost:3000/api/restoreItem/"+ item_id[1],
      }).done(function( data ) {
        currentItemsTable.clear()

        if( data.RemoveBy == item_id[1]){
          $("#deleted_" + data.RemoveBy).parent().parent().remove()
          $("#responseMessage").append('<h5 class="animated fadeOut">'+data.Message+'</h5>');

          data.Items.forEach(function(value){

            let rmvButton = '<td><button id="'+value.ID+'" class="btn btn-danger remove">Remove</button></td>'

            let array = [value.ID, value.ItemName, value.ItemCost, value.ItemPrice, value.Category, rmvButton]

            $('#currentItems').DataTable().row.add(array).draw()

          })
        }else{
          $("#responseMessage").append('<h5 class="animated fadeOut">Error! Try Again!</h5>');
        }



          })
      })



});//end of doc ready
