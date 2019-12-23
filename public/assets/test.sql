select orders.created_at, orders.amount, items.item_name, users.first_name, users.last_name, stores.store_name
from inventory.orders
inner join inventory.items on orders.item_id = items.id
inner join inventory.users on orders.user_uuid = users.user_uuid
inner join inventory.stores on orders.store_id = stores.id
where orders.created_at between '20:28:24.0000000-05:00' and '20:28:24.999999-05:00'
and orders.user_uuid = '1bbd12ee-7acd-42f2-9849-c86c94b341bc';
