box.cfg{}
p = box.schema.space.create("posts", {if_not_exists= true} )
p:create_index('primary',{parts = {1,"unsigned"},if_not_exists= true})
p:create_index('user_id',{parts = {2,"unsigned"},if_not_exists= true})