from pathlib import Path

import torch.jit
from PIL import Image
from torchvision.transforms.functional import to_tensor, resize

image_path = "/home/miguel/Pictures/trump.jpg"
model_path = Path(__file__).parent.parent / "models" / "resnet50_imagenet" / "1" / "model.pt"
model = torch.jit.load(str(model_path))

img = Image.open(image_path)
img = resize(img, (224, 224))
img = to_tensor(img)
x = img.unsqueeze(0)

with torch.no_grad():
    pred = model(x)

print(pred)
