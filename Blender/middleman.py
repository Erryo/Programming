import bpy
import bmesh

pixel_pos = {
    0:(0,0),
    1:(1,1),
    2:(2,0),
    }

class Lightpipe:

    def __init__(self,pipes=1,p_size=16):
        global pixel_pos

        self.pipes = pipes
        self.p_size = p_size
        # Clear the scene
        bpy.ops.object.select_all(action='SELECT')
        bpy.ops.object.delete(use_global=False)

        bpy.ops.outliner.orphans_purge()
        bpy.ops.outliner.orphans_purge()
        bpy.ops.outliner.orphans_purge()

        material = bpy.data.materials.new("matea")
        material.use_nodes = True

        glass_node = material.node_tree.nodes.new("ShaderNodeBsdfGlass")
        m_output = material.node_tree.nodes['Material Output']
        material.node_tree.links.new(glass_node.outputs[0],m_output.inputs[0])

        glass_node.inputs[1].default_value = 1

        view_layer = bpy.context.view_layer

        for i in range(self.pipes):
            for j in range(self.p_size):
                for k in range(3):
                    # Create new light datablock.

                    mesh = bpy.data.meshes.new('Basic_Sphere')
                    pixel_data = bpy.data.lights.new(name="New Light", type='POINT')
                    # Create new object with our light datablock.

                    mesh_object = bpy.data.objects.new(name=f"m_{i}_{j}_{k}", object_data=mesh)
                    pixel_object = bpy.data.objects.new(name=f"{i}_{j}_{k}", object_data=pixel_data)

                    # Link light object to the active collection of current view layer,
                    # so that it'll appear in the current scene.
                    view_layer.active_layer_collection.collection.objects.link(pixel_object)


                    # Place light to a specified location.
                    pixel_object.location = ((0.0 + i * 6) + pixel_pos.get(k)[0]  , 0.0 +pixel_pos.get(k)[1] , 0.0 + j)
aaaaaa;<D-A><D-A><D-A><D-A><D-A><D-A><D-A><D-A>aa
                    # And finally select it and make it active.
                    pixel_object.select_set(True)
                    view_layer.objects.active = pixel_object

                    # Insert first frame
                    pixel_object.data.color = [0,0,0]
                    pixel_object.data.keyframe_insert("color",frame=0)


                    bm = bmesh.new()
                    bmesh.ops.create_uvsphere(bm, u_segments=16, v_segments=8, radius=0.3)
                    bm.to_mesh(mesh)
                    bm.free()

                    view_layer.active_layer_collection.collection.objects.link(mesh_object)

                    # Place light to a specified location.
                    mesh_object.location = ((0.0 + i * 6) + pixel_pos.get(k)[0]  , 0.0 +pixel_pos.get(k)[1] , 0.0 + j)
                    mesh_object.data.materials.append(material)
                    # And finally select it and make it active.
                    mesh_object.select_set(True)
                    view_layer.objects.active = mesh_object

    def transform_color(self,color):
        blender_color = []
        for i in range(3):
            blender_color.append(color[i]/255)
        #blender_color.append(1.0)
        return blender_color

    def set_pixel_0(self, p, x, color):
        color = self.transform_color(color)
        light_obj = bpy.data.objects[f"{p}_{x}_0"]
        light_obj.data.energy = 100
        light_obj.data.color = color

    def set_pixel_1(self, p, x, color):
        color = self.transform_color(color)
        light_obj = bpy.data.objects[f"{p}_{x}_1"]
        light_obj.data.energy = 100
        light_obj.data.color = color
    def set_pixel_2(self, p, x, color):
        color = self.transform_color(color)
        light_obj = bpy.data.objects[f"{p}_{x}_2"]
        light_obj.data.energy = 100
        light_obj.data.color = color

    def set_pixel_y(self, p, x, color, y):
        if y == 0:
            self.set_pixel_0(p, x, color)
        if y == 1:
            self.set_pixel_1(p, x, color)
        if y == 2:
            self.set_pixel_2(p, x, color)

    def show(self,frame):
        for p in range(self.pipes):
            for x in range(self.p_size):
                for k in range(3):
                    light_obj = bpy.data.objects[f"{p}_{x}_{k}"]
                    light_obj.data.keyframe_insert("color",frame=frame)
