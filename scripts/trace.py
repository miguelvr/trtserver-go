import shutil
from pathlib import Path

import torch.jit
from torch import nn
from torch.nn import functional as F
from torchvision.models import resnet50

model_config = """name: "resnet50_imagenet"
platform: "pytorch_libtorch"
max_batch_size: 32
input [
  {
    name: "INPUT__0"
    data_type: TYPE_FP32
    dims: [ 3, 224, 224 ]
    format: FORMAT_NCHW
  }
]
output [
  {
    name: "OUTPUT__0"
    data_type: TYPE_INT32
    dims: [ 1 ]
    label_filename: "imagenet_labels.txt"
  },
  {
    name: "OUTPUT__1"
    data_type: TYPE_FP32
    dims: [ 1 ]
  }
]
"""


class ResNet50(nn.Module):
    mean = [0.485, 0.456, 0.406]
    std = [0.229, 0.224, 0.225]

    def __init__(self):
        super(ResNet50, self).__init__()
        self.model = resnet50(pretrained=True)

    def preprocess(self, tensor):
        dtype = tensor.dtype
        mean = torch.as_tensor(self.mean, dtype=dtype, device=tensor.device)
        std = torch.as_tensor(self.std, dtype=dtype, device=tensor.device)
        tensor.sub_(mean[None, :, None, None]).div_(std[None, :, None, None])
        return tensor

    def forward(self, x):
        x = self.model(self.preprocess(x))
        x = F.softmax(x, dim=-1)
        prob, label = torch.max(x, dim=-1, keepdim=True)
        return label.to(torch.int32), prob.to(torch.float32)


def main():
    model = ResNet50()
    model.eval()

    x = torch.rand(2, 3, 224, 224)
    traced_model = torch.jit.trace(model, (x,))

    assets_dir = Path(__file__).parent.parent / "assets"
    output_dir = assets_dir / "models" / "resnet50_imagenet" / "1"
    if not output_dir.exists():
        output_dir.mkdir(parents=True)

    output_path = str(output_dir / "model.pt")
    torch.jit.save(traced_model, output_path)

    with open(str(output_dir.parent / "config.pbtxt"), "w") as f:
        f.write(model_config)

    shutil.copy(str(assets_dir / "imagenet_labels.txt"), str(output_dir.parent / "imagenet_labels.txt"))


if __name__ == '__main__':
    main()
