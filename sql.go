
package main

var queries map[string]string = map[string]string{

  "getUser" : `SELECT id, user_uuid, user_email, username, password_hash, status, reset_time, password_reset_hash
               FROM inventory.users
               WHERE username = ?;`,

 "getUserByUuid" : `SELECT id, user_uuid, user_email, username, password_hash, status, reset_time, password_reset_hash
              FROM inventory.users
              WHERE user_uuid = ?;`,

  "resetPass" : `UPDATE inventory.users
                 SET password_hash = ?, status = 'active', password_reset_hash = null, reset_time = null
                 WHERE user_uuid = ?;`,

  "requestReset" : `UPDATE inventory.users
                    SET status = 'reset', password_reset_hash = ?, reset_time = ?
                    WHERE username = ?;`,

  "getOrderEmail" : `select orders.created_at, orders.amount, items.item_name, users.first_name, users.last_name, stores.store_name
                    from inventory.orders
                    inner join inventory.items on orders.item_id = items.id
                    inner join inventory.users on orders.user_uuid = users.user_uuid
                    inner join inventory.stores on orders.store_id = stores.id
                    where orders.created_at between ? and ?
                    and orders.user_uuid = ?;`,

}
