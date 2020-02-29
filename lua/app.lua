box.cfg{}
box.schema.sequence.create('P',{if_not_exists=true})
p = box.schema.space.create("posts", {if_not_exists= true} )
p:create_index('primary',{sequence='P',if_not_exists= true})
p:create_index('user_id',{unique=false, parts = {2,"unsigned"},if_not_exists= true})
--box.space.posts.index.primary:drop()