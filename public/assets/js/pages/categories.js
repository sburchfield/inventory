$(document).ready( function () {

  const currentCategoriesTable = $('#currentCategories').DataTable({
  "columnDefs": [
    { className: "category", "targets": [ 1 ] }
  ]
});
  const deletedCategoriesTable = $('#deletedCategories').DataTable({
  "columnDefs": [
    { className: "category", "targets": [ 1 ] }
  ]
});

  $("#categoriesForm").submit(function(e){
    e.preventDefault()

    let data = $("#categoriesForm").serialize()

    $.ajax({
      method: "POST",
      url: "http://localhost:3000/updateCategories",
      data: data
    }).done(function( data ) {

      console.log(data);

        $("#responseMessage").empty()
        $("#responseMessage").append('<h5 class="animated fadeOut">'+data.Message+'</h5>');

        currentCategoriesTable.clear()

        data.Categories.forEach(function(value){

          let rmvButton = '<td><button id="'+value.ID+'" class="btn btn-danger remove">Remove</button></td>'

          let array = [value.ID, value.Category, value.Description, rmvButton]

          $('#currentCategories').DataTable().row.add(array).draw()

            $("#categoriesForm")[0].reset()

        })



    })

  });//end of submit


  $(document).on("click", ".remove", function(){

    let category_id = $(this).attr('id')
    let category = $(this).parents("tr").find(".category").html()

    $("#responseMessage").empty();


    bootbox.confirm({
        title: "Destroy Category?",
        message: "Do you want to destroy this category and its items?",
        buttons: {
            cancel: {
                label: 'Cancel'
            },
            confirm: {
                label: 'Confirm'
            }
        },
        callback: function (result) {

          if(result){
            $.ajax({
              method: "GET",
              url: "http://localhost:3000/removeCategory/"+ category_id + "/" + category,
            }).done(function( data ) {

                deletedCategoriesTable.clear()

                if (data.RemoveBy == category_id){
                  $("#" + data.RemoveBy).parent().parent().remove()
                  $("#responseMessage").append('<h5 class="animated fadeOut">'+data.Message+'</h5>');

                  data.Categories.forEach(function(value){

                    let restoreButton = '<td><button id="deleted_'+value.ID+'" class="btn btn-success restore">Restore</button></td>'

                    let array = [value.ID, value.Category, value.Description, restoreButton]

                    $('#deletedCategories').DataTable().row.add(array).draw()

                  })

                }else{
                  $("#responseMessage").append('<h5 class="animated fadeOut">Error! Try Again!</h5>');
                }

                })

          }
        }
    });

    })

    $(document).on("click", ".restore", function(){

      let category_id = $(this).attr('id').split("_")
      let category = $(this).parents("tr").find(".category").html()

      $("#responseMessage").empty();

      bootbox.confirm({
          title: "Restore Category?",
          message: "Do you want to restore this category and all of its associated items?",
          buttons: {
              cancel: {
                  label: '<i class="fa fa-times"></i> Cancel'
              },
              confirm: {
                  label: '<i class="fa fa-check"></i> Confirm'
              }
          },
          callback: function (result) {
              if(result){
                $.ajax({
                  method: "GET",
                  url: "http://localhost:3000/restoreCategory/"+ category_id[1] + "/" + category,
                }).done(function( data ) {
                  currentCategoriesTable.clear()

                  if( data.RemoveBy == category_id[1]){
                    $("#deleted_" + data.RemoveBy).parent().parent().remove()
                    $("#responseMessage").append('<h5 class="animated fadeOut">'+data.Message+'</h5>');

                    data.Categories.forEach(function(value){

                      let rmvButton = '<td><button id="'+value.ID+'" class="btn btn-danger remove">Remove</button></td>'

                      let array = [value.ID, value.Category, value.Description, rmvButton]

                      $('#currentCategories').DataTable().row.add(array).draw()

                    })
                  }else{
                    $("#responseMessage").append('<h5 class="animated fadeOut">Error! Try Again!</h5>');
                  }



                    })
              }
          }
      });


      })


});//end of doc ready
